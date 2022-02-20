package Controllers

import (
	"net/http"

	jsonEncoder "github.com/plouiserre/exposecongressman/JsonEncoder"
	"github.com/plouiserre/exposecongressman/Manager"
	models "github.com/plouiserre/exposecongressman/Models"
	repository "github.com/plouiserre/exposecongressman/Repository"
	services "github.com/plouiserre/exposecongressman/Services"
)

func Congressmans(w http.ResponseWriter, r *http.Request) {
	congressmanJsonEncoder := jsonEncoder.CongressmanJsonEncoder{
		W: w,
	}

	congressmans := models.CongressmansModel{}

	GetAll(congressmanJsonEncoder, r, congressmans)
}

func Congressman(w http.ResponseWriter, r *http.Request) {
	repo, _ := InitCongressmanRepository()

	congressmanJsonEncoder := jsonEncoder.CongressmanJsonEncoder{
		W: w,
	}

	congressman := models.CongressmanModel{}

	GetById(congressmanJsonEncoder, r, "congressman", *repo.LogManager, congressman)
}

func CreateCongressman(w http.ResponseWriter, r *http.Request) {
	repo, _ := InitCongressmanRepository()

	congressmanJsonEncoder := jsonEncoder.CongressmanJsonEncoder{
		W: w,
	}

	congressman := models.CongressmanModel{}

	CreateEntity(congressmanJsonEncoder, r, *repo.LogManager, congressman)
}

func UpdateCongressman(w http.ResponseWriter, r *http.Request) {
	repo, _ := InitCongressmanRepository()

	congressmanJsonEncoder := jsonEncoder.CongressmanJsonEncoder{
		W: w,
	}

	congressman := models.CongressmanModel{}

	UpdateEntity(congressmanJsonEncoder, r, repo, *repo.LogManager, congressman, congressman)
}

func DeleteCongressman(w http.ResponseWriter, r *http.Request) {
	repo, _ := InitCongressmanRepository()

	congressmanJsonEncoder := jsonEncoder.CongressmanJsonEncoder{
		W: w,
	}

	DeleteEntity(congressmanJsonEncoder, r, repo, *repo.LogManager)
}

func InitCongressmanRepository() (repository.CongressmanRepository, Manager.LogManager) {
	entityService := services.EntityService{}
	repo, logManager := entityService.InitRepository(0)
	congressmanRepo := repo.(repository.CongressmanRepository)
	return congressmanRepo, logManager
}
