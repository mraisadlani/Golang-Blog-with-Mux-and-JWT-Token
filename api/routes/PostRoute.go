package routes

import (
	"go-blog-jwt-token/api/controllers"
	"net/http"
)

var postRoute = []Route{
	Route{
		URI: "/PostAPI/GetAllPost",
		Method: http.MethodGet,
		Handler: controllers.GetAllPost,
		Auth: false,
	},
	Route{
		URI: "/PostAPI/CreatePost",
		Method: http.MethodPost,
		Handler: controllers.CreatePost,
		Auth: false,
	},
	Route{
		URI: "/PostAPI/{id_article}/FindPost/{id_post}",
		Method: http.MethodGet,
		Handler: controllers.FindPost,
		Auth: false,
	},
	Route{
		URI: "/PostAPI/UpdatePost",
		Method: http.MethodPost,
		Handler: controllers.UpdatePost,
		Auth: false,
	},
	Route{
		URI: "/PostAPI/{id_article}/DeletePost/{id_post}",
		Method: http.MethodPost,
		Handler: controllers.DeletePost,
		Auth: false,
	},
	Route{
		URI: "/PostAPI/{id_article}/PublishPost/{id_post}",
		Method: http.MethodPost,
		Handler: controllers.PublishPost,
		Auth: false,
	},
	Route{
		URI: "/PostAPI/{id_article}/CancelPost/{id_post}",
		Method: http.MethodPost,
		Handler: controllers.CancelPost,
		Auth: false,
	},
}