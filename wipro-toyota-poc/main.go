package main

import (
	"log"
	"net/http"
	"wipro-toyota-poc/Dbconfig"
	"wipro-toyota-poc/Routers"

	"github.com/gorilla/mux"
)

func main() {
	// Create mux router
	router := mux.NewRouter()

	//run database
	Dbconfig.ConnectDB()

	//routes
	Routers.InitializeRouter(router) //add this

	// at 6000 port listeninggo
	log.Fatal(http.ListenAndServe(":6000", router))
}
