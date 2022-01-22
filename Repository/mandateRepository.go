package Repository

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	manager "github.com/plouiserre/exposecongressman/Manager"
	models "github.com/plouiserre/exposecongressman/Models"
)

type MandateRepository struct {
	LogManager *manager.LogManager
}

//TODO factoriser avec les autres repository
func (mr *MandateRepository) InitDB() (db *sql.DB) {
	db, err := sql.Open("mysql", "ProcessDeputesData:ASimpleP@ssW0rd@/PROCESSDEPUTES")

	if err != nil {
		mr.LogManager.WriteErrorLog("Erreur connexion " + err.Error())
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return db
}

func (mr *MandateRepository) AllMandates() (*models.MandatesModel, bool) {
	var mandates models.MandatesModel
	db := mr.InitDB()
	noError := true

	rows, err := db.Query("select * FROM PROCESSDEPUTES.Mandate;")

	if err != nil {
		mr.LogManager.WriteErrorLog("Erreur requête " + err.Error())
		noError = false
	} else {
		defer rows.Close()

		for rows.Next() {
			var mandate models.MandateModel
			err := rows.Scan(&mandate.Id, &mandate.Uid, &mandate.TermOffice, &mandate.TypeOrgane,
				&mandate.StartDate, &mandate.EndDate, &mandate.Precedence, &mandate.PrincipleNaming,
				&mandate.QualityCode, &mandate.QualityLabel, &mandate.QualityLabelSex,
				&mandate.RefBody, &mandate.CongressmanId)

			if err != nil {
				mr.LogManager.WriteErrorLog("Erreur récupération du résultat " + err.Error())
				noError = false
			}

			mandates = append(mandates, mandate)
		}
	}

	return &mandates, noError
}

func (mr *MandateRepository) GetMandate(id int) (*models.MandateModel, bool) {
	var mandate models.MandateModel
	db := mr.InitDB()
	noError := true

	row, err := db.Query("select * FROM PROCESSDEPUTES.Mandate where MandateId=?;", id)

	if err != nil {
		mr.LogManager.WriteErrorLog("Erreur requête " + err.Error())
		noError = false
	} else {
		if row.Next() {
			errScan := row.Scan(&mandate.Id, &mandate.Uid, &mandate.TermOffice, &mandate.TypeOrgane,
				&mandate.StartDate, &mandate.EndDate, &mandate.Precedence, &mandate.PrincipleNaming,
				&mandate.QualityCode, &mandate.QualityLabel, &mandate.QualityLabelSex,
				&mandate.RefBody, &mandate.CongressmanId)

			if errScan != nil {
				mr.LogManager.WriteErrorLog("Erreur récupération du résultat " + errScan.Error())
				noError = false
			}
		}
		row.Close()
	}
	if mandate != (models.MandateModel{}) {
		return &mandate, noError
	} else {
		return nil, noError
	}
}

func (mr *MandateRepository) InsertMandate(mandate *models.MandateModel) (int64, bool) {
	db := mr.InitDB()
	var lid int64
	noError := true

	if mandate == nil {
		mr.LogManager.WriteErrorLog("No Data send to insert")
	} else {
		queryMandate := "INSERT INTO PROCESSDEPUTES.Mandate(MandateUid, TermOffice, StartDate, EndDate, Precedence, PrincipleNoming, QualityCode, QualityLabel, QualityLabelSex, RefBody, CongressManId) VALUES (?,?,?,?,?,?,?,?,?,?,?)"
		stmt, errPrepare := db.Prepare(queryMandate)
		if errPrepare != nil {
			mr.LogManager.WriteErrorLog("Erreur récupération du résultat " + errPrepare.Error())
			noError = false
		} else {
			res, errExec := stmt.Exec(mandate.Id, mandate.Uid, mandate.TermOffice, mandate.TypeOrgane, mandate.StartDate, mandate.EndDate,
				mandate.Precedence, mandate.PrincipleNaming, mandate.QualityCode, mandate.QualityLabel, mandate.QualityLabelSex,
				mandate.RefBody, mandate.CongressmanId)
			if errExec != nil {
				mr.LogManager.WriteErrorLog("Mandate Repository : Erreur exécution requête " + errExec.Error())
				noError = false
			} else {
				var errGetLastId error
				lid, errGetLastId = res.LastInsertId()
				if errGetLastId != nil {
					mr.LogManager.WriteErrorLog("Erreur lors de la récupération de l'id enregistré " + errGetLastId.Error())
				}
			}
		}
	}
	defer db.Close()

	return lid, noError
}

func (mr *MandateRepository) UpdateCongressMan(mandate *models.MandateModel, id int) bool {
	db := mr.InitDB()
	noError := true

	queryMandate := "UPDATE  PROCESSDEPUTES.Mandate SET MandateUid=?, TermOffice=?, StartDate=?, EndDate=?, Precedence=?, PrincipleNoming=?, QualityCode=?, QualityLabel=?, QualityLabelSex=?, RefBody=?, CongressManId=? WHERE MandateId = ?"
	stmt, errPrepare := db.Prepare(queryMandate)
	if errPrepare != nil {
		mr.LogManager.WriteErrorLog("Erreur récupération du résultat " + errPrepare.Error())
		noError = false
	} else {
		_, errExec := stmt.Exec(mandate.Uid, mandate.TermOffice, mandate.StartDate, mandate.EndDate, mandate.Precedence,
			mandate.PrincipleNaming, mandate.QualityCode, mandate.QualityLabel, mandate.QualityLabelSex,
			mandate.RefBody, id)
		if errExec != nil {
			mr.LogManager.WriteErrorLog("Congressman Repository : Erreur exécution requête " + errExec.Error())
			noError = false
		}
	}

	defer db.Close()
	return noError
}

func (mr *MandateRepository) DeleteCongressMan(id int) (int64, bool) {
	var nbDelete int64
	db := mr.InitDB()
	noError := true

	queryCongressMan := "DELETE FROM PROCESSDEPUTES.Mandate WHERE MandateId = ?"
	stmt, errPrepare := db.Prepare(queryCongressMan)
	if errPrepare != nil {
		mr.LogManager.WriteErrorLog("Erreur récupération du résultat " + errPrepare.Error())
		noError = false
	} else {
		result, errExec := stmt.Exec(id)
		if errExec != nil {
			mr.LogManager.WriteErrorLog("Congressman Repository : Erreur exécution requête " + errExec.Error())
			noError = false
		}
		nbDelete, _ = result.RowsAffected()
	}
	defer db.Close()

	return nbDelete, noError
}
