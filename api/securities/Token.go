package securities

import (
	"github.com/dgrijalva/jwt-go"
	"go-blog-jwt-token/api/configs"
	"time"
)

const AuthTokenValidTime = time.Hour * 3

func GenerateToken(username string) (string, error) {
	authToken, err := CreateAuthToken(username)

	if err != nil {
		return "", err
	}

	return authToken, nil
}

func CreateAuthToken(username string) (string, error) {
	authTokenExp := time.Now().Add(AuthTokenValidTime).Unix()

	claims := jwt.StandardClaims{
		Subject: username,
		ExpiresAt: authTokenExp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(configs.Config.SECRETKEY))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}