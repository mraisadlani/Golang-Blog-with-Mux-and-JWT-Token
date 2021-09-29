package services

import (
	"go-blog-jwt-token/api/entities"
	"go-blog-jwt-token/api/payloads/request"
	"go-blog-jwt-token/api/repositories/impl"
	"go-blog-jwt-token/api/utils"
	"gorm.io/gorm"
	"net/http"
)

func GetAllRole(r *http.Request, pagination *request.Pagination, db *gorm.DB) (interface{}, error) {
	repo := impl.NewRoleRepositoryImpl(db)
	set, err, totalPages := repo.GetAll(pagination)

	if err != nil {
		return nil, err
	}

	err = utils.SetupPagination(r, set, pagination, totalPages)

	if err != nil {
		return nil, err
	}

	return set, err
}

func InsertRole(db *gorm.DB, roleReq request.RoleRequest) (bool, error) {
	repo := impl.NewRoleRepositoryImpl(db)
	set, err := repo.Insert(roleReq)

	if err != nil {
		return false, err
	}

	return set, err
}

func FindByRoleName(db *gorm.DB, rolename string) (entities.Role, error) {
	repo := impl.NewRoleRepositoryImpl(db)
	set, err := repo.FindByNamaRole(rolename)

	if err != nil {
		return entities.Role{}, err
	}

	return set, err
}

func EditRole(db *gorm.DB, roleId int64, roleReq request.RoleRequest) (bool, error) {
	repo := impl.NewRoleRepositoryImpl(db)
	set, err := repo.Update(roleReq, roleId)

	if err != nil {
		return false, err
	}

	return set, err
}

func DeleteRole(db *gorm.DB, roleId int64) (bool, error) {
	repo := impl.NewRoleRepositoryImpl(db)
	set, err := repo.Delete(roleId)

	if err != nil {
		return false, err
	}

	return set, err
}

func FindByIdRole(db *gorm.DB, roleId int64) (entities.Role, error) {
	repo := impl.NewRoleRepositoryImpl(db)
	set, err := repo.FindById(roleId)

	if err != nil {
		return entities.Role{}, err
	}

	return set, err
}

func StatusRole(db *gorm.DB, get entities.Role) (bool, error) {
	var err error
	repo := impl.NewRoleRepositoryImpl(db)

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