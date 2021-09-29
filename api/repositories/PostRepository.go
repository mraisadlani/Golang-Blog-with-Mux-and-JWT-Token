package repositories

import (
	"go-blog-jwt-token/api/entities"
	"go-blog-jwt-token/api/payloads/request"
)

type PostRepository interface {
	GetAll(*request.Pagination) (interface{}, error, int)
	CreatePost(request.PostRequest, string, []string, []string) (bool, error)
	FindPost(int64, int64) (entities.Post, error)
	UpdatePost(request.PostRequest, []string, string, []string) (bool, error)
	DeletePost(int64, int64) (bool, error)
	PublishPost(int64, int64) (bool, error)
	CancelPost(int64, int64) (bool, error)
}