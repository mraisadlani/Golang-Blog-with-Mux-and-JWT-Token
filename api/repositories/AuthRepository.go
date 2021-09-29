package repositories

import (
	"go-blog-jwt-token/api/entities"
	"go-blog-jwt-token/api/payloads/request"
)

type AuthRepository interface {
	CheckUser(request.LoginRequest) (entities.User, error)
	ForgetPassword(request.ForgetPasswordRequest, int64) (bool, error)
}