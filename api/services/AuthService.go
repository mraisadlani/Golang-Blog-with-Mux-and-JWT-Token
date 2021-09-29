package services

import (
	"go-blog-jwt-token/api/entities"
	"go-blog-jwt-token/api/payloads/request"
	"go-blog-jwt-token/api/repositories/impl"
	"gorm.io/gorm"
)

func CheckLogin(db *gorm.DB, loginRequest request.LoginRequest) (entities.User, error) {
	repo := impl.NewAuthRepositoryImpl(db)

	get, err := repo.CheckUser(loginRequest)

	if err != nil {
		return entities.User{}, err
	}

	return get, nil
}

func ForgetPassword(db *gorm.DB, forget request.ForgetPasswordRequest, idUser int64) (bool, error) {
	repo := impl.NewAuthRepositoryImpl(db)

	get, err := repo.ForgetPassword(forget, idUser)

	if err != nil {
		return false, err
	}

	return get, nil
}