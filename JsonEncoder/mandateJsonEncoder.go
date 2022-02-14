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

func (mj MandateJsonEncoder) EncodeEntities(entityModel models.EntityModel) {
	json.NewEncoder(mj.W).Encode(entityModel.Mandates)
}

func (mj MandateJsonEncoder) EncodeEntity(model models.EntityModel) {
	json.NewEncoder(mj.W).Encode(model.Mandate)
}

func (mj MandateJsonEncoder) WriteHeader(statusCode int) {
	mj.W.WriteHeader(statusCode)
}

func (mj MandateJsonEncoder) SetHeader() {
	mj.W.Header().Set("Content-type", "application/json;charset=UTF-8")
}

func (mj MandateJsonEncoder) UnmarshalEntity(body []byte, logManager Manager.LogManager) (models.EntityModel, bool) {
	var mandate models.MandateModel
	var entityResult models.EntityModel

	noError := true

	errJson := json.Unmarshal(body, &mandate)

	if errJson != nil {
		mj.WriteHeader(http.StatusBadRequest)
		logManager.WriteErrorLog(errJson.Error())
		noError = false
	}

	entityResult.Mandate = mandate
	return entityResult, noError
}

func (mj MandateJsonEncoder) ResponseEntityCreated(model models.EntityModel, lid int64) {
	mandate := model.Mandate
	mandate.Id = lid
	mj.WriteHeader(http.StatusCreated)
	//TODO à améliorer
	mj.EncodeEntity(model)
}
