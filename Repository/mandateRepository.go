package Repository

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/plouiserre/exposecongressman/Manager"
	manager "github.com/plouiserre/exposecongressman/Manager"
	models "github.com/plouiserre/exposecongressman/Models"
)

type MandateRepository struct {
	LogManager *manager.LogManager
}

//TODO factoriser avec les autres repository
func (mr MandateRepository) InitDB() (db *sql.DB) {
	db, err := sql.Open("mysql", "ProcessDeputesData:ASimpleP@ssW0rd@/PROCESSDEPUTES")

	if err != nil {
		mr.LogManager.WriteErrorLog("Erreur connexion " + err.Error())
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return db
}

/*func (mr *MandateRepository) AllMandates() (*models.MandatesModel, bool) {
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
				&mandate.StartDate, &mandate.EndDate, &mandate.Precedence, &mandate.PrincipleNoming,
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
}*/

//TODO quand ca marchera redéplacer tout le code de AllMandates dans GetAll()
func (mr MandateRepository) GetAll() (*models.EntityModel, bool) {
	var mandates models.MandatesModel
	var entities models.EntityModel
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
				&mandate.StartDate, &mandate.EndDate, &mandate.Precedence, &mandate.PrincipleNoming,
				&mandate.QualityCode, &mandate.QualityLabel, &mandate.QualityLabelSex,
				&mandate.RefBody, &mandate.CongressmanId)

			if err != nil {
				mr.LogManager.WriteErrorLog("Erreur récupération du résultat " + err.Error())
				noError = false
			}

			mandates = append(mandates, mandate)
		}
		entities.Mandates = mandates
	}

	return &entities, noError
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
				&mandate.StartDate, &mandate.EndDate, &mandate.Precedence, &mandate.PrincipleNoming,
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

func (mr MandateRepository) GetById(id int) (*models.EntityModel, bool) {
	var entity models.EntityModel

	mandate, noError := mr.GetMandate(id)
	if mandate != nil {
		entity.Mandate = *mandate
		return &entity, noError
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
		queryMandate := "INSERT INTO PROCESSDEPUTES.Mandate(MandateUid, TermOffice, TypeOrgane, StartDate, EndDate, Precedence, PrincipleNoming, QualityCode, QualityLabel, QualityLabelSex, RefBody, CongressManId) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)"
		stmt, errPrepare := db.Prepare(queryMandate)
		if errPrepare != nil {
			mr.LogManager.WriteErrorLog("Erreur récupération du résultat " + errPrepare.Error())
			noError = false
		} else {
			res, errExec := stmt.Exec(mandate.Uid, mandate.TermOffice, mandate.TypeOrgane, mandate.StartDate, mandate.EndDate,
				mandate.Precedence, mandate.PrincipleNoming, mandate.QualityCode, mandate.QualityLabel, mandate.QualityLabelSex,
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

func (mr MandateRepository) CreateEntity(entity *models.EntityModel) (int64, bool) {
	lid, noError := mr.InsertMandate(&entity.Mandate)

	return lid, noError
}

func (mr *MandateRepository) UpdateMandate(mandate *models.MandateModel, id int) bool {
	db := mr.InitDB()
	noError := true

	queryMandate := "UPDATE  PROCESSDEPUTES.Mandate SET MandateUid=?, TermOffice=?, TypeOrgane=?,StartDate=?, EndDate=?, Precedence=?, PrincipleNoming=?, QualityCode=?, QualityLabel=?, QualityLabelSex=?, RefBody=?, CongressManId=? WHERE MandateId = ?"
	stmt, errPrepare := db.Prepare(queryMandate)
	if errPrepare != nil {
		mr.LogManager.WriteErrorLog("Erreur récupération du résultat " + errPrepare.Error())
		noError = false
	} else {
		_, errExec := stmt.Exec(mandate.Uid, mandate.TermOffice, mandate.TypeOrgane, mandate.StartDate,
			mandate.EndDate, mandate.Precedence, mandate.PrincipleNoming, mandate.QualityCode, mandate.QualityLabel,
			mandate.QualityLabelSex, mandate.RefBody, mandate.CongressmanId, id)
		if errExec != nil {
			mr.LogManager.WriteErrorLog("Mandate Repository : Erreur exécution requête " + errExec.Error())
			noError = false
		}
	}

	defer db.Close()
	return noError
}

func (mr MandateRepository) UpdateEntity(entity *models.EntityModel, id int) bool {
	noError := mr.UpdateMandate(&entity.Mandate, id)
	return noError
}

func (mr *MandateRepository) DeleteMandate(id int) (int64, bool) {
	var nbDelete int64
	db := mr.InitDB()
	noError := true

	queryMandate := "DELETE FROM PROCESSDEPUTES.Mandate WHERE MandateId = ?"
	stmt, errPrepare := db.Prepare(queryMandate)
	if errPrepare != nil {
		mr.LogManager.WriteErrorLog("Erreur récupération du résultat " + errPrepare.Error())
		noError = false
	} else {
		result, errExec := stmt.Exec(id)
		if errExec != nil {
			mr.LogManager.WriteErrorLog("Mandate Repository : Erreur exécution requête " + errExec.Error())
			noError = false
		}
		nbDelete, _ = result.RowsAffected()
	}
	defer db.Close()

	return nbDelete, noError
}

func (mr MandateRepository) DeleteEntity(id int) (int64, bool) {
	nbDelete, noError := mr.DeleteMandate(id)
	return nbDelete, noError
}

func (mr MandateRepository) InitRepository() (IRepository, Manager.LogManager) {
	return nil, manager.LogManager{}
}
