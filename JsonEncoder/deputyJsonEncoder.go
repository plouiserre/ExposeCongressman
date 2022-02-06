package JsonEncoder

import (
	"encoding/json"
	"net/http"

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
