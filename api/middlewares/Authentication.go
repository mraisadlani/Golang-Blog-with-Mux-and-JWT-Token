package middlewares

import (
	"errors"
	"go-blog-jwt-token/api/payloads/response"
	"go-blog-jwt-token/api/securities"
	"net/http"
)

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		_, err := securities.Authorization(r)

		if err != nil {
			response.ResponseError(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		next(w, r)
	}
}