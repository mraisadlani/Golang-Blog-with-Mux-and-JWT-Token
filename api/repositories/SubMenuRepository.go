package repositories

import (
	"go-blog-jwt-token/api/entities"
	"go-blog-jwt-token/api/payloads/request"
)

type SubMenuRepository interface {
	GetAll(*request.Pagination) (interface{}, error, int)
	FindBySubMenu(string, string) (entities.SubMenu, error)
	CreateSubMenu(int64, []request.SubMenuRequest) (bool, error)
	UpdateSubMenu(int64, int64, request.SubMenuRequest) (bool, error)
	DeleteSubMenu(int64, int64) (bool, error)
	FindById(int64, int64) (entities.SubMenu, error)
	Enable(int64) (bool, error)
	Disable(int64) (bool, error)
}