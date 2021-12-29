package main

import (
	"github.com/gorilla/mux"
	controllers "github.com/plouiserre/exposecongressman/Controllers"
)

func InitializeRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/congressmans").Name("Cars").HandlerFunc(controllers.CongressMans)
	router.Methods("GET").Path("/congressman/{id}").Name("Car").HandlerFunc(controllers.CongressMan)
	return router
}
