package Controllers

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	jsonEncoder "github.com/plouiserre/exposecongressman/JsonEncoder"
	"github.com/plouiserre/exposecongressman/Manager"
	model "github.com/plouiserre/exposecongressman/Models"
	repository "github.com/plouiserre/exposecongressman/Repository"
	modelRequest "github.com/plouiserre/exposecongressman/Request"
	services "github.com/plouiserre/exposecongressman/Services"
)

func GetAll(request modelRequest.ModelRequest) {
	repositoryBase := InitBaseController(request.JsonEncoder)

	entityService := services.EntityService{
		Models:         request.Models,
		RepositoryBase: repositoryBase,
	}

	entities, noError := entityService.GetAll()

	if noError {
		request.JsonEncoder.WriteHeader(http.StatusOK)
		request.JsonEncoder.EncodeEntities(*entities)
	} else {
		request.JsonEncoder.WriteHeader(http.StatusInternalServerError)
	}
}

func GetById(request modelRequest.ModelRequest, entityName string) {
	repositoryBase := InitBaseController(request.JsonEncoder)

	entityService := services.EntityService{
		IModel:         request.Model,
		RepositoryBase: repositoryBase,
	}

	vars := mux.Vars(request.Request)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		request.JsonEncoder.WriteHeader(http.StatusBadRequest)
		request.LogManager.WriteErrorLog("Error cast " + err.Error())
	} else {
		entity, noError := entityService.GetById(id)
		if !noError {
			request.JsonEncoder.WriteHeader(http.StatusInternalServerError)
			request.LogManager.WriteErrorLog("Error during the recovery of the entity")
		} else if entity != nil {
			request.JsonEncoder.WriteHeader(http.StatusOK)
			request.JsonEncoder.EncodeEntity(*entity)
		} else {
			badId := strconv.Itoa(id)
			request.JsonEncoder.WriteHeader(http.StatusNotFound)
			request.LogManager.WriteErrorLog("No " + entityName + " with the Id " + badId)
		}
	}
}

func CreateEntity(request modelRequest.ModelRequest) {
	repositoryBase := InitBaseController(request.JsonEncoder)

	entityService := services.EntityService{
		IModel:         request.Model,
		RepositoryBase: repositoryBase,
	}

	body, err := ioutil.ReadAll(request.Request.Body)

	if err != nil {
		request.JsonEncoder.WriteHeader(http.StatusBadRequest)
		request.LogManager.WriteErrorLog("Error Body " + err.Error())
	} else {
		entity, noErrorMarhsal := request.JsonEncoder.UnmarshalEntity(body, request.LogManager)
		if !noErrorMarhsal {
			request.JsonEncoder.WriteHeader(http.StatusBadRequest)
		} else {
			lid, noErrorCreation := entityService.CreateEntity(&entity)
			if noErrorCreation {
				request.JsonEncoder.ResponseEntity(entity, lid, http.StatusCreated)
			}
		}
	}
}

func UpdateEntity(request modelRequest.ModelRequest) {
	repositoryBase := InitBaseController(request.JsonEncoder)

	entityService := services.EntityService{
		IModel:         request.Model,
		RepositoryBase: repositoryBase,
	}

	vars := mux.Vars(request.Request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		request.JsonEncoder.WriteHeader(http.StatusBadRequest)
		request.LogManager.WriteErrorLog("Error Body " + err.Error())
	} else {
		entity, noError := entityService.GetById(id)
		if !noError {
			request.JsonEncoder.WriteHeader(http.StatusInternalServerError)
		} else if entity == nil {
			request.JsonEncoder.WriteHeader(http.StatusNotFound)
			request.LogManager.WriteErrorLog("No congressman find with this id " + vars["id"])
		} else {
			body, errBody := ioutil.ReadAll(request.Request.Body)
			if errBody != nil {
				request.JsonEncoder.WriteHeader(http.StatusBadRequest)
				request.LogManager.WriteErrorLog(err.Error())
			} else {
				entity, noErrorMarhsal := request.JsonEncoder.UnmarshalEntity(body, request.LogManager)
				if noErrorMarhsal {
					updateId := int64(id)
					updateId, noError := entityService.UpdateEntity(&entity, updateId)
					if !noError {
						request.JsonEncoder.WriteHeader(http.StatusInternalServerError)
					} else {
						request.JsonEncoder.ResponseEntity(entity, updateId, http.StatusOK)
					}
				}
			}
		}
	}
}

func DeleteEntity(request modelRequest.ModelRequest) {
	repositoryBase := InitBaseController(request.JsonEncoder)

	entityService := services.EntityService{
		IModel:         request.Model,
		RepositoryBase: repositoryBase,
	}

	vars := mux.Vars(request.Request)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		request.JsonEncoder.WriteHeader(http.StatusBadRequest)
		request.LogManager.WriteErrorLog("Error Body " + err.Error())
	} else {
		nbDelete, noError := entityService.DeleteEntity(id)

		if !noError {
			request.JsonEncoder.WriteHeader(http.StatusInternalServerError)
		} else if nbDelete > 0 {
			request.JsonEncoder.WriteHeader(http.StatusNoContent)
		} else {
			request.JsonEncoder.WriteHeader(http.StatusNotFound)
		}
	}
}

func InitBaseController(jsonEncoder jsonEncoder.IJsonEncoder) repository.RepositoryBase {
	jsonEncoder.SetHeader()

	logManager := Manager.LogManager{}
	logManager.InitLog()

	repositoryBase := repository.RepositoryBase{
		LogManager: &logManager,
	}

	return repositoryBase
}

func InitLogManager() Manager.LogManager {
	entityService := services.EntityService{}
	logManager := entityService.InitLogManager()
	return logManager
}

func InitRequestModel(jsonEncoder jsonEncoder.IJsonEncoder, r *http.Request, logManager Manager.LogManager, model model.IModel, models model.IModels) modelRequest.ModelRequest {
	modelRequest := modelRequest.ModelRequest{
		JsonEncoder: jsonEncoder,
		Request:     r,
		LogManager:  logManager,
		Model:       model,
		Models:      models,
	}
	return modelRequest
}
