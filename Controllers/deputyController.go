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

func Deputies(w http.ResponseWriter, r *http.Request) {
	repo, _ := InitDeputyRepository()

	deputyJsonEncoder := jsonEncoder.DeputyJsonEncoder{
		W: w,
	}

	GetAll(deputyJsonEncoder, r, repo)
}

func Deputy(w http.ResponseWriter, r *http.Request) {
	/*repo, logManager := InitDeputyRepository()
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logManager.WriteErrorLog("Error cast " + err.Error())
	} else {
		deputy, noError := repo.GetDeputy(id)
		if !noError {
			w.WriteHeader(http.StatusInternalServerError)
		} else if deputy != nil {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(deputy)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}*/
	repo, _ := InitDeputyRepository()

	deputyJsonEncoder := jsonEncoder.DeputyJsonEncoder{
		W: w,
	}

	GetById(deputyJsonEncoder, r, repo, *repo.LogManager)
}

func CreateDeputy(w http.ResponseWriter, r *http.Request) {
	repo, logManager := InitDeputyRepository()
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logManager.WriteErrorLog("Error Body " + err.Error())
	} else {
		var deputy models.DeputyModel
		errJson := json.Unmarshal(body, &deputy)

		if errJson != nil {
			w.WriteHeader(http.StatusBadRequest)
			logManager.WriteErrorLog(err.Error())
		}
		lid, noError := repo.InsertDeputy(&deputy)
		if !noError {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			deputy.Id = lid
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(deputy)
		}
	}
}

func UpdateDeputy(w http.ResponseWriter, r *http.Request) {
	repo, logManager := InitDeputyRepository()
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logManager.WriteErrorLog("Error Body " + err.Error())
	} else {
		deputy, noError := repo.GetDeputy(id)
		if !noError {
			w.WriteHeader(http.StatusInternalServerError)
		} else if deputy == nil {
			w.WriteHeader(http.StatusNotFound)
			logManager.WriteErrorLog("No Deputy find with this id " + vars["id"])
		} else {
			body, errBody := ioutil.ReadAll(r.Body)
			if errBody != nil {
				w.WriteHeader(http.StatusBadRequest)
				logManager.WriteErrorLog(err.Error())
			} else {
				errJson := json.Unmarshal(body, &deputy)
				if errJson != nil {
					w.WriteHeader(http.StatusBadRequest)
					logManager.WriteErrorLog(err.Error())
				} else {
					noError := repo.UpdateDeputy(deputy, id)
					if !noError {
						w.WriteHeader(http.StatusInternalServerError)
					} else {
						w.WriteHeader(http.StatusOK)
						json.NewEncoder(w).Encode(deputy)
					}
				}
			}
		}
	}
}

func DeleteDeputy(w http.ResponseWriter, r *http.Request) {
	repo, logManager := InitDeputyRepository()
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logManager.WriteErrorLog("Error Body " + err.Error())
	} else {
		nbDelete, noError := repo.DeleteDeputy(id)

		if !noError {
			w.WriteHeader(http.StatusInternalServerError)
		} else if nbDelete > 0 {
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func InitDeputyRepository() (repository.DeputyRepository, Manager.LogManager) {
	entityService := services.EntityService{}
	repo, logManager := entityService.InitRepository(2)
	deputyRepo := repo.(repository.DeputyRepository)
	return deputyRepo, logManager
}
