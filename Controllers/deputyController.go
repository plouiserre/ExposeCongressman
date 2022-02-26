package Controllers

import (
	"net/http"

	jsonEncoder "github.com/plouiserre/exposecongressman/JsonEncoder"
	models "github.com/plouiserre/exposecongressman/Models"
)

func Deputies(w http.ResponseWriter, r *http.Request) {
	deputyJsonEncoder := jsonEncoder.DeputyJsonEncoder{
		W: w,
	}

	deputies := models.DeputiesModel{}

	GetAll(deputyJsonEncoder, r, deputies)
}

func Deputy(w http.ResponseWriter, r *http.Request) {
	logManager := InitLogManager()

	deputyJsonEncoder := jsonEncoder.DeputyJsonEncoder{
		W: w,
	}

	deputy := models.DeputyModel{}

	GetById(deputyJsonEncoder, r, "deputy", logManager, deputy)
}

func CreateDeputy(w http.ResponseWriter, r *http.Request) {
	logManager := InitLogManager()

	deputyJsonEncoder := jsonEncoder.DeputyJsonEncoder{
		W: w,
	}

	deputy := models.DeputyModel{}

	CreateEntity(deputyJsonEncoder, r, logManager, deputy)
}

func UpdateDeputy(w http.ResponseWriter, r *http.Request) {
	logManager := InitLogManager()

	deputyJsonEncoder := jsonEncoder.DeputyJsonEncoder{
		W: w,
	}

	deputy := models.DeputyModel{}

	UpdateEntity(deputyJsonEncoder, r, logManager, deputy, deputy)
}

func DeleteDeputy(w http.ResponseWriter, r *http.Request) {
	logManager := InitLogManager()

	deputyJsonEncoder := jsonEncoder.DeputyJsonEncoder{
		W: w,
	}

	deputy := models.DeputyModel{}

	DeleteEntity(deputyJsonEncoder, r, logManager, deputy)
}
