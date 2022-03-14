package Service

import (
	models "github.com/plouiserre/exposecongressman/Models"
	repository "github.com/plouiserre/exposecongressman/Repository"
)

type CongressmanService struct {
	CongressmanRepository repository.CongressmanRepository
}

func (cs *CongressmanService) GetMandatesFromCongressman(id int64) models.CongressmanMandatesModel {
	/*mandates := []models.MandateModel{

		models.MandateModel{
			Id:              113,
			Uid:             "PM768381",
			TermOffice:      15,
			TypeOrgane:      "MISINFO",
			StartDate:       "2019-11-20 00:00:00",
			Precedence:      "5",
			PrincipleNoming: "1",
			QualityCode:     "Co-rapporteur",
			QualityLabel:    "Co-rapporteur",
			QualityLabelSex: "Co-rapporteur",
			CongressmanId:   "5",
		},
		models.MandateModel{
			Id:              114,
			Uid:             "PM768379",
			TermOffice:      15,
			TypeOrgane:      "MISINFO",
			StartDate:       "2019-11-20 00:00:00",
			Precedence:      "10",
			PrincipleNoming: "0",
			QualityCode:     "Membre",
			QualityLabel:    "Membre",
			QualityLabelSex: "Membre",
			CongressmanId:   "5",
		},
	}
	congressmanMandates := models.CongressmanMandatesModel{
		CongressmanId: 5,
		Mandates:      mandates,
	}
	return congressmanMandates*/
	mandates := cs.CongressmanRepository.GetCongressmansMandates(id)
	congressmanMandates := models.CongressmanMandatesModel{
		CongressmanId: id,
		Mandates:      *mandates,
	}
	return congressmanMandates
}
