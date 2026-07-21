package api

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/humatest"
)

func TestUnifiedResponses(t *testing.T) {
	originalNewError := huma.NewError
	originalNewErrorWithContext := huma.NewErrorWithContext
	t.Cleanup(func() {
		huma.NewError = originalNewError
		huma.NewErrorWithContext = originalNewErrorWithContext
	})

	ConfigureHumaErrors()
	_, testAPI := humatest.New(t, huma.DefaultConfig("Test API", "1.0.0"))

	type input struct {
		Value int `query:"value" required:"true"`
	}

	huma.Get(testAPI, "/test", func(_ context.Context, in *input) (*BodyResponse[int], error) {
		return NewSuccessResponse(in.Value), nil
	})

	t.Run("success", func(t *testing.T) {
		response := testAPI.Get("/test?value=42")
		if response.Code != http.StatusOK {
			t.Fatalf("expected status 200, got %d", response.Code)
		}

		var body Response[int]
		if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
			t.Fatalf("decode response: %v", err)
		}
		if body.Code != CodeSuccess || body.Msg != "success" || body.Data != 42 {
			t.Fatalf("unexpected success response: %+v", body)
		}
	})

	t.Run("validation error", func(t *testing.T) {
		response := testAPI.Get("/test?value=invalid")
		if response.Code != http.StatusUnprocessableEntity {
			t.Fatalf("expected status 422, got %d", response.Code)
		}

		var body ErrorResponse
		if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
			t.Fatalf("decode response: %v", err)
		}
		if body.Code != CodeFailure || body.Msg != "参数校验失败" || body.Data != nil {
			t.Fatalf("unexpected error response: %+v", body)
		}
	})
}

func TestNewCodeError(t *testing.T) {
	err := NewCodeError(http.StatusNotFound, 10001, "文章不存在")
	if err.GetStatus() != http.StatusNotFound || err.Code != 10001 || err.Msg != "文章不存在" || err.Data != nil {
		t.Fatalf("unexpected code error: %+v", err)
	}
}

func TestInternalErrorMessage(t *testing.T) {
	tests := []struct {
		name     string
		message  []string
		expected string
	}{
		{name: "default", expected: "服务器内部错误"},
		{name: "with context", message: []string{"文章分类加载失败"}, expected: "服务器内部错误：文章分类加载失败"},
		{name: "empty context", message: []string{""}, expected: "服务器内部错误"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := InternalError(test.message...)
			if err.GetStatus() != http.StatusInternalServerError || err.Code != CodeFailure || err.Msg != test.expected || err.Data != nil {
				t.Fatalf("unexpected internal error: %+v", err)
			}
		})
	}
}
