package JsonEncoder

import (
	Manager "github.com/plouiserre/exposecongressman/Manager"
	models "github.com/plouiserre/exposecongressman/Models"
)

type IJsonEncoder interface {
	WriteHeader(httpStatusCode int)
	EncodeEntities(models models.IModels)
	EncodeEntity(model models.IModel)
	SetHeader()
	UnmarshalEntity(body []byte, logManager Manager.LogManager) (models.IModel, bool)
	ResponseEntity(model models.IModel, lid int64, statusCode int)
}
