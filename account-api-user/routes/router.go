package routes

import (
	"account-api-user/structs"

	"github.com/gorilla/mux"
)

var routes = make([]structs.Route, 2, 4)

//NewRouter ...
func NewRouter() *mux.Router {

	routes = append(routes, AccountRoutes...)

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
