package routes

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"go-blog-jwt-token/api/middlewares"
	_ "go-blog-jwt-token/docs"
	"net/http"
)

type Route struct {
	URI     string
	Method  string
	Handler func(w http.ResponseWriter, r *http.Request)
	Auth    bool
}

func Load() []Route {
	routes := welcomeRoute
	routes = append(routes, userRoute...)
	routes = append(routes, authRoute...)
	routes = append(routes, permissionRoute...)
	routes = append(routes, menuRoute...)
	routes = append(routes, subMenuRoute...)
	routes = append(routes, roleRoute...)
	routes = append(routes, articleRoute...)
	routes = append(routes, postRoute...)
	routes = append(routes, tagRoute...)
	routes = append(routes, categoryRoute...)

	return routes
}

func SetupWithMiddleware(r *mux.Router) *mux.Router {
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	for _, route := range Load() {
		api := r.PathPrefix("/api").Subrouter()

		if route.Auth {
			api.HandleFunc(route.URI,
				middlewares.SetMiddlewareCors(
					middlewares.SetMiddlewareAuthentication(
						route.Handler,
					),
				),
			).Methods(route.Method)
		} else {
			api.HandleFunc(route.URI,
				middlewares.SetMiddlewareCors(
					route.Handler,
				),
			).Methods(route.Method)
		}
	}

	return r
}