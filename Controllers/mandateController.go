package Controllers

import (
	"net/http"

	jsonEncoder "github.com/plouiserre/exposecongressman/JsonEncoder"
	"github.com/plouiserre/exposecongressman/Manager"
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

func Mandate(w http.ResponseWriter, r *http.Request) {
	repo, _ := InitMandateRepository()

	MandateJsonEncoder := jsonEncoder.MandateJsonEncoder{
		W: w,
	}

	GetById(MandateJsonEncoder, r, repo, "mandate", *repo.LogManager)
}

func CreateMandate(w http.ResponseWriter, r *http.Request) {
	repo, _ := InitMandateRepository()

	MandateJsonEncoder := jsonEncoder.MandateJsonEncoder{
		W: w,
	}
	CreateEntity(MandateJsonEncoder, r, repo, *repo.LogManager)
}

func UpdateMandate(w http.ResponseWriter, r *http.Request) {
	repo, _ := InitMandateRepository()

	MandateJsonEncoder := jsonEncoder.MandateJsonEncoder{
		W: w,
	}
	UpdateEntity(MandateJsonEncoder, r, repo, *repo.LogManager)
}

func DeleteMandate(w http.ResponseWriter, r *http.Request) {
	/*repo, logManager := InitMandateRepository()
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
	}*/
	repo, _ := InitMandateRepository()

	MandateJsonEncoder := jsonEncoder.MandateJsonEncoder{
		W: w,
	}
	DeleteEntity(MandateJsonEncoder, r, repo, *repo.LogManager)
}

func InitMandateRepository() (repository.MandateRepository, Manager.LogManager) {
	entityService := services.EntityService{}
	repo, logManager := entityService.InitRepository(1)
	mandateRepo := repo.(repository.MandateRepository)
	return mandateRepo, logManager
}
