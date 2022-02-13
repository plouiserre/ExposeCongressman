package Controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	jsonEncoder "github.com/plouiserre/exposecongressman/JsonEncoder"
	"github.com/plouiserre/exposecongressman/Manager"
	repository "github.com/plouiserre/exposecongressman/Repository"
	services "github.com/plouiserre/exposecongressman/Services"
)

func Congressmans(w http.ResponseWriter, r *http.Request) {
	repo, _ := InitCongressmanRepository()

	congressmanJsonEncoder := jsonEncoder.CongressmanJsonEncoder{
		W: w,
	}

	GetAll(congressmanJsonEncoder, r, repo)
}

func Congressman(w http.ResponseWriter, r *http.Request) {
	repo, _ := InitCongressmanRepository()

	congressmanJsonEncoder := jsonEncoder.CongressmanJsonEncoder{
		W: w,
	}

	GetById(congressmanJsonEncoder, r, repo, "congressman", *repo.LogManager)
}

func CreateCongressman(w http.ResponseWriter, r *http.Request) {
	repo, _ := InitCongressmanRepository()

	congressmanJsonEncoder := jsonEncoder.CongressmanJsonEncoder{
		W: w,
	}

	CreateEntity(congressmanJsonEncoder, r, repo, *repo.LogManager)
}

func UpdateCongressman(w http.ResponseWriter, r *http.Request) {
	repo, _ := InitCongressmanRepository()

	congressmanJsonEncoder := jsonEncoder.CongressmanJsonEncoder{
		W: w,
	}

	UpdateEntity(congressmanJsonEncoder, r, repo, *repo.LogManager)
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

func InitCongressmanRepository() (repository.CongressmanRepository, Manager.LogManager) {
	entityService := services.EntityService{}
	repo, logManager := entityService.InitRepository(0)
	congressmanRepo := repo.(repository.CongressmanRepository)
	return congressmanRepo, logManager
}
