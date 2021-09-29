package routes

import (
	"go-blog-jwt-token/api/controllers"
	"net/http"
)

var permissionRoute = []Route{
	Route{
		URI: "/PermissionAPI/{id_user}/GetPermission/{id_menu}",
		Method: http.MethodGet,
		Handler: controllers.GetPermission,
		Auth: false,
	},
	Route{
		URI: "/PermissionAPI/CreatePermission",
		Method: http.MethodPost,
		Handler: controllers.CreatePermission,
		Auth: false,
	},
	Route{
		URI: "/PermissionAPI/UpdatePermission",
		Method: http.MethodPost,
		Handler: controllers.UpdatePermission,
		Auth: false,
	},
	Route{
		URI: "/PermissionAPI/{id_permission}/DeletePermission",
		Method: http.MethodPost,
		Handler: controllers.DeletePermission,
		Auth: false,
	},
}