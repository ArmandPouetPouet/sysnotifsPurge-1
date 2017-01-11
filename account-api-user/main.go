package main

import (
	"account-api-user/routes"
	"log"
	"net/http"
)

func main() {

	//Setup the API routes
	router := routes.NewRouter()

	log.Fatal(http.ListenAndServe(":5000", router))
}
