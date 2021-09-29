package services

import (
	"go-blog-jwt-token/api/entities"
	"go-blog-jwt-token/api/payloads/request"
	"go-blog-jwt-token/api/repositories/impl"
	"go-blog-jwt-token/api/utils"
	"gorm.io/gorm"
	"net/http"
)

func GetAllUser(r *http.Request, pagination *request.Pagination, db *gorm.DB) (interface{}, error) {
	repo := impl.NewUserRepositoryImpl(db)
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

func InsertUser(db *gorm.DB, userReq request.UserRequest) (bool, error) {
	repo := impl.NewUserRepositoryImpl(db)
	set, err := repo.Insert(userReq)

	if err != nil {
		return false, err
	}

	return set, err
}

func FindById(db *gorm.DB, userId int64) (entities.User, error) {
	repo := impl.NewUserRepositoryImpl(db)
	set, err := repo.FindById(userId)

	if err != nil {
		return entities.User{}, err
	}

	return set, err
}

func EditUser(db *gorm.DB, userId int64, userReq request.UserRequest) (bool, error) {
	repo := impl.NewUserRepositoryImpl(db)
	set, err := repo.Update(userReq, userId)

	if err != nil {
		return false, err
	}

	return set, err
}

func DeleteUser(db *gorm.DB, userId int64) (bool, error) {
	repo := impl.NewUserRepositoryImpl(db)
	set, err := repo.Delete(userId)

	if err != nil {
		return false, err
	}

	return set, err
}

func StatusUser(db *gorm.DB, get entities.User) (bool, error) {
	var err error
	repo := impl.NewUserRepositoryImpl(db)

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

func UploadImage(db *gorm.DB, filename string, userId int64) (bool, error) {
	repo := impl.NewUserRepositoryImpl(db)

	get, err := repo.UploadImage(filename, userId)

	if err != nil {
		return false, err
	}

	return get, nil
}