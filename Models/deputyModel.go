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

func (dm DeputyModel) QueryGetById(db *sql.DB, id int) (*sql.Rows, error) {
	row, err := db.Query("select * FROM PROCESSDEPUTES.Deputy where DeputyId=?;", id)
	return row, err
}
func (dm DeputyModel) RowsScanGetById(row *sql.Rows, logManager *manager.LogManager) (EntityModel, bool) {
	var deputy DeputyModel
	var entity EntityModel
	noError := true
	errScan := row.Scan(&deputy.Id, &deputy.StartDate, &deputy.EndDate, &deputy.RefDeputy, &deputy.MandateId)

	if errScan != nil {
		logManager.WriteErrorLog("Erreur récupération du résultat " + errScan.Error())
		noError = false
	}
	entity.Deputy = deputy
	return entity, noError
}

func (dm DeputyModel) IsEntityFill(entity EntityModel, logManager *manager.LogManager) bool {
	if entity.Deputy != (DeputyModel{}) {
		return true
	} else {
		logManager.WriteErrorLog("No Data send to insert")
		return false
	}
}

func (dm DeputyModel) PrepareCreateQuery(db *sql.DB, logManager *manager.LogManager) (*sql.Stmt, bool) {
	noError := true
	queryDeputy := "INSERT INTO PROCESSDEPUTES.Deputy(StartDate, EndDate, RefDeputy, MandateId) VALUES (?,?,?,?)"
	stmt, errPrepare := db.Prepare(queryDeputy)
	if errPrepare != nil {
		logManager.WriteErrorLog("Erreur récupération du résultat " + errPrepare.Error())
		noError = false
	}
	return stmt, noError
}

func (dm DeputyModel) ExecuteCreateQuery(stmt *sql.Stmt, model EntityModel) (sql.Result, string, error) {
	deputy := model.Deputy
	res, errExec := stmt.Exec(deputy.StartDate, deputy.EndDate, deputy.RefDeputy, deputy.MandateId)
	return res, "Deputy", errExec
}

type DeputiesModel []DeputyModel

func (dms DeputiesModel) GetQuery(db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query("select * FROM PROCESSDEPUTES.Deputy;")
	return rows, err
}

func (dms DeputiesModel) RowsScanGetEntities(rows *sql.Rows, logManager *manager.LogManager) (EntityModel, bool) {
	var deputies DeputiesModel
	var entities EntityModel
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
	entities.Deputies = deputies
	return entities, noError
}
