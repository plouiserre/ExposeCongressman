package Controllers

import (
	"net/http"

	jsonEncoder "github.com/plouiserre/exposecongressman/JsonEncoder"
	"github.com/plouiserre/exposecongressman/Manager"
	models "github.com/plouiserre/exposecongressman/Models"
)

func Deputies(w http.ResponseWriter, r *http.Request) {
	deputyJsonEncoder := jsonEncoder.DeputyJsonEncoder{
		W: w,
	}

	modelRequest := InitRequestModel(deputyJsonEncoder, r, Manager.LogManager{}, models.DeputyModel{}, models.DeputiesModel{})

	GetAll(modelRequest)
}

func Deputy(w http.ResponseWriter, r *http.Request) {
	logManager := InitLogManager()

	deputyJsonEncoder := jsonEncoder.DeputyJsonEncoder{
		W: w,
	}

	modelRequest := InitRequestModel(deputyJsonEncoder, r, logManager, models.DeputyModel{}, models.DeputiesModel{})

	GetById(modelRequest, "deputy")
}

func CreateDeputy(w http.ResponseWriter, r *http.Request) {
	logManager := InitLogManager()

	deputyJsonEncoder := jsonEncoder.DeputyJsonEncoder{
		W: w,
	}

	modelRequest := InitRequestModel(deputyJsonEncoder, r, logManager, models.DeputyModel{}, models.DeputiesModel{})

	CreateEntity(modelRequest)
}

func UpdateDeputy(w http.ResponseWriter, r *http.Request) {
	logManager := InitLogManager()

	deputyJsonEncoder := jsonEncoder.DeputyJsonEncoder{
		W: w,
	}

	modelRequest := InitRequestModel(deputyJsonEncoder, r, logManager, models.DeputyModel{}, models.DeputiesModel{})

	UpdateEntity(modelRequest)
}

func DeleteDeputy(w http.ResponseWriter, r *http.Request) {
	logManager := InitLogManager()

	deputyJsonEncoder := jsonEncoder.DeputyJsonEncoder{
		W: w,
	}

	modelRequest := InitRequestModel(deputyJsonEncoder, r, logManager, models.DeputyModel{}, models.DeputiesModel{})

	DeleteEntity(modelRequest)
}
