package repositories

import (
	"go-blog-jwt-token/api/entities"
	"go-blog-jwt-token/api/payloads/request"
)

type TagRepository interface {
	GetAll(*request.Pagination) (interface{}, error, int)
	FindTag(int64) (entities.Tag, error)
	CreateTag(request.TagRequest) (bool, error)
	UpdateTag(request.TagRequest, int64) (bool, error)
	DeleteTag(int64) (bool, error)
}