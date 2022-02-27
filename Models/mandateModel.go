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

func (mm MandateModel) QueryGetById() string {
	query := "select * FROM PROCESSDEPUTES.Mandate where MandateId=?;"
	return query
}

func (mm MandateModel) RowsScanGetById(row *sql.Rows, logManager *manager.LogManager) (IModel, bool) {
	var mandate MandateModel
	noError := true

	errScan := row.Scan(&mandate.Id, &mandate.Uid, &mandate.TermOffice, &mandate.TypeOrgane,
		&mandate.StartDate, &mandate.EndDate, &mandate.Precedence, &mandate.PrincipleNoming,
		&mandate.QualityCode, &mandate.QualityLabel, &mandate.QualityLabelSex,
		&mandate.RefBody, &mandate.CongressmanId)

	if errScan != nil {
		logManager.WriteErrorLog("Erreur récupération du résultat " + errScan.Error())
		noError = false

	}

	return mandate, noError
}

//TODO to delete ??
func (mm MandateModel) IsEntityFill(model IModel, logManager *manager.LogManager) bool {
	mandate := model.(MandateModel)
	if mandate != (MandateModel{}) {
		return true
	} else {
		logManager.WriteErrorLog("No Data send to insert")
		return false
	}
}

func (mm MandateModel) QueryCreate() string {
	queryMandate := "INSERT INTO PROCESSDEPUTES.Mandate(MandateUid, TermOffice, TypeOrgane, StartDate, EndDate, Precedence, PrincipleNoming, QualityCode, QualityLabel, QualityLabelSex, RefBody, CongressManId) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)"

	return queryMandate
}

func (mm MandateModel) ExecuteCreateQuery(model IModel, stmt *sql.Stmt) (sql.Result, string, error) {
	res, errExec := stmt.Exec(mm.Uid, mm.TermOffice, mm.TypeOrgane, mm.StartDate, mm.EndDate, mm.Precedence, mm.PrincipleNoming,
		mm.QualityCode, mm.QualityLabel, mm.QualityLabelSex, mm.RefBody, mm.CongressmanId)
	return res, "Mandate ", errExec
}

func (mm MandateModel) QueryUpdate() string {
	queryMandate := "UPDATE  PROCESSDEPUTES.Mandate SET MandateUid=?, TermOffice=?, TypeOrgane=?,StartDate=?, EndDate=?, Precedence=?, PrincipleNoming=?, QualityCode=?, QualityLabel=?, QualityLabelSex=?, RefBody=?, CongressManId=? WHERE MandateId = ?"
	return queryMandate
}

func (mm MandateModel) ExecuteUpdateQuery(stmt *sql.Stmt, id int64) (string, error) {
	_, errExec := stmt.Exec(mm.Uid, mm.TermOffice, mm.TypeOrgane, mm.StartDate, mm.EndDate, mm.Precedence, mm.PrincipleNoming, mm.QualityCode,
		mm.QualityLabel, mm.QualityLabelSex, mm.RefBody, mm.CongressmanId, id)
	return "Mandate ", errExec
}

func (mm MandateModel) QueryDelete() string {
	queryDeputy := "DELETE FROM PROCESSDEPUTES.Mandate WHERE MandateId = ?"
	return queryDeputy
}

type MandatesModel []MandateModel

func (mms MandatesModel) GetQuery(db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query("select * FROM PROCESSDEPUTES.Mandate;")
	return rows, err
}

//TODO à la fin on ne renverra plus de EntityModel mais un IModels
func (mms MandatesModel) RowsScanGetEntities(rows *sql.Rows, logManager *manager.LogManager) (IModels, bool) {
	var mandates MandatesModel
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

	return mandates, noError
}
