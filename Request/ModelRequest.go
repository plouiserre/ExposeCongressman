package Request

import (
	"net/http"

	jsonEncoder "github.com/plouiserre/exposecongressman/JsonEncoder"
	"github.com/plouiserre/exposecongressman/Manager"
	model "github.com/plouiserre/exposecongressman/Models"
)

type ModelRequest struct {
	JsonEncoder jsonEncoder.IJsonEncoder
	Request     *http.Request
	LogManager  Manager.LogManager
	Model       model.IModel
	Models      model.IModels
}
