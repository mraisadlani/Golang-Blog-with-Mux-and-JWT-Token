package routes

import (
	"go-blog-jwt-token/api/controllers"
	"net/http"
)

var menuRoute = []Route{
	Route{
		URI: "/MenuAPI/",
		Method: http.MethodGet,
		Handler: controllers.GetAllMenu,
		Auth: false,
	},
	Route{
		URI: "/MenuAPI/{id_menu}/GetMenu",
		Method: http.MethodGet,
		Handler: controllers.GetMenu,
		Auth: false,
	},
	Route{
		URI: "/MenuAPI/CreateMenu",
		Method: http.MethodPost,
		Handler: controllers.CreateMenu,
		Auth: false,
	},
	Route{
		URI: "/MenuAPI/{id_menu}/UpdateMenu",
		Method: http.MethodPost,
		Handler: controllers.UpdateMenu,
		Auth: false,
	},
	Route{
		URI: "/MenuAPI/{id_menu}/DeleteMenu",
		Method: http.MethodPost,
		Handler: controllers.DeleteMenu,
		Auth: false,
	},
	Route{
		URI: "/MenuAPI/{id_menu}/StatusMenu",
		Method: http.MethodPost,
		Handler: controllers.StatusMenu,
		Auth: false,
	},
}