package Service

import (
	models "github.com/plouiserre/exposecongressman/Models"
	repository "github.com/plouiserre/exposecongressman/Repository"
)

type CongressmanService struct {
	CongressmanRepository repository.CongressmanRepository
}

func (cs *CongressmanService) GetMandatesFromCongressman(id int) models.CongressmanMandatesModel {
	mandates := cs.CongressmanRepository.GetCongressmansMandates(id)
	congressmanMandates := models.CongressmanMandatesModel{
		CongressmanId: id,
		Mandates:      *mandates,
	}
	return congressmanMandates
}

func (cs *CongressmanService) GetCongressmansFromDepartment(id int) models.CongressmanDepartmentsModel {
	congressmanDepartments := cs.CongressmanRepository.GetCongressmansDepartments(id)
	return congressmanDepartments
}

func (cs *CongressmanService) GetCongressmansFromJobs(jobs string) models.CongressmansModel {
	congressmans := cs.CongressmanRepository.GetCongressmansJobs(jobs)
	return *congressmans
}
