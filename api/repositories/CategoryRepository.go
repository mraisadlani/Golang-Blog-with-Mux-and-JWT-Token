package repositories

import (
	"go-blog-jwt-token/api/entities"
	"go-blog-jwt-token/api/payloads/request"
)

type CategoryRepository interface {
	GetAll(*request.Pagination) (interface{}, error, int)
	FindCategory(int64) (entities.Category, error)
	CreateCategory(request.CategoryRequest) (bool, error)
	UpdateCategory(request.CategoryRequest, int64) (bool, error)
	DeleteCategory(int64) (bool, error)
}