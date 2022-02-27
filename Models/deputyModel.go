package Models

import (
	"database/sql"

	manager "github.com/plouiserre/exposecongressman/Manager"
)

type DeputyModel struct {
	Id        int64          `json:"Id"`
	StartDate string         `json:"StartDate"`
	EndDate   sql.NullString `json:"EndDate"`
	RefDeputy string         `json:"RefDeputy"`
	MandateId int64          `json:"MandateId"`
}

func (dm DeputyModel) QueryGetById() string {
	query := "select * FROM PROCESSDEPUTES.Deputy where DeputyId=?;"
	return query
}

func (dm DeputyModel) RowsScanGetById(row *sql.Rows, logManager *manager.LogManager) (IModel, bool) {
	var deputy DeputyModel
	noError := true
	errScan := row.Scan(&deputy.Id, &deputy.StartDate, &deputy.EndDate, &deputy.RefDeputy, &deputy.MandateId)

	if errScan != nil {
		logManager.WriteErrorLog("Erreur récupération du résultat " + errScan.Error())
		noError = false
	}
	return deputy, noError
}

func (dm DeputyModel) IsEntityFill(model IModel, logManager *manager.LogManager) bool {
	deputy := model.(DeputyModel)
	if deputy != (DeputyModel{}) {
		return true
	} else {
		logManager.WriteErrorLog("No Data send to insert")
		return false
	}
}

func (dm DeputyModel) QueryCreate() string {
	queryDeputy := "INSERT INTO PROCESSDEPUTES.Deputy(StartDate, EndDate, RefDeputy, MandateId) VALUES (?,?,?,?)"

	return queryDeputy
}

func (dm DeputyModel) ExecuteCreateQuery(model IModel, stmt *sql.Stmt) (sql.Result, string, error) {
	res, errExec := stmt.Exec(dm.StartDate, dm.EndDate, dm.RefDeputy, dm.MandateId)
	return res, "Deputy", errExec
}

func (dm DeputyModel) QueryUpdate() string {
	queryDeputy := "UPDATE  PROCESSDEPUTES.Deputy SET StartDate=?, EndDate=?, RefDeputy=?, MandateId=? WHERE DeputyId = ?"
	return queryDeputy
}

func (dm DeputyModel) ExecuteUpdateQuery(stmt *sql.Stmt, id int64) (string, error) {
	_, errExec := stmt.Exec(dm.StartDate, dm.EndDate, dm.RefDeputy, dm.MandateId, id)
	return "Deputy ", errExec
}

//TODO mieux factoriser en mettant juste la requête
func (dm DeputyModel) QueryDelete() string {
	queryDeputy := "DELETE FROM PROCESSDEPUTES.Deputy WHERE DeputyId = ?"
	return queryDeputy
}

type DeputiesModel []DeputyModel

func (dms DeputiesModel) GetQuery(db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query("select * FROM PROCESSDEPUTES.Deputy;")
	return rows, err
}

func (dms DeputiesModel) RowsScanGetEntities(rows *sql.Rows, logManager *manager.LogManager) (IModels, bool) {
	var deputies DeputiesModel
	noError := true

	for rows.Next() {
		var deputy DeputyModel
		err := rows.Scan(&deputy.Id, &deputy.StartDate, &deputy.EndDate, &deputy.RefDeputy, &deputy.MandateId)

		if err != nil {
			logManager.WriteErrorLog("Erreur récupération du résultat " + err.Error())
			noError = false
		}

		deputies = append(deputies, deputy)
	}
	return deputies, noError
}
