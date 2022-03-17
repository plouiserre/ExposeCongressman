package Controllers

import (
	"encoding/json"
	"net/http"

	jsonEncoder "github.com/plouiserre/exposecongressman/JsonEncoder"
	"github.com/plouiserre/exposecongressman/Manager"
	models "github.com/plouiserre/exposecongressman/Models"
	services "github.com/plouiserre/exposecongressman/Services"
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

//TODO retravailler cette méthode pour la partie argument
func CongressmansMandates(w http.ResponseWriter, r *http.Request) {
	congressmanService := services.CongressmanService{}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	response := congressmanService.GetMandatesFromCongressman(5)
	json.NewEncoder(w).Encode(response)
}

//TODO
//1 - controller OK
//2 - service mock + controller branchement
//3 - repository branchement BDD + service branchement

func CongressmansByDepartment(w http.ResponseWriter, r *http.Request) {
	congressmanService := services.CongressmanService{}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	response := congressmanService.GetCongressmansFromDepartment(33)
	json.NewEncoder(w).Encode(response)
}
