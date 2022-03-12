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
	router.HandleFunc("/congressman/{id}/mandates", controllers.CongressmansMandates).Methods("GET")

	router.HandleFunc("/deputies/", controllers.Deputies).Methods("GET")
	router.HandleFunc("/deputies/", controllers.CreateDeputy).Methods("POST")
	router.HandleFunc("/deputy/{id}", controllers.Deputy).Methods("GET")
	router.HandleFunc("/deputy/{id}", controllers.UpdateDeputy).Methods("PUT")
	router.HandleFunc("/deputy/{id}", controllers.DeleteDeputy).Methods("DELETE")

	router.HandleFunc("/mandates/", controllers.Mandates).Methods("GET")
	router.HandleFunc("/mandates/", controllers.CreateMandate).Methods("POST")
	router.HandleFunc("/mandate/{id}", controllers.Mandate).Methods("GET")
	router.HandleFunc("/mandate/{id}", controllers.UpdateMandate).Methods("PUT")
	router.HandleFunc("/mandate/{id}", controllers.DeleteMandate).Methods("DELETE")

	return router
}
