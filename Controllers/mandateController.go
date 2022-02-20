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
	MandateJsonEncoder := jsonEncoder.MandateJsonEncoder{
		W: w,
	}

	mandates := models.MandatesModel{}

	GetAll(MandateJsonEncoder, r, mandates)
}

func Mandate(w http.ResponseWriter, r *http.Request) {
	repo, _ := InitMandateRepository()

	MandateJsonEncoder := jsonEncoder.MandateJsonEncoder{
		W: w,
	}

	mandate := models.MandateModel{}

	GetById(MandateJsonEncoder, r, "mandate", *repo.LogManager, mandate)
}

func CreateMandate(w http.ResponseWriter, r *http.Request) {
	repo, _ := InitMandateRepository()

	MandateJsonEncoder := jsonEncoder.MandateJsonEncoder{
		W: w,
	}

	mandate := models.MandateModel{}

	CreateEntity(MandateJsonEncoder, r, *repo.LogManager, mandate)
}

func UpdateMandate(w http.ResponseWriter, r *http.Request) {
	repo, _ := InitMandateRepository()

	MandateJsonEncoder := jsonEncoder.MandateJsonEncoder{
		W: w,
	}

	mandate := models.MandateModel{}

	//TODO changer ca car troooooop moche
	UpdateEntity(MandateJsonEncoder, r, repo, *repo.LogManager, mandate, mandate)
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
