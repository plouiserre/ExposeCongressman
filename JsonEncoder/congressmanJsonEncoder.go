package JsonEncoder

import (
	"encoding/json"
	"net/http"

	Manager "github.com/plouiserre/exposecongressman/Manager"
	models "github.com/plouiserre/exposecongressman/Models"
)

type CongressmanJsonEncoder struct {
	W http.ResponseWriter
}

func (cj CongressmanJsonEncoder) EncodeEntities(entities models.IModels) {
	congressmans := entities.(models.CongressmansModel)
	json.NewEncoder(cj.W).Encode(congressmans)
}

func (cj CongressmanJsonEncoder) EncodeEntity(entity models.IModel) {
	congressman := entity.(models.CongressmanModel)
	json.NewEncoder(cj.W).Encode(congressman)
}

func (cj CongressmanJsonEncoder) WriteHeader(statusCode int) {
	cj.W.WriteHeader(statusCode)
}

func (cj CongressmanJsonEncoder) SetHeader() {
	cj.W.Header().Set("Content-type", "application/json;charset=UTF-8")
}

func (cj CongressmanJsonEncoder) UnmarshalEntity(body []byte, logManager Manager.LogManager) (models.IModel, bool) {
	var congressman models.CongressmanModel

	noError := true

	errJson := json.Unmarshal(body, &congressman)

	if errJson != nil {
		cj.WriteHeader(http.StatusBadRequest)
		logManager.WriteErrorLog(errJson.Error())
		noError = false
	}
	return congressman, noError
}

func (cj CongressmanJsonEncoder) ResponseEntity(entity models.IModel, lid int64, statusCode int) {
	congressman := entity.(models.CongressmanModel)
	congressman.Id = lid
	cj.WriteHeader(statusCode)
	cj.EncodeEntity(congressman)
}
