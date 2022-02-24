package JsonEncoder

import (
	Manager "github.com/plouiserre/exposecongressman/Manager"
	models "github.com/plouiserre/exposecongressman/Models"
)

type IJsonEncoder interface {
	WriteHeader(httpStatusCode int)
	EncodeEntities(models models.EntityModel)
	EncodeEntity(model models.EntityModel)
	SetHeader()
	UnmarshalEntity(body []byte, logManager Manager.LogManager) (models.EntityModel, bool)
	ResponseEntity(model models.EntityModel, lid int64, statusCode int)
}
