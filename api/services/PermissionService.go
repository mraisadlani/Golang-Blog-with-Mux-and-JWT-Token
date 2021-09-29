package services

import (
	"go-blog-jwt-token/api/entities"
	"go-blog-jwt-token/api/payloads/request"
	"go-blog-jwt-token/api/repositories/impl"
	"gorm.io/gorm"
)

func FindPermission(db *gorm.DB, IdUser int64, IdMenu int64) (entities.Permission, error) {
	repo := impl.NewPermissionRepositoryImpl(db)

	get, err := repo.FindPermission(IdUser, IdMenu)

	if err != nil {
		return entities.Permission{}, err
	}

	return get, nil
}

func CreatePermission(db *gorm.DB, permission []request.PermissionRequest) (bool, error) {
	repo := impl.NewPermissionRepositoryImpl(db)

	get, err := repo.CreatePermission(permission)

	if err != nil {
		return false, err
	}

	return get, nil
}

func UpdatePermission(db *gorm.DB, permission []request.PermissionRequest) (bool, error) {
	repo := impl.NewPermissionRepositoryImpl(db)

	get, err := repo.UpdatePermission(permission)

	if err != nil {
		return false, err
	}

	return get, nil
}

func DeletePermission(db *gorm.DB, IdPermission int64) (bool, error) {
	repo := impl.NewPermissionRepositoryImpl(db)

	get, err := repo.DeletePermission(IdPermission)

	if err != nil {
		return false, err
	}

	return get, nil
}