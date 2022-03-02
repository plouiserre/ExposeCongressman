package JsonEncoder

import (
	"encoding/json"
	"net/http"

	Manager "github.com/plouiserre/exposecongressman/Manager"
	models "github.com/plouiserre/exposecongressman/Models"
)

type MandateJsonEncoder struct {
	W http.ResponseWriter
}

func (mj MandateJsonEncoder) EncodeEntities(entities models.IModels) {
	mandates := entities.(models.MandatesModel)
	json.NewEncoder(mj.W).Encode(mandates)
}

func (mj MandateJsonEncoder) EncodeEntity(entity models.IModel) {
	mandate := entity.(models.MandateModel)
	json.NewEncoder(mj.W).Encode(mandate)
}

func (mj MandateJsonEncoder) WriteHeader(statusCode int) {
	mj.W.WriteHeader(statusCode)
}

func (mj MandateJsonEncoder) SetHeader() {
	mj.W.Header().Set("Content-type", "application/json;charset=UTF-8")
}

func (mj MandateJsonEncoder) UnmarshalEntity(body []byte, logManager Manager.LogManager) (models.IModel, bool) {
	var mandate models.MandateModel

	noError := true

	errJson := json.Unmarshal(body, &mandate)

	if errJson != nil {
		mj.WriteHeader(http.StatusBadRequest)
		logManager.WriteErrorLog(errJson.Error())
		noError = false
	}

	return mandate, noError
}

func (mj MandateJsonEncoder) ResponseEntity(entity models.IModel, lid int64, statusCode int) {
	mandate := entity.(models.MandateModel)
	mandate.Id = lid
	mj.WriteHeader(statusCode)
	mj.EncodeEntity(mandate)
}
