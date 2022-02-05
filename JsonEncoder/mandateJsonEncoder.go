package JsonEncoder

import (
	"encoding/json"
	"net/http"

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
