package Controllers

import (
	"net/http"

	jsonEncoder "github.com/plouiserre/exposecongressman/JsonEncoder"
	"github.com/plouiserre/exposecongressman/Manager"
	models "github.com/plouiserre/exposecongressman/Models"
)

func Congressmans(w http.ResponseWriter, r *http.Request) {
	congressmanJsonEncoder := jsonEncoder.CongressmanJsonEncoder{
		W: w,
	}

	modelRequest := InitRequestModel(congressmanJsonEncoder, r, Manager.LogManager{}, models.CongressmanModel{}, models.CongressmansModel{})

	GetAll(modelRequest)
}

func Congressman(w http.ResponseWriter, r *http.Request) {
	logManager := InitLogManager()

	congressmanJsonEncoder := jsonEncoder.CongressmanJsonEncoder{
		W: w,
	}

	modelRequest := InitRequestModel(congressmanJsonEncoder, r, logManager, models.CongressmanModel{}, models.CongressmansModel{})

	GetById(modelRequest, "congressman")
}

func CreateCongressman(w http.ResponseWriter, r *http.Request) {
	logManager := InitLogManager()

	congressmanJsonEncoder := jsonEncoder.CongressmanJsonEncoder{
		W: w,
	}

	modelRequest := InitRequestModel(congressmanJsonEncoder, r, logManager, models.CongressmanModel{}, models.CongressmansModel{})

	CreateEntity(modelRequest)
}

func UpdateCongressman(w http.ResponseWriter, r *http.Request) {
	logManager := InitLogManager()

	congressmanJsonEncoder := jsonEncoder.CongressmanJsonEncoder{
		W: w,
	}

	modelRequest := InitRequestModel(congressmanJsonEncoder, r, logManager, models.CongressmanModel{}, models.CongressmansModel{})

	UpdateEntity(modelRequest)
}

func DeleteCongressman(w http.ResponseWriter, r *http.Request) {
	logManager := InitLogManager()

	congressmanJsonEncoder := jsonEncoder.CongressmanJsonEncoder{
		W: w,
	}

	modelRequest := InitRequestModel(congressmanJsonEncoder, r, logManager, models.CongressmanModel{}, models.CongressmansModel{})

	DeleteEntity(modelRequest)
}
