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

func (cm CongressmanModel) QueryGetById() string {
	query := "select * FROM PROCESSDEPUTES.CongressMan where CongressManId=?;"
	return query
}

func (cm CongressmanModel) RowsScanGetById(row *sql.Rows, logManager *manager.LogManager) (IModel, bool) {
	var congressman CongressmanModel
	noError := true

	errScan := row.Scan(&congressman.Id, &congressman.Uid, &congressman.Civility, &congressman.FirstName,
		&congressman.LastName, &congressman.Alpha, &congressman.Trigramme, &congressman.BirthDate,
		&congressman.BirthCity, &congressman.BirthDepartment, &congressman.BirthCountry,
		&congressman.Jobtitle, &congressman.CatSocPro, &congressman.FamSocPro)

	if errScan != nil {
		logManager.WriteErrorLog("Erreur récupération du résultat " + errScan.Error())
		noError = false
	}

	return congressman, noError
}

func (cm CongressmanModel) IsEntityFill(model IModel, logManager *manager.LogManager) bool {
	congressman := model.(CongressmanModel)
	if congressman != (CongressmanModel{}) {
		return true
	} else {
		logManager.WriteErrorLog("No Data send to insert")
		return false
	}
}
func (cm CongressmanModel) QueryCreate() string {
	queryCongressMan := "INSERT INTO PROCESSDEPUTES.Congressman(CongressManUid, Civility, FirstName, LastName, Alpha, Trigramme, BirthDate, BirthCity, BirthDepartment, BirthCountry, JobTitle, CatSocPro, FamSocPro) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)"

	return queryCongressMan
}

func (cm CongressmanModel) ExecuteCreateQuery(model IModel, stmt *sql.Stmt) (sql.Result, string, error) {
	congressman := model.(CongressmanModel)
	res, errExec := stmt.Exec(congressman.Uid, congressman.Civility, congressman.FirstName, congressman.LastName, congressman.Alpha,
		congressman.Trigramme, congressman.BirthDate, congressman.BirthCity, congressman.BirthDepartment, congressman.BirthCountry,
		congressman.Jobtitle, congressman.CatSocPro, congressman.FamSocPro)
	return res, "Congressman", errExec
}

func (cm CongressmanModel) QueryUpdate() string {
	queryCongressMan := "UPDATE  PROCESSDEPUTES.Congressman SET Civility=?, FirstName=?, LastName=?, Alpha=?, Trigramme=?, BirthDate=?, BirthCity=?, BirthDepartment=?, BirthCountry=?, JobTitle=?, CatSocPro=?, FamSocPro=? WHERE CongressManId = ?"
	return queryCongressMan
}

func (cm CongressmanModel) ExecuteUpdateQuery(stmt *sql.Stmt, id int64) (string, error) {
	_, errExec := stmt.Exec(cm.Civility, cm.FirstName, cm.LastName, cm.Alpha, cm.Trigramme, cm.BirthDate, cm.BirthCity, cm.BirthDepartment,
		cm.BirthCountry, cm.Jobtitle, cm.CatSocPro, cm.FamSocPro, id)
	return "Congressman ", errExec
}

func (cm CongressmanModel) QueryDelete() string {
	queryDeputy := "DELETE FROM PROCESSDEPUTES.Congressman WHERE CongressManId = ?"
	return queryDeputy
}

type CongressmansModel []CongressmanModel

func (cms CongressmansModel) GetQuery(db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query("select * FROM PROCESSDEPUTES.CongressMan;")
	return rows, err
}

func (cms CongressmansModel) RowsScanGetEntities(rows *sql.Rows, logManager *manager.LogManager) (IModels, bool) {
	var congressmans CongressmansModel
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
	return congressmans, noError
}
