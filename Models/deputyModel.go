package Models

type DeputyModel struct {
	Id        int64  `json:"Id"`
	StartDate string `json:"StartDate"`
	EndDate   string `json:"EndDate"`
	RefDeputy string `json:"RefDeputy"`
	MandateId int64  `json:"MandateId"`
}

type DeputiesModel []DeputyModel
