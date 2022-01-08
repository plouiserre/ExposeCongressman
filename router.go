package main

import (
	"github.com/gorilla/mux"
	controllers "github.com/plouiserre/exposecongressman/Controllers"
)

func InitializeRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/congressmans/", controllers.CongressMans).Methods("GET")
	router.HandleFunc("/congressmans/", controllers.CreateCongressMan).Methods("POST")
	router.HandleFunc("/congressman/{id}", controllers.CongressMan).Methods("GET")
	router.HandleFunc("/congressman/{id}", controllers.UpdateCongressMan).Methods("PUT")
	router.HandleFunc("/congressman/{id}", controllers.DeleteCongressMan).Methods("DELETE")

	return router
}
