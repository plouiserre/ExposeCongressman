package JsonEncoder

import (
	models "github.com/plouiserre/exposecongressman/Models"
)

type IJsonEncoder interface {
	WriteHeader(httpStatusCode int)
	EncodeEntities(models models.EntityModel)
	EncodeEntity(model models.EntityModel)
	SetHeader()
}
