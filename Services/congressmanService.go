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

func (cs *CongressmanService) GetCongressmansFromDepartment(id int64) models.CongressmanDepartmentsModel {
	congressmanDepartments := models.CongressmanDepartmentsModel{
		{
			Id:            1,
			Uid:           "PA1008",
			Civility:      "M.",
			FirstName:     "Alain",
			LastName:      "David",
			Alpha:         "DavidA",
			BirthDate:     "1949-06-02 00:00:00",
			Region:        "Nouvelle-Aquitaine",
			Department:    "Gironde",
			Departmentnum: 33,
		},
		{
			Id:            100,
			Uid:           "PA606507",
			Civility:      "M.",
			FirstName:     "Florent",
			LastName:      "Boudié",
			Alpha:         "Boudie",
			BirthDate:     "1973-09-22 00:00:00",
			Region:        "Nouvelle-Aquitaine",
			Department:    "Gironde",
			Departmentnum: 33,
		},
		{
			Id:            246,
			Uid:           "PA719570",
			Civility:      "Mme",
			FirstName:     "Catherine",
			LastName:      "Fabre",
			Alpha:         "Fabre",
			BirthDate:     "1978-09-19 00:00:00",
			Region:        "Nouvelle-Aquitaine",
			Department:    "Gironde",
			Departmentnum: 33,
		},
	}
	return congressmanDepartments
}
