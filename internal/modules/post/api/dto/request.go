package dto

type GetPostRequest struct {
	ID int `path:"id" minimum:"1" doc:"说说 ID"`
}

// CreatePostRequest 对应 POST /post 的 JSON 请求体。
type CreatePostRequest struct {
	Body struct {
		Content string `json:"content" minLength:"1" maxLength:"100" doc:"说说内容"`
	}
}

// UpdatePostRequest 对应 PUT /post/{id}。
type UpdatePostRequest struct {
	ID int `path:"id" minimum:"1" doc:"文章 ID"`

	Body struct {
		Content string `json:"content" minLength:"1" maxLength:"100" doc:"文章内容"`
	}
}

// DeletePostRequest 对应 DELETE /post/{id}。
type DeletePostRequest struct {
	ID int `path:"id" minimum:"1" doc:"文章 ID"`
}
