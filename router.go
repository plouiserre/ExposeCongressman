package main

import (
	"github.com/gorilla/mux"
	controllers "github.com/plouiserre/exposecongressman/Controllers"
)

func InitializeRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/congressmans/", controllers.Congressmans).Methods("GET")
	router.HandleFunc("/congressmans/", controllers.CreateCongressman).Methods("POST")
	router.HandleFunc("/congressman/{id}", controllers.Congressman).Methods("GET")
	router.HandleFunc("/congressman/{id}", controllers.UpdateCongressman).Methods("PUT")
	router.HandleFunc("/congressman/{id}", controllers.DeleteCongressman).Methods("DELETE")

	return router
}
