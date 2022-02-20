package Controllers

import (
	"net/http"

	jsonEncoder "github.com/plouiserre/exposecongressman/JsonEncoder"
	"github.com/plouiserre/exposecongressman/Manager"
	models "github.com/plouiserre/exposecongressman/Models"
	repository "github.com/plouiserre/exposecongressman/Repository"
	services "github.com/plouiserre/exposecongressman/Services"
)

func Deputies(w http.ResponseWriter, r *http.Request) {
	deputyJsonEncoder := jsonEncoder.DeputyJsonEncoder{
		W: w,
	}

	deputies := models.DeputiesModel{}

	GetAll(deputyJsonEncoder, r, deputies)
}

func Deputy(w http.ResponseWriter, r *http.Request) {
	repo, _ := InitDeputyRepository()

	deputyJsonEncoder := jsonEncoder.DeputyJsonEncoder{
		W: w,
	}

	deputy := models.DeputyModel{}

	GetById(deputyJsonEncoder, r, "deputy", *repo.LogManager, deputy)
}

func CreateDeputy(w http.ResponseWriter, r *http.Request) {
	repo, _ := InitDeputyRepository()

	deputyJsonEncoder := jsonEncoder.DeputyJsonEncoder{
		W: w,
	}

	deputy := models.DeputyModel{}

	CreateEntity(deputyJsonEncoder, r, *repo.LogManager, deputy)
}

func UpdateDeputy(w http.ResponseWriter, r *http.Request) {
	repo, _ := InitDeputyRepository()

	deputyJsonEncoder := jsonEncoder.DeputyJsonEncoder{
		W: w,
	}

	deputy := models.DeputyModel{}

	UpdateEntity(deputyJsonEncoder, r, repo, *repo.LogManager, deputy, deputy)
}

func DeleteDeputy(w http.ResponseWriter, r *http.Request) {
	repo, _ := InitDeputyRepository()

	deputyJsonEncoder := jsonEncoder.DeputyJsonEncoder{
		W: w,
	}

	DeleteEntity(deputyJsonEncoder, r, repo, *repo.LogManager)
}

func InitDeputyRepository() (repository.DeputyRepository, Manager.LogManager) {
	entityService := services.EntityService{}
	repo, logManager := entityService.InitRepository(2)
	deputyRepo := repo.(repository.DeputyRepository)
	return deputyRepo, logManager
}
