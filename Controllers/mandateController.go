package Controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/plouiserre/exposecongressman/Manager"
	repository "github.com/plouiserre/exposecongressman/Repository"
)

func Mandates(w http.ResponseWriter, r *http.Request) {
	repo, _ := InitMandateRepository()
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	mandates, noError := repo.AllMandates()

	if noError {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mandates)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func Mandate(w http.ResponseWriter, r *http.Request) {
	repo, logManager := InitMandateRepository()
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logManager.WriteErrorLog("Error cast " + err.Error())
	} else {
		mandate, noError := repo.GetMandate(id)
		if !noError {
			w.WriteHeader(http.StatusInternalServerError)
		} else if mandate != nil {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(mandate)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func CreateMandate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	fmt.Println("CreateMandate called")
}

func UpdateMandate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	fmt.Println("UpdateMandate called")
}

func DeleteMandate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	fmt.Println("DeleteMandate called")
}

//TODO fixe this when you are multiple controllers
func InitMandateRepository() (repository.MandateRepository, Manager.LogManager) {
	logManager := Manager.LogManager{}
	logManager.InitLog()
	mandateRepository := repository.MandateRepository{
		LogManager: &logManager,
	}
	return mandateRepository, logManager
}
