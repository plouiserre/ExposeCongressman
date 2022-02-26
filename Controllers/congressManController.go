package Controllers

import (
	"net/http"

	jsonEncoder "github.com/plouiserre/exposecongressman/JsonEncoder"
	models "github.com/plouiserre/exposecongressman/Models"
)

func Congressmans(w http.ResponseWriter, r *http.Request) {
	congressmanJsonEncoder := jsonEncoder.CongressmanJsonEncoder{
		W: w,
	}

	congressmans := models.CongressmansModel{}

	GetAll(congressmanJsonEncoder, r, congressmans)
}

func Congressman(w http.ResponseWriter, r *http.Request) {
	logManager := InitLogManager()

	congressmanJsonEncoder := jsonEncoder.CongressmanJsonEncoder{
		W: w,
	}

	congressman := models.CongressmanModel{}

	GetById(congressmanJsonEncoder, r, "congressman", logManager, congressman)
}

func CreateCongressman(w http.ResponseWriter, r *http.Request) {
	logManager := InitLogManager()

	congressmanJsonEncoder := jsonEncoder.CongressmanJsonEncoder{
		W: w,
	}

	congressman := models.CongressmanModel{}

	CreateEntity(congressmanJsonEncoder, r, logManager, congressman)
}

func UpdateCongressman(w http.ResponseWriter, r *http.Request) {
	logManager := InitLogManager()

	congressmanJsonEncoder := jsonEncoder.CongressmanJsonEncoder{
		W: w,
	}

	congressman := models.CongressmanModel{}

	UpdateEntity(congressmanJsonEncoder, r, logManager, congressman)
}

func DeleteCongressman(w http.ResponseWriter, r *http.Request) {
	logManager := InitLogManager()

	congressmanJsonEncoder := jsonEncoder.CongressmanJsonEncoder{
		W: w,
	}

	congressman := models.CongressmanModel{}

	DeleteEntity(congressmanJsonEncoder, r, logManager, congressman)
}
