package JsonEncoder

import (
	"encoding/json"
	"net/http"

	Manager "github.com/plouiserre/exposecongressman/Manager"
	models "github.com/plouiserre/exposecongressman/Models"
)

type DeputyJsonEncoder struct {
	W http.ResponseWriter
}

func (dj DeputyJsonEncoder) EncodeEntities(entities models.IModels) {
	deputies := entities.(models.DeputiesModel)
	json.NewEncoder(dj.W).Encode(deputies)
}

func (dj DeputyJsonEncoder) EncodeEntity(entity models.IModel) {
	deputy := entity.(models.DeputyModel)
	json.NewEncoder(dj.W).Encode(deputy)
}

func (dj DeputyJsonEncoder) WriteHeader(statusCode int) {
	dj.W.WriteHeader(statusCode)
}

func (dj DeputyJsonEncoder) SetHeader() {
	dj.W.Header().Set("Content-type", "application/json;charset=UTF-8")
}

func (dj DeputyJsonEncoder) UnmarshalEntity(body []byte, logManager Manager.LogManager) (models.IModel, bool) {
	var deputy models.DeputyModel

	noError := true

	errJson := json.Unmarshal(body, &deputy)

	if errJson != nil {
		dj.WriteHeader(http.StatusBadRequest)
		logManager.WriteErrorLog(errJson.Error())
		noError = false
	}
	return deputy, noError
}

func (dj DeputyJsonEncoder) ResponseEntity(entity models.IModel, lid int64, statusCode int) {
	deputy := entity.(models.DeputyModel)
	deputy.Id = lid
	dj.WriteHeader(statusCode)
	dj.EncodeEntity(deputy)
}
