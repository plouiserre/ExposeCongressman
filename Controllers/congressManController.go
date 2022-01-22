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

func Congressmans(w http.ResponseWriter, r *http.Request) {
	repo, _ := InitCongressmanRepository()
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	congressmans, noError := repo.AllCongressMans()

	if noError {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(congressmans)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func Congressman(w http.ResponseWriter, r *http.Request) {
	repo, logManager := InitCongressmanRepository()
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logManager.WriteErrorLog("Error cast " + err.Error())
	} else {
		congressman, noError := repo.GetCongressMan(id)
		if !noError {
			w.WriteHeader(http.StatusInternalServerError)
		} else if congressman != nil {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(congressman)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func CreateCongressman(w http.ResponseWriter, r *http.Request) {
	repo, logManager := InitCongressmanRepository()
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logManager.WriteErrorLog("Error Body " + err.Error())
	} else {
		var congressman models.CongressmanModel

		errJson := json.Unmarshal(body, &congressman)

		if errJson != nil {
			w.WriteHeader(http.StatusBadRequest)
			logManager.WriteErrorLog(err.Error())
		}

		lid, noError := repo.InsertCongressMan(&congressman)
		if !noError {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			congressman.Id = lid
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(congressman)
		}
	}
}

func UpdateCongressman(w http.ResponseWriter, r *http.Request) {
	repo, logManager := InitCongressmanRepository()
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logManager.WriteErrorLog("Error Body " + err.Error())
	} else {
		congressman, noError := repo.GetCongressMan(id)
		if !noError {
			w.WriteHeader(http.StatusInternalServerError)
		} else if congressman == nil {
			w.WriteHeader(http.StatusNotFound)
			logManager.WriteErrorLog("No congressman find with this id " + vars["id"])
		} else {
			body, errBody := ioutil.ReadAll(r.Body)
			if errBody != nil {
				w.WriteHeader(http.StatusBadRequest)
				logManager.WriteErrorLog(err.Error())
			} else {
				errJson := json.Unmarshal(body, &congressman)
				if errJson != nil {
					w.WriteHeader(http.StatusBadRequest)
					logManager.WriteErrorLog(err.Error())
				} else {
					noError := repo.UpdateCongressMan(congressman, id)
					if !noError {
						w.WriteHeader(http.StatusInternalServerError)
					} else {
						w.WriteHeader(http.StatusOK)
						json.NewEncoder(w).Encode(congressman)
					}
				}
			}
		}
	}
}

func DeleteCongressman(w http.ResponseWriter, r *http.Request) {
	repo, logManager := InitCongressmanRepository()
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logManager.WriteErrorLog("Error Body " + err.Error())
	} else {
		nbDelete, noError := repo.DeleteCongressMan(id)

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
func InitCongressmanRepository() (repository.CongressmanRepository, Manager.LogManager) {
	logManager := Manager.LogManager{}
	logManager.InitLog()
	congressmanRepository := repository.CongressmanRepository{
		LogManager: &logManager,
	}
	return congressmanRepository, logManager
}
