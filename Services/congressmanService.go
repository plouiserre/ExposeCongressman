package Service

import (
	models "github.com/plouiserre/exposecongressman/Models"
	repository "github.com/plouiserre/exposecongressman/Repository"
)

type CongressmanService struct {
	CongressmanRepository repository.CongressmanRepository
}

func (cs *CongressmanService) GetMandatesFromCongressman(id int64) models.CongressmanMandatesModel {
	mandates := cs.CongressmanRepository.GetCongressmansMandates(id)
	congressmanMandates := models.CongressmanMandatesModel{
		CongressmanId: id,
		Mandates:      *mandates,
	}
	return congressmanMandates
}
