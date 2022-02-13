package Controllers

import (
	"net/http"

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
