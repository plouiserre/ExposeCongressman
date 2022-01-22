package Models

import "database/sql"

type MandateModel struct {
	Id              int64          `json:"Id"`
	Uid             string         `json:"Uid"`
	TermOffice      int64          `json:"TermOffice"`
	TypeOrgane      string         `json:"TypeOrgane"`
	StartDate       string         `json:"StartDate"`
	EndDate         sql.NullString `json:"EndDate"`
	Precedence      string         `json:"Precedence"`
	PrincipleNoming string         `json:"PrincipleNoming"`
	QualityCode     string         `json:"QualityCode"`
	QualityLabel    string         `json:"QualityLabel"`
	QualityLabelSex string         `json:"QualityLabelSex"`
	RefBody         sql.NullString `json:"RefBody"`
	CongressmanId   string         `json:"CongressmanId"`
}

type MandatesModel []MandateModel
