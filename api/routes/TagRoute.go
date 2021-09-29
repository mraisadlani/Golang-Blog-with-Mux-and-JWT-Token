package routes

import (
	"go-blog-jwt-token/api/controllers"
	"net/http"
)

var tagRoute = []Route{
	Route{
		URI: "/TagAPI/",
		Method: http.MethodGet,
		Handler: controllers.GetAllTag,
		Auth: false,
	},
	Route{
		URI: "/TagAPI/{id_tag}/GetTag",
		Method: http.MethodGet,
		Handler: controllers.GetTag,
		Auth: false,
	},
	Route{
		URI: "/TagAPI/CreateTag",
		Method: http.MethodPost,
		Handler: controllers.CreateTag,
		Auth: false,
	},
	Route{
		URI: "/TagAPI/{id_tag}/UpdateTag",
		Method: http.MethodPost,
		Handler: controllers.UpdateTag,
		Auth: false,
	},
	Route{
		URI: "/TagAPI/{id_tag}/DeleteTag",
		Method: http.MethodPost,
		Handler: controllers.DeleteTag,
		Auth: false,
	},
}