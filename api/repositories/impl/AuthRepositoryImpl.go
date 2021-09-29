package impl

import (
	"errors"
	"go-blog-jwt-token/api/entities"
	"go-blog-jwt-token/api/payloads/request"
	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepositoryImpl(db *gorm.DB) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{db}
}

func (r *AuthRepositoryImpl) CheckUser(loginRequest request.LoginRequest) (entities.User, error) {
	var user entities.User

	r.db.Where("username=?", loginRequest.Username).Take(&user)

	if user.Username == "" {
		return entities.User{}, errors.New("Username tidak ditemukan")
	}

	if user.Status != true {
		return entities.User{}, errors.New("Account tidak aktif")
	}

	return user, nil
}

func (r *AuthRepositoryImpl) ForgetPassword(forgetRequest request.ForgetPasswordRequest, userId int64) (bool, error) {
	var user entities.User

	r.db.Where("id_user=?", userId).Take(&user)

	if user.ID == 0 {
		return false, errors.New("Id user tidak ditemukan")
	}

	row := r.db.Model(&user).Where("id_user=?", userId).Update("password", forgetRequest.NewPassword)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil
}