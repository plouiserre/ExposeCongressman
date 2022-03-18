package Controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	jsonEncoder "github.com/plouiserre/exposecongressman/JsonEncoder"
	"github.com/plouiserre/exposecongressman/Manager"
	models "github.com/plouiserre/exposecongressman/Models"
	services "github.com/plouiserre/exposecongressman/Services"
)

//TODO mettre les logs

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

func CongressmansMandates(w http.ResponseWriter, r *http.Request) {
	id, noError := GetIdParameters(w, r)
	if noError {
		congressmanService := services.CongressmanService{}

		w.Header().Set("Content-type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		response := congressmanService.GetMandatesFromCongressman(id)
		json.NewEncoder(w).Encode(response)
	}
}

func CongressmansByDepartment(w http.ResponseWriter, r *http.Request) {
	id, noError := GetIdParameters(w, r)
	if noError {
		congressmanService := services.CongressmanService{}

		w.Header().Set("Content-type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		response := congressmanService.GetCongressmansFromDepartment(id)
		json.NewEncoder(w).Encode(response)
	}
}

func GetIdParameters(w http.ResponseWriter, r *http.Request) (int, bool) {
	logManager := InitLogManager()
	vars := mux.Vars(r)
	var noError bool
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logManager.WriteErrorLog("Error cast " + err.Error())
		noError = false
	} else {
		noError = true
	}
	return id, noError
}

func CongressmansByJobs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jobs := vars["jobs"]

	congressmanService := services.CongressmanService{}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	response := congressmanService.GetCongressmansFromJobs(jobs)
	json.NewEncoder(w).Encode(response)

}
