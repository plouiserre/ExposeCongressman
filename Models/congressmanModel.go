package Models

import (
	"database/sql"

	manager "github.com/plouiserre/exposecongressman/Manager"
)

type CongressmanModel struct {
	Id              int64  `json:"Id"`
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

func (cm CongressmanModel) QueryGetById(db *sql.DB, id int) (*sql.Rows, error) {
	return nil, nil
}
func (cm CongressmanModel) RowsScanGetById(rows *sql.Rows, logManager *manager.LogManager) (EntityModel, bool) {
	return EntityModel{}, false
}

type CongressmansModel []CongressmanModel

func (cms CongressmansModel) GetQuery(db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query("select * FROM PROCESSDEPUTES.CongressMan;")
	return rows, err
}

func (cms CongressmansModel) RowsScanGetEntities(rows *sql.Rows, logManager *manager.LogManager) (EntityModel, bool) {
	var congressmans CongressmansModel
	var entities EntityModel
	noError := true

	for rows.Next() {
		var congressman CongressmanModel
		err := rows.Scan(&congressman.Id, &congressman.Uid, &congressman.Civility, &congressman.FirstName,
			&congressman.LastName, &congressman.Alpha, &congressman.Trigramme, &congressman.BirthDate,
			&congressman.BirthCity, &congressman.BirthDepartment, &congressman.BirthCountry,
			&congressman.Jobtitle, &congressman.CatSocPro, &congressman.FamSocPro)

		if err != nil {
			logManager.WriteErrorLog("Erreur récupération du résultat " + err.Error())
			noError = false
		}

		congressmans = append(congressmans, congressman)
	}
	entities.Congressmans = congressmans
	return entities, noError
}
