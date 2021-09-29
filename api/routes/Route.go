package routes

import "github.com/gorilla/mux"

func SetupRoute() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	return SetupWithMiddleware(r)
}