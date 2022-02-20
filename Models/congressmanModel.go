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
	row, err := db.Query("select * FROM PROCESSDEPUTES.CongressMan where CongressManId=?;", id)
	return row, err
}

func (cm CongressmanModel) RowsScanGetById(row *sql.Rows, logManager *manager.LogManager) (EntityModel, bool) {
	var congressman CongressmanModel
	var entity EntityModel
	noError := true

	errScan := row.Scan(&congressman.Id, &congressman.Uid, &congressman.Civility, &congressman.FirstName,
		&congressman.LastName, &congressman.Alpha, &congressman.Trigramme, &congressman.BirthDate,
		&congressman.BirthCity, &congressman.BirthDepartment, &congressman.BirthCountry,
		&congressman.Jobtitle, &congressman.CatSocPro, &congressman.FamSocPro)

	if errScan != nil {
		logManager.WriteErrorLog("Erreur récupération du résultat " + errScan.Error())
		noError = false
	}
	entity.Congressman = congressman
	return entity, noError
}

func (cm CongressmanModel) IsEntityFill(entity EntityModel, logManager *manager.LogManager) bool {
	if entity.Congressman != (CongressmanModel{}) {
		return true
	} else {
		logManager.WriteErrorLog("No Data send to insert")
		return false
	}
}
func (cm CongressmanModel) PrepareCreateQuery(db *sql.DB, logManager *manager.LogManager) (*sql.Stmt, bool) {
	noError := true
	queryCongressMan := "INSERT INTO PROCESSDEPUTES.Congressman(CongressManUid, Civility, FirstName, LastName, Alpha, Trigramme, BirthDate, BirthCity, BirthDepartment, BirthCountry, JobTitle, CatSocPro, FamSocPro) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)"
	stmt, errPrepare := db.Prepare(queryCongressMan)
	if errPrepare != nil {
		logManager.WriteErrorLog("Erreur récupération du résultat " + errPrepare.Error())
		noError = false
	}
	return stmt, noError
}

func (cm CongressmanModel) ExecuteCreateQuery(stmt *sql.Stmt, model EntityModel) (sql.Result, string, error) {
	congressman := model.Congressman
	res, errExec := stmt.Exec(congressman.Uid, congressman.Civility, congressman.FirstName,
		congressman.LastName, congressman.Alpha, congressman.Trigramme, congressman.BirthDate,
		congressman.BirthCity, congressman.BirthDepartment, congressman.BirthCountry,
		congressman.Jobtitle, congressman.CatSocPro, congressman.FamSocPro)
	return res, "Congressman", errExec
}

func (cm CongressmanModel) PrepareUpdateQuery(db *sql.DB, logManager *manager.LogManager) (*sql.Stmt, bool) {
	noError := true
	queryCongressMan := "UPDATE  PROCESSDEPUTES.Congressman SET Civility=?, FirstName=?, LastName=?, Alpha=?, Trigramme=?, BirthDate=?, BirthCity=?, BirthDepartment=?, BirthCountry=?, JobTitle=?, CatSocPro=?, FamSocPro=? WHERE CongressManId = ?"
	stmt, errPrepare := db.Prepare(queryCongressMan)
	if errPrepare != nil {
		logManager.WriteErrorLog("Erreur récupération du résultat " + errPrepare.Error())
		noError = false
	}
	return stmt, noError
}

func (cm CongressmanModel) ExecuteUpdateQuery(stmt *sql.Stmt, model EntityModel, id int) (string, error) {
	congressman := model.Congressman
	_, errExec := stmt.Exec(congressman.Civility, congressman.FirstName, congressman.LastName,
		congressman.Alpha, congressman.Trigramme, congressman.BirthDate,
		congressman.BirthCity, congressman.BirthDepartment, congressman.BirthCountry,
		congressman.Jobtitle, congressman.CatSocPro, congressman.FamSocPro, id)
	return "Congressman ", errExec
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
