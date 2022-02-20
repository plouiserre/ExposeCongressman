package Models

import (
	"database/sql"

	manager "github.com/plouiserre/exposecongressman/Manager"
)

//séparer en plusieurs fichiers

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

func (mm MandateModel) QueryGetById(db *sql.DB, id int) (*sql.Rows, error) {
	row, err := db.Query("select * FROM PROCESSDEPUTES.Mandate where MandateId=?;", id)
	return row, err
}

func (mm MandateModel) RowsScanGetById(row *sql.Rows, logManager *manager.LogManager) (EntityModel, bool) {
	var mandate MandateModel
	var entity EntityModel
	noError := true

	errScan := row.Scan(&mandate.Id, &mandate.Uid, &mandate.TermOffice, &mandate.TypeOrgane,
		&mandate.StartDate, &mandate.EndDate, &mandate.Precedence, &mandate.PrincipleNoming,
		&mandate.QualityCode, &mandate.QualityLabel, &mandate.QualityLabelSex,
		&mandate.RefBody, &mandate.CongressmanId)

	if errScan != nil {
		logManager.WriteErrorLog("Erreur récupération du résultat " + errScan.Error())
		noError = false

	}
	entity.Mandate = mandate
	return entity, noError
}

func (mm MandateModel) IsEntityFill(model EntityModel, logManager *manager.LogManager) bool {
	if model.Mandate != (MandateModel{}) {
		return true
	} else {
		logManager.WriteErrorLog("No Data send to insert")
		return false
	}
}

func (mm MandateModel) PrepareCreateQuery(db *sql.DB, logManager *manager.LogManager) (*sql.Stmt, bool) {
	noError := true
	queryMandate := "INSERT INTO PROCESSDEPUTES.Mandate(MandateUid, TermOffice, TypeOrgane, StartDate, EndDate, Precedence, PrincipleNoming, QualityCode, QualityLabel, QualityLabelSex, RefBody, CongressManId) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)"
	stmt, errPrepare := db.Prepare(queryMandate)
	if errPrepare != nil {
		logManager.WriteErrorLog("Erreur récupération du résultat " + errPrepare.Error())
		noError = false
	}
	return stmt, noError
}

func (mm MandateModel) ExecuteCreateQuery(stmt *sql.Stmt, model EntityModel) (sql.Result, string, error) {
	mandate := model.Mandate
	res, errExec := stmt.Exec(mandate.Uid, mandate.TermOffice, mandate.TypeOrgane, mandate.StartDate, mandate.EndDate,
		mandate.Precedence, mandate.PrincipleNoming, mandate.QualityCode, mandate.QualityLabel, mandate.QualityLabelSex,
		mandate.RefBody, mandate.CongressmanId)
	return res, "Mandate ", errExec
}

func (mm MandateModel) PrepareUpdateQuery(db *sql.DB, logManager *manager.LogManager) (*sql.Stmt, bool) {
	noError := true
	queryMandate := "UPDATE  PROCESSDEPUTES.Mandate SET MandateUid=?, TermOffice=?, TypeOrgane=?,StartDate=?, EndDate=?, Precedence=?, PrincipleNoming=?, QualityCode=?, QualityLabel=?, QualityLabelSex=?, RefBody=?, CongressManId=? WHERE MandateId = ?"
	stmt, errPrepare := db.Prepare(queryMandate)
	if errPrepare != nil {
		logManager.WriteErrorLog("Erreur récupération du résultat " + errPrepare.Error())
		noError = false
	}
	return stmt, noError
}

func (mm MandateModel) ExecuteUpdateQuery(stmt *sql.Stmt, model EntityModel, id int) (string, error) {
	mandate := model.Mandate
	_, errExec := stmt.Exec(mandate.Uid, mandate.TermOffice, mandate.TypeOrgane, mandate.StartDate,
		mandate.EndDate, mandate.Precedence, mandate.PrincipleNoming, mandate.QualityCode, mandate.QualityLabel,
		mandate.QualityLabelSex, mandate.RefBody, mandate.CongressmanId, id)
	return "Mandate ", errExec
}

type MandatesModel []MandateModel

func (mms MandatesModel) GetQuery(db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query("select * FROM PROCESSDEPUTES.Mandate;")
	return rows, err
}

//TODO à la fin on ne renverra plus de EntityModel mais un IModels
func (mms MandatesModel) RowsScanGetEntities(rows *sql.Rows, logManager *manager.LogManager) (EntityModel, bool) {
	var mandates MandatesModel
	var entities EntityModel
	noError := true

	for rows.Next() {
		var mandate MandateModel
		err := rows.Scan(&mandate.Id, &mandate.Uid, &mandate.TermOffice, &mandate.TypeOrgane,
			&mandate.StartDate, &mandate.EndDate, &mandate.Precedence, &mandate.PrincipleNoming,
			&mandate.QualityCode, &mandate.QualityLabel, &mandate.QualityLabelSex,
			&mandate.RefBody, &mandate.CongressmanId)

		if err != nil {
			logManager.WriteErrorLog("Erreur récupération du résultat " + err.Error())
			noError = false
		}

		mandates = append(mandates, mandate)
	}
	entities.Mandates = mandates
	return entities, noError
}
