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

func (dj DeputyJsonEncoder) EncodeEntities(entityModel models.EntityModel) {
	json.NewEncoder(dj.W).Encode(entityModel.Deputies)
}

func (dj DeputyJsonEncoder) EncodeEntity(model models.EntityModel) {
	json.NewEncoder(dj.W).Encode(model.Deputy)
}

func (dj DeputyJsonEncoder) WriteHeader(statusCode int) {
	dj.W.WriteHeader(statusCode)
}

func (dj DeputyJsonEncoder) SetHeader() {
	dj.W.Header().Set("Content-type", "application/json;charset=UTF-8")
}

func (dj DeputyJsonEncoder) UnmarshalEntity(body []byte, logManager Manager.LogManager) (models.EntityModel, bool) {
	var deputy models.DeputyModel
	var entityResult models.EntityModel

	noError := true

	errJson := json.Unmarshal(body, &deputy)

	if errJson != nil {
		dj.WriteHeader(http.StatusBadRequest)
		logManager.WriteErrorLog(errJson.Error())
		noError = false
	}
	entityResult.Deputy = deputy
	return entityResult, noError
}

func (dj DeputyJsonEncoder) ResponseEntityCreated(model models.EntityModel, lid int64) {
	model.Deputy.Id = lid
	dj.WriteHeader(http.StatusCreated)
	dj.EncodeEntity(model)
}
