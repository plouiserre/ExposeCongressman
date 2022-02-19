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
