package Controllers

import (
	"net/http"

	jsonEncoder "github.com/plouiserre/exposecongressman/JsonEncoder"
	models "github.com/plouiserre/exposecongressman/Models"
)

func Mandates(w http.ResponseWriter, r *http.Request) {
	MandateJsonEncoder := jsonEncoder.MandateJsonEncoder{
		W: w,
	}

	mandates := models.MandatesModel{}

	GetAll(MandateJsonEncoder, r, mandates)
}

func Mandate(w http.ResponseWriter, r *http.Request) {
	logManager := InitLogManager()
	MandateJsonEncoder := jsonEncoder.MandateJsonEncoder{
		W: w,
	}

	mandate := models.MandateModel{}

	GetById(MandateJsonEncoder, r, "mandate", logManager, mandate)
}

func CreateMandate(w http.ResponseWriter, r *http.Request) {
	logManager := InitLogManager()

	MandateJsonEncoder := jsonEncoder.MandateJsonEncoder{
		W: w,
	}

	mandate := models.MandateModel{}

	CreateEntity(MandateJsonEncoder, r, logManager, mandate)
}

func UpdateMandate(w http.ResponseWriter, r *http.Request) {
	logManager := InitLogManager()

	MandateJsonEncoder := jsonEncoder.MandateJsonEncoder{
		W: w,
	}

	mandate := models.MandateModel{}

	//TODO changer ca car troooooop moche
	UpdateEntity(MandateJsonEncoder, r, logManager, mandate, mandate)
}

func DeleteMandate(w http.ResponseWriter, r *http.Request) {
	logManager := InitLogManager()

	MandateJsonEncoder := jsonEncoder.MandateJsonEncoder{
		W: w,
	}

	mandate := models.MandateModel{}

	DeleteEntity(MandateJsonEncoder, r, logManager, mandate)
}
