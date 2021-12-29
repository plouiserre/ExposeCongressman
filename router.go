package main

import (
	"github.com/gorilla/mux"
	"github.com/plouiserre/exposecongressman/controllers"
)

func InitializeRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/congressmans").Name("Cars").HandlerFunc(controllers.CongressMans)
	return router
}
