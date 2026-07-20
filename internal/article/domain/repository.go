package domain

import (
	"blog-2026ddd-server/internal/article/dto"
	"blog-2026ddd-server/internal/shared/api"
	"context"
)

type Repository interface {
	//Create(ctx context.Context, article *Article) error
	//Delete(ctx context.Context, id uint) error
	//Update(ctx context.Context, article *Article) error
	//List(ctx context.Context) ([]*dto.ArticleListItem, error)
	//FindByID(ctx context.Context, id uint) (*Article, error)

	// --- 文章 Article 相关 ---
	//CreateArticle(ctx context.Context, article *Article) error
	//UpdateArticle(ctx context.Context, article *Article) error
	//DeleteArticle(ctx context.Context, id uint) error
	GetArticleByID(ctx context.Context, id uint) (*Article, error)
	//ListArticles(ctx context.Context) ([]*dto.ArticleListItem, error)
	ListArticles(ctx context.Context, param api.Page) (api.PageResult[*dto.ArticleListItem], error)

	//ListArticles(ctx context.Context, page, pageSize int) ([]*Article, int64, error)

	// --- 分类 ArticleType 相关 ---
	//CreateType(ctx context.Context, articleType *ArticleType) error
	//UpdateType(ctx context.Context, articleType *ArticleType) error
	//DeleteType(ctx context.Context, id uint) error
	//GetTypeByID(ctx context.Context, id uint) (*ArticleType, error)
	//ListTypes(ctx context.Context) ([]*ArticleType, error)
}
