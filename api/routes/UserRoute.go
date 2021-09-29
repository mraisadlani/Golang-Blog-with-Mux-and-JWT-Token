package routes

import (
	"go-blog-jwt-token/api/controllers"
	"net/http"
)

var userRoute = []Route{
	Route{
		URI: "/UserAPI/GetAllUsers",
		Method: http.MethodGet,
		Handler: controllers.GetAllUser,
		Auth: false,
	},
	Route{
		URI: "/UserAPI/CreateUser",
		Method: http.MethodPost,
		Handler: controllers.CreateUser,
		Auth: false,
	},
	Route{
		URI: "/UserAPI/{id_user}/FindById",
		Method: http.MethodPost,
		Handler: controllers.FindById,
		Auth: false,
	},
	Route{
		URI: "/UserAPI/{id_user}/UpdateUser",
		Method: http.MethodPost,
		Handler: controllers.UpdateUser,
		Auth: false,
	},
	Route{
		URI: "/UserAPI/{id_user}/DeleteUser",
		Method: http.MethodPost,
		Handler: controllers.DeleteUser,
		Auth: false,
	},
	Route{
		URI: "/UserAPI/{id_user}/StatusUser",
		Method: http.MethodPost,
		Handler: controllers.StatusUser,
		Auth: false,
	},
	Route{
		URI: "/UserAPI/{id_user}/UploadImage",
		Method: http.MethodPost,
		Handler: controllers.UploadImage,
		Auth: false,
	},
}