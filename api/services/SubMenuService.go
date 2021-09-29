package services

import (
	"go-blog-jwt-token/api/entities"
	"go-blog-jwt-token/api/payloads/request"
	"go-blog-jwt-token/api/repositories/impl"
	"go-blog-jwt-token/api/utils"
	"gorm.io/gorm"
	"net/http"
)

func GetAllSubMenu(r *http.Request, db *gorm.DB, pagination *request.Pagination) (interface{}, error) {
	repo := impl.NewSubMenuRepositoryImpl(db)

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

func FindSubMenu(db *gorm.DB, namaMenu string, namaSubMenu string) (entities.SubMenu, error) {
	repo := impl.NewSubMenuRepositoryImpl(db)

	get, err := repo.FindBySubMenu(namaMenu, namaSubMenu)

	if err != nil {
		return entities.SubMenu{}, err
	}

	return get, nil
}

func CreateSubMenu(db *gorm.DB, idMenu int64, subMenus []request.SubMenuRequest) (bool, error) {
	repo := impl.NewSubMenuRepositoryImpl(db)

	get, err := repo.CreateSubMenu(idMenu, subMenus)

	if err != nil {
		return false, err
	}

	return get, nil
}

func UpdateSubMenu(db *gorm.DB, idMenu int64, idSubMenu int64, subMenus request.SubMenuRequest) (bool, error) {
	repo := impl.NewSubMenuRepositoryImpl(db)

	get, err := repo.UpdateSubMenu(idMenu, idSubMenu, subMenus)

	if err != nil {
		return false, err
	}

	return get, nil
}

func DeleteSubMenu(db *gorm.DB, idMenu int64, idSubMenu int64) (bool, error) {
	repo := impl.NewSubMenuRepositoryImpl(db)

	get, err := repo.DeleteSubMenu(idMenu, idSubMenu)

	if err != nil {
		return false, err
	}

	return get, nil
}

func FindByIdSubMenu(db *gorm.DB, idMenu int64, idSubMenu int64) (entities.SubMenu, error) {
	repo := impl.NewSubMenuRepositoryImpl(db)

	get, err := repo.FindById(idMenu, idSubMenu)

	if err != nil {
		return entities.SubMenu{}, err
	}

	return get, nil
}

func StatusSubMenu(db *gorm.DB, get entities.SubMenu) (bool, error) {
	var err error
	repo := impl.NewSubMenuRepositoryImpl(db)

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