package Controllers

import (
	"net/http"

	jsonEncoder "github.com/plouiserre/exposecongressman/JsonEncoder"
	repository "github.com/plouiserre/exposecongressman/Repository"
	services "github.com/plouiserre/exposecongressman/Services"
)

//TODO voir si on peut pas mettre le repo init une seule fois
func GetAll(jsonEncoder jsonEncoder.IJsonEncoder, r *http.Request, repo repository.IRepository) {
	jsonEncoder.SetHeader()

	entityService := services.EntityService{}

	entities, noError := entityService.GetAll(repo)

	if noError {
		jsonEncoder.WriteHeader(http.StatusOK)
		jsonEncoder.EncodeEntities(*entities)
	} else {
		jsonEncoder.WriteHeader(http.StatusInternalServerError)
	}
}
