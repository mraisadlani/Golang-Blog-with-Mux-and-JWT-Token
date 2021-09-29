package repositories

import (
	"go-blog-jwt-token/api/entities"
	"go-blog-jwt-token/api/payloads/request"
)

type UserRepository interface {
	GetAll(*request.Pagination) (interface{}, error, int)
	FindById(int64) (entities.User, error)
	FindByEmail(string) (entities.User, error)
	FindByUsername(string) (entities.User, error)
	Insert(request.UserRequest) (bool, error)
	Update(request.UserRequest, int64) (bool, error)
	Delete(int64) (bool, error)
	Enable(int64) (bool, error)
	Disable(int64) (bool, error)
	UploadImage(string, int64) (bool, error)
}