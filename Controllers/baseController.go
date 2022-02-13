package Controllers

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	jsonEncoder "github.com/plouiserre/exposecongressman/JsonEncoder"
	"github.com/plouiserre/exposecongressman/Manager"
	repository "github.com/plouiserre/exposecongressman/Repository"
	services "github.com/plouiserre/exposecongressman/Services"
)

//TODO mettre moins de paramètre dans ses méthodes

func GetAll(jsonEncoder jsonEncoder.IJsonEncoder, r *http.Request, repo repository.IRepository) {
	jsonEncoder.SetHeader()

	entityService := services.EntityService{}

	entities, noError := entityService.GetAll(repo)

	if noError {
		jsonEncoder.WriteHeader(http.StatusOK)
		jsonEncoder.EncodeEntities(*entities)
	} else {
		jsonEncoder.WriteHeader(http.StatusInternalServerError)
	}
}

//TODO optimiser les paramètres de la méthode et de entityService
func GetById(jsonEncoder jsonEncoder.IJsonEncoder, r *http.Request, repo repository.IRepository, entityName string, logManager Manager.LogManager) {
	jsonEncoder.SetHeader()

	entityService := services.EntityService{}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		jsonEncoder.WriteHeader(http.StatusBadRequest)
		logManager.WriteErrorLog("Error cast " + err.Error())
	} else {
		entity, noError := entityService.GetById(id, repo)
		if !noError {
			jsonEncoder.WriteHeader(http.StatusInternalServerError)
			logManager.WriteErrorLog("Error during the recovery of the entity")
		} else if entity != nil {
			jsonEncoder.WriteHeader(http.StatusOK)
			jsonEncoder.EncodeEntity(*entity)
		} else {
			badId := strconv.Itoa(id)
			jsonEncoder.WriteHeader(http.StatusNotFound)
			logManager.WriteErrorLog("No " + entityName + " with the Id " + badId)
		}
	}
}

func CreateEntity(jsonEncoder jsonEncoder.IJsonEncoder, r *http.Request, repo repository.IRepository, logManager Manager.LogManager) {
	jsonEncoder.SetHeader()

	entityService := services.EntityService{}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		jsonEncoder.WriteHeader(http.StatusBadRequest)
		logManager.WriteErrorLog("Error Body " + err.Error())
	} else {
		entity, noErrorMarhsal := jsonEncoder.UnmarshalEntity(body, logManager)
		if !noErrorMarhsal {
			jsonEncoder.WriteHeader(http.StatusBadRequest)
		} else {
			lid, noErrorCreation := entityService.CreateEntity(repo, &entity)
			if noErrorCreation {
				jsonEncoder.ResponseEntityCreated(entity, lid)
			}
		}
	}
}

func UpdateEntity(jsonEncoder jsonEncoder.IJsonEncoder, r *http.Request, repo repository.IRepository, logManager Manager.LogManager) {
	jsonEncoder.SetHeader()

	entityService := services.EntityService{}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		jsonEncoder.WriteHeader(http.StatusBadRequest)
		logManager.WriteErrorLog("Error Body " + err.Error())
	} else {
		entity, noError := entityService.GetById(id, repo)
		if !noError {
			jsonEncoder.WriteHeader(http.StatusInternalServerError)
		} else if entity == nil {
			jsonEncoder.WriteHeader(http.StatusNotFound)
			logManager.WriteErrorLog("No congressman find with this id " + vars["id"])
		} else {
			body, errBody := ioutil.ReadAll(r.Body)
			if errBody != nil {
				jsonEncoder.WriteHeader(http.StatusBadRequest)
				logManager.WriteErrorLog(err.Error())
			} else {
				entity, noErrorMarhsal := jsonEncoder.UnmarshalEntity(body, logManager)
				if noErrorMarhsal {
					noError := entityService.UpdateEntity(repo, &entity, id)
					if !noError {
						jsonEncoder.WriteHeader(http.StatusInternalServerError)
					} else {
						jsonEncoder.WriteHeader(http.StatusOK)
						jsonEncoder.EncodeEntity(entity)
					}
				}
			}
		}
	}
}
