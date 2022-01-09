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

	router.HandleFunc("/deputies/", controllers.Deputies).Methods("GET")
	router.HandleFunc("/deputies/", controllers.CreateDeputy).Methods("POST")
	router.HandleFunc("/deputy/{id}", controllers.Deputy).Methods("GET")
	router.HandleFunc("/deputy/{id}", controllers.UpdateDeputy).Methods("PUT")
	router.HandleFunc("/deputy/{id}", controllers.DeleteDeputy).Methods("DELETE")

	return router
}
