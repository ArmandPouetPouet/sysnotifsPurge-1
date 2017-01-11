package routes

import (
	"account-api-user/handlers"
	"account-api-user/structs"
)

//AccountRoutes ...
var AccountRoutes = structs.Routes{
	// GET
	// get user information
	structs.Route{
		Name:        "CheckRatio",
		Method:      "GET",
		Pattern:     "/CheckRatio",
		HandlerFunc: handlers.CheckRatio,
	},
}
