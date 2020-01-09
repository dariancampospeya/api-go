package router

import (
	"api/router/routes"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter().StricSlash(true)
	return routes.SetupRoutes(r)
}
