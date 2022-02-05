package Controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	jsonEncoder "github.com/plouiserre/exposecongressman/JsonEncoder"
	"github.com/plouiserre/exposecongressman/Manager"
	models "github.com/plouiserre/exposecongressman/Models"
	repository "github.com/plouiserre/exposecongressman/Repository"
	services "github.com/plouiserre/exposecongressman/Services"
)

func Mandates(w http.ResponseWriter, r *http.Request) {
	repo, _ := InitMandateRepository()

	MandateJsonEncoder := jsonEncoder.MandateJsonEncoder{
		W: w,
	}

	GetAll(MandateJsonEncoder, r, repo)
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
	repo, logManager := InitMandateRepository()
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logManager.WriteErrorLog("Error Body " + err.Error())
	} else {
		nbDelete, noError := repo.DeleteMandate(id)

		if !noError {
			w.WriteHeader(http.StatusInternalServerError)
		} else if nbDelete > 0 {
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func InitMandateRepository() (repository.MandateRepository, Manager.LogManager) {
	entityService := services.EntityService{}
	repo, logManager := entityService.InitRepository(1)
	mandateRepo := repo.(repository.MandateRepository)
	return mandateRepo, logManager
}
