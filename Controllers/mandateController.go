package Controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/plouiserre/exposecongressman/Manager"
	models "github.com/plouiserre/exposecongressman/Models"
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

//TODO voir comment factoriser entre les autres méthodes GET
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
	repo, logManager := InitMandateRepository()
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logManager.WriteErrorLog("Error Body " + err.Error())
	} else {
		var mandate models.MandateModel

		errJson := json.Unmarshal(body, &mandate)

		if errJson != nil {
			w.WriteHeader(http.StatusBadRequest)
			logManager.WriteErrorLog(err.Error())
		}

		lid, noError := repo.InsertMandate(&mandate)
		if !noError {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			mandate.Id = lid
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(mandate)
		}
	}
}

func UpdateMandate(w http.ResponseWriter, r *http.Request) {
	repo, logManager := InitMandateRepository()
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logManager.WriteErrorLog("Error Body " + err.Error())
	} else {
		mandate, noError := repo.GetMandate(id)
		if !noError {
			w.WriteHeader(http.StatusInternalServerError)
		} else if mandate == nil {
			w.WriteHeader(http.StatusNotFound)
			logManager.WriteErrorLog("No congressman find with this id " + vars["id"])
		} else {
			body, errBody := ioutil.ReadAll(r.Body)
			if errBody != nil {
				w.WriteHeader(http.StatusBadRequest)
				logManager.WriteErrorLog(err.Error())
			} else {
				errJson := json.Unmarshal(body, &mandate)
				if errJson != nil {
					w.WriteHeader(http.StatusBadRequest)
					logManager.WriteErrorLog(err.Error())
				} else {
					noError := repo.UpdateMandate(mandate, id)
					if !noError {
						w.WriteHeader(http.StatusInternalServerError)
					} else {
						w.WriteHeader(http.StatusOK)
						json.NewEncoder(w).Encode(mandate)
					}
				}
			}
		}
	}
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
