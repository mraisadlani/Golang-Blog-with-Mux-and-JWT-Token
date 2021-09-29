package routes

import (
	"go-blog-jwt-token/api/controllers"
	"net/http"
)

var roleRoute = []Route{
	Route{
		URI: "/RoleAPI/GetAllRoles",
		Method: http.MethodGet,
		Handler: controllers.GetAllRole,
		Auth: false,
	},
	Route{
		URI: "/RoleAPI/CreateRole",
		Method: http.MethodPost,
		Handler: controllers.CreateRole,
		Auth: false,
	},
	Route{
		URI: "/RoleAPI/{role_name}/FindByRoleName",
		Method: http.MethodPost,
		Handler: controllers.FindByRoleName,
		Auth: false,
	},
	Route{
		URI: "/RoleAPI/{id_role}/UpdateRole",
		Method: http.MethodPost,
		Handler: controllers.UpdateRole,
		Auth: false,
	},
	Route{
		URI: "/RoleAPI/{id_role}/DeleteRole",
		Method: http.MethodPost,
		Handler: controllers.DeleteRole,
		Auth: false,
	},
	Route{
		URI: "/RoleAPI/{id_role}/StatusRole",
		Method: http.MethodPost,
		Handler: controllers.StatusRole,
		Auth: false,
	},
}