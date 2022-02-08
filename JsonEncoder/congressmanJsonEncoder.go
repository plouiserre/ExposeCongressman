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

func (cj CongressmanJsonEncoder) EncodeEntities(entityModel models.EntityModel) {
	json.NewEncoder(cj.W).Encode(entityModel.Congressmans)
}

func (cj CongressmanJsonEncoder) EncodeEntity(model models.EntityModel) {
	json.NewEncoder(cj.W).Encode(model.Congressman)
}

func (cj CongressmanJsonEncoder) WriteHeader(statusCode int) {
	cj.W.WriteHeader(statusCode)
}

func (cj CongressmanJsonEncoder) SetHeader() {
	cj.W.Header().Set("Content-type", "application/json;charset=UTF-8")
}

func (cj CongressmanJsonEncoder) UnmarshalEntity(body []byte, logManager Manager.LogManager) models.EntityModel {
	return models.EntityModel{}
}

func (cj CongressmanJsonEncoder) ResponseEntityCreated(model models.EntityModel, lid int64) {

}
