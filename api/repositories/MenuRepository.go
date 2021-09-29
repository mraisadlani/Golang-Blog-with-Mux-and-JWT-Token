package repositories

import (
	"go-blog-jwt-token/api/entities"
	"go-blog-jwt-token/api/payloads/request"
)

type MenuRepository interface {
	GetAll(*request.Pagination) (interface{}, error, int)
	FindByMenu(int64) (entities.Menu, error)
	CreateMenu(request.MenuRequest) (bool, error)
	UpdateMenu(request.MenuRequest, int64) (bool, error)
	DeleteMenu(int64) (bool, error)
	Enable(int64) (bool, error)
	Disable(int64) (bool, error)
}