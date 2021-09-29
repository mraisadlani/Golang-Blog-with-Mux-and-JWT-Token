package services

import (
	"go-blog-jwt-token/api/entities"
	"go-blog-jwt-token/api/payloads/request"
	"go-blog-jwt-token/api/repositories/impl"
	"go-blog-jwt-token/api/utils"
	"gorm.io/gorm"
	"net/http"
)

func GetAllMenu(r *http.Request, db *gorm.DB, pagination *request.Pagination) (interface{}, error) {
	repo := impl.NewMenuRepositoryImpl(db)

	get, err, totalPages := repo.GetAll(pagination)

	if err != nil {
		return nil, err
	}

	err = utils.SetupPagination(r, get, pagination, totalPages)

	if err != nil {
		return nil, err
	}

	return get, nil
}

func FindMenu(db *gorm.DB, IdMenu int64) (entities.Menu, error) {
	repo := impl.NewMenuRepositoryImpl(db)

	get, err := repo.FindByMenu(IdMenu)

	if err != nil {
		return entities.Menu{}, err
	}

	return get, nil
}

func CreateMenu(db *gorm.DB, menus request.MenuRequest) (bool, error) {
	repo := impl.NewMenuRepositoryImpl(db)

	get, err := repo.CreateMenu(menus)

	if err != nil {
		return false, err
	}

	return get, nil
}

func UpdateMenu(db *gorm.DB, menus request.MenuRequest, IdMenu int64) (bool, error) {
	repo := impl.NewMenuRepositoryImpl(db)

	get, err := repo.UpdateMenu(menus, IdMenu)

	if err != nil {
		return false, err
	}

	return get, nil
}

func DeleteMenu(db *gorm.DB, IdMenu int64) (bool, error) {
	repo := impl.NewMenuRepositoryImpl(db)

	get, err := repo.DeleteMenu(IdMenu)

	if err != nil {
		return false, err
	}

	return get, nil
}

func StatusMenu(db *gorm.DB, get entities.Menu) (bool, error) {
	var err error
	repo := impl.NewMenuRepositoryImpl(db)

	if get.Status == false {
		_, err = repo.Enable(get.ID)
	} else {
		_, err = repo.Disable(get.ID)
	}

	if err != nil {
		return false, err
	}

	return true, err
}