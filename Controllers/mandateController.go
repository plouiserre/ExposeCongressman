package Controllers

import (
	"net/http"

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

	mandates := models.MandatesModel{}

	GetAll(MandateJsonEncoder, r, repo, mandates)
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
