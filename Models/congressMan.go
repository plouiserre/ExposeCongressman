package models

type CongressMan struct {
	Id              int    `json:"Id"`
	Uid             string `json:"Uid"`
	Civility        string `json:"Civility"`
	FirstName       string `json:"FirstName"`
	LastName        string `json:"LastName"`
	Alpha           string `json:"Alpha"`
	Trigramme       string `json:"Trigramme"`
	BirthDate       string `json:"BirthDate"`
	BirthCity       string `json:"BirthCity"`
	BirthDepartment string `json:"BirthDepartment"`
	BirthCountry    string `json:"BirthCountry"`
	Jobtitle        string `json:"Jobtitle"`
	CatSocPro       string `json:"CatSocPro"`
	FamSocPro       string `json:"FamSocPro"`
}

type CongressMans []CongressMan

//TODO externaliser ensuite dans une classe repository quand on aura arrêter de mocker
func AllCongressMans() *CongressMans {
	var congressMans CongressMans
	firstCongressMan := CongressMan{
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
		FamSocPro:       ",'Cadres et professions intellectuelles supérieures'",
	}
	secondCongressMan := CongressMan{
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
		FamSocPro:       ",'Cadres et professions intellectuelles supérieures'",
	}
	congressMans = append(congressMans, firstCongressMan)
	congressMans = append(congressMans, secondCongressMan)
	return &congressMans
}
