package repositories

import (
	"go-blog-jwt-token/api/entities"
	"go-blog-jwt-token/api/payloads/request"
)

type ArticleRepository interface {
	GetAll(*request.Pagination) (interface{}, error, int)
	FindArticle(int64) (entities.Article, error)
	CreateArticle(request.ArticleRequest) (bool, error)
	UpdateArticle(request.ArticleRequest, int64) (bool, error)
	DeleteArticle(int64) (bool, error)
}