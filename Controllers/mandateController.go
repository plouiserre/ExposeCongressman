package Controllers

import (
	"net/http"

	jsonEncoder "github.com/plouiserre/exposecongressman/JsonEncoder"
	"github.com/plouiserre/exposecongressman/Manager"
	models "github.com/plouiserre/exposecongressman/Models"
)

func Mandates(w http.ResponseWriter, r *http.Request) {
	MandateJsonEncoder := jsonEncoder.MandateJsonEncoder{
		W: w,
	}

	modelRequest := InitRequestModel(MandateJsonEncoder, r, Manager.LogManager{}, models.MandateModel{}, models.MandatesModel{})

	GetAll(modelRequest)
}

func Mandate(w http.ResponseWriter, r *http.Request) {
	logManager := InitLogManager()
	MandateJsonEncoder := jsonEncoder.MandateJsonEncoder{
		W: w,
	}

	modelRequest := InitRequestModel(MandateJsonEncoder, r, logManager, models.MandateModel{}, models.MandatesModel{})

	GetById(modelRequest, "mandate")
}

func CreateMandate(w http.ResponseWriter, r *http.Request) {
	logManager := InitLogManager()

	MandateJsonEncoder := jsonEncoder.MandateJsonEncoder{
		W: w,
	}

	modelRequest := InitRequestModel(MandateJsonEncoder, r, logManager, models.MandateModel{}, models.MandatesModel{})

	CreateEntity(modelRequest)
}

func UpdateMandate(w http.ResponseWriter, r *http.Request) {
	logManager := InitLogManager()

	MandateJsonEncoder := jsonEncoder.MandateJsonEncoder{
		W: w,
	}

	modelRequest := InitRequestModel(MandateJsonEncoder, r, logManager, models.MandateModel{}, models.MandatesModel{})

	UpdateEntity(modelRequest)
}

func DeleteMandate(w http.ResponseWriter, r *http.Request) {
	logManager := InitLogManager()

	MandateJsonEncoder := jsonEncoder.MandateJsonEncoder{
		W: w,
	}

	modelRequest := InitRequestModel(MandateJsonEncoder, r, logManager, models.MandateModel{}, models.MandatesModel{})

	DeleteEntity(modelRequest)
}
