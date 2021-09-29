package repositories

import (
	"go-blog-jwt-token/api/entities"
	"go-blog-jwt-token/api/payloads/request"
)

type PermissionRepository interface {
	FindPermission(int64, int64) (entities.Permission, error)
	CreatePermission([]request.PermissionRequest) (bool, error)
	UpdatePermission([]request.PermissionRequest) (bool, error)
	DeletePermission(int64) (bool, error)
}