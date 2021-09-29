package routes

import (
	"go-blog-jwt-token/api/controllers"
	"net/http"
)

var subMenuRoute = []Route{
	Route{
		URI: "/SubMenuAPI/",
		Method: http.MethodGet,
		Handler: controllers.GetAllSubMenu,
		Auth: false,
	},
	Route{
		URI: "/SubMenuAPI/{nama_menu}/GetSubMenu/{nama_sub_menu}",
		Method: http.MethodGet,
		Handler: controllers.GetSubMenu,
		Auth: false,
	},
	Route{
		URI: "/SubMenuAPI/{id_menu}/CreateSubMenu",
		Method: http.MethodPost,
		Handler: controllers.CreateSubMenu,
		Auth: false,
	},
	Route{
		URI: "/SubMenuAPI/{id_menu}/UpdateSubMenu/{id_sub_menu}",
		Method: http.MethodPost,
		Handler: controllers.UpdateSubMenu,
		Auth: false,
	},
	Route{
		URI: "/SubMenuAPI/{id_menu}/DeleteSubMenu/{id_sub_menu}",
		Method: http.MethodPost,
		Handler: controllers.DeleteSubMenu,
		Auth: false,
	},
	Route{
		URI: "/SubMenuAPI/{id_menu}/StatusSubMenu/{id_sub_menu}",
		Method: http.MethodPost,
		Handler: controllers.StatusSubMenu,
		Auth: false,
	},
}