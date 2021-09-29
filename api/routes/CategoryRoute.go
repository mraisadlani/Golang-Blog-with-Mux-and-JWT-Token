package routes

import (
	"go-blog-jwt-token/api/controllers"
	"net/http"
)

var categoryRoute = []Route{
	Route{
		URI: "/CategoryAPI/",
		Method: http.MethodGet,
		Handler: controllers.GetAllCategory,
		Auth: false,
	},
	Route{
		URI: "/CategoryAPI/{id_category}/GetCategory",
		Method: http.MethodGet,
		Handler: controllers.GetCategory,
		Auth: false,
	},
	Route{
		URI: "/CategoryAPI/CreateCategory",
		Method: http.MethodPost,
		Handler: controllers.CreateCategory,
		Auth: false,
	},
	Route{
		URI: "/CategoryAPI/{id_category}/UpdateCategory",
		Method: http.MethodPost,
		Handler: controllers.UpdateCategory,
		Auth: false,
	},
	Route{
		URI: "/CategoryAPI/{id_category}/DeleteCategory",
		Method: http.MethodPost,
		Handler: controllers.DeleteCategory,
		Auth: false,
	},
}