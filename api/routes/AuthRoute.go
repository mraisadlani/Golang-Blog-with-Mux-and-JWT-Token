package routes

import (
	"go-blog-jwt-token/api/controllers"
	"net/http"
)

var authRoute = []Route{
	Route{
		URI:     "/AuthAPI/DoLogin",
		Method:  http.MethodPost,
		Handler: controllers.DoLogin,
		Auth:    false,
	},
	Route{
		URI:     "/AuthAPI/{id_user}/ForgetPassword",
		Method:  http.MethodPost,
		Handler: controllers.ForgetPassword,
		Auth:    false,
	},
}