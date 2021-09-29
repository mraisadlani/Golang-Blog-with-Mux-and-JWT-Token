package routes

import (
	"go-blog-jwt-token/api/controllers"
	"net/http"
)

var welcomeRoute = []Route{
	Route{
		URI: "/",
		Method: http.MethodGet,
		Handler: controllers.Welcome,
		Auth: false,
	},
}