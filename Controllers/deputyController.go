package Controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/plouiserre/exposecongressman/Manager"
	models "github.com/plouiserre/exposecongressman/Models"
	repository "github.com/plouiserre/exposecongressman/Repository"
)

func Deputies(w http.ResponseWriter, r *http.Request) {
	repo, _ := InitDeputyRepository()
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	deputies, noError := repo.AllDeputies()

	if noError {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(deputies)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func Deputy(w http.ResponseWriter, r *http.Request) {
	repo, logManager := InitDeputyRepository()
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
	}
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

//TODO fixe this when you are multiple controllers
func InitDeputyRepository() (repository.DeputyRepository, Manager.LogManager) {
	logManager := Manager.LogManager{}
	logManager.InitLog()
	deputyRepository := repository.DeputyRepository{
		LogManager: &logManager,
	}
	return deputyRepository, logManager
}
