package repositories

import (
	"go-blog-jwt-token/api/entities"
	"go-blog-jwt-token/api/payloads/request"
)

type RoleRepository interface {
	GetAll(*request.Pagination) (interface{}, error, int)
	FindByNamaRole(string) (entities.Role, error)
	FindById(int64) (entities.Role, error)
	Insert(request.RoleRequest) (bool, error)
	Update(request.RoleRequest, int64) (bool, error)
	Delete(int64) (bool, error)
	Enable(int64) (bool, error)
	Disable(int64) (bool, error)
}