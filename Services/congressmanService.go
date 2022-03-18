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
	/*congressmanDepartments := models.CongressmanDepartmentsModel{
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
	}*/
	congressmanDepartments := cs.CongressmanRepository.GetCongressmansDepartments(id)
	return congressmanDepartments
}

func (cs *CongressmanService) GetCongressmansFromJobs(jobs string) models.CongressmansModel {
	congressmans := models.CongressmansModel{
		{
			Id:              1,
			Uid:             "PA1008",
			Civility:        "M.",
			FirstName:       "Alain",
			LastName:        "David",
			Alpha:           "DavidA",
			Trigramme:       "ADA",
			BirthDate:       "1949-06-02 00:00:00",
			BirthCity:       "Libourne",
			BirthDepartment: "Gironde",
			BirthCountry:    "France",
			Jobtitle:        "Ingénieur",
			CatSocPro:       "Cadres d'entreprise",
			FamSocPro:       "Cadres et professions intellectuelles supérieures",
		},
		{
			Id:              2,
			Uid:             "PA1012",
			Civility:        "M.",
			FirstName:       "Charles",
			LastName:        "de la Verpillière",
			Alpha:           "Verpilliere (de la)",
			Trigramme:       "CDL",
			BirthDate:       "1954-05-31 00:00:00",
			BirthCity:       "Bourg-en-Bresse",
			BirthDepartment: "Ain",
			BirthCountry:    "France",
			Jobtitle:        "Conseiller d'État",
			CatSocPro:       "Cadres de la fonction publique, professions intellectuelles et  artistiques",
			FamSocPro:       "Cadres et professions intellectuelles supérieures",
		},
		{
			Id:              3,
			Uid:             "PA1029",
			Civility:        "M.",
			FirstName:       "Bernard",
			LastName:        "Deflesselles",
			Alpha:           "Deflesselles",
			Trigramme:       "BDF",
			BirthDate:       "1953-10-16 00:00:00",
			BirthCity:       "Paris 6ème",
			BirthDepartment: "Paris",
			BirthCountry:    "France",
			Jobtitle:        "Ingénieur",
			CatSocPro:       "Cadres d'entreprise",
			FamSocPro:       "Cadres et professions intellectuelles supérieures",
		},
	}
	return congressmans
}
