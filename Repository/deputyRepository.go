package Repository

import (
	"database/sql"
	"time"

	manager "github.com/plouiserre/exposecongressman/Manager"
	models "github.com/plouiserre/exposecongressman/Models"
)

type DeputyRepository struct {
	LogManager *manager.LogManager
}

func (dr DeputyRepository) InitDB() (db *sql.DB) {
	db, err := sql.Open("mysql", "ProcessDeputesData:ASimpleP@ssW0rd@/PROCESSDEPUTES")

	if err != nil {
		dr.LogManager.WriteErrorLog("Erreur connexion " + err.Error())
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return db
}

func (dr *DeputyRepository) AllDeputies() (*models.DeputiesModel, bool) {
	var deputies models.DeputiesModel
	db := dr.InitDB()
	noError := true

	rows, err := db.Query("select * FROM PROCESSDEPUTES.Deputy;")

	if err != nil {
		dr.LogManager.WriteErrorLog("Erreur requête " + err.Error())
		noError = false
	} else {
		defer rows.Close()

		for rows.Next() {
			var deputy models.DeputyModel
			err := rows.Scan(&deputy.Id, &deputy.StartDate, &deputy.EndDate, &deputy.RefDeputy, &deputy.MandateId)

			if err != nil {
				dr.LogManager.WriteErrorLog("Erreur récupération du résultat " + err.Error())
				noError = false
			}

			deputies = append(deputies, deputy)
		}
	}

	return &deputies, noError
}

func (dr *DeputyRepository) GetDeputy(id int) (*models.DeputyModel, bool) {
	var deputy models.DeputyModel
	db := dr.InitDB()
	noError := true

	row, err := db.Query("select * FROM PROCESSDEPUTES.Deputy where DeputyId=?;", id)

	if err != nil {
		dr.LogManager.WriteErrorLog("Erreur requête " + err.Error())
		noError = false
	} else {
		if row.Next() {
			errScan := row.Scan(&deputy.Id, &deputy.StartDate, &deputy.EndDate, &deputy.RefDeputy, &deputy.MandateId)

			if errScan != nil {
				dr.LogManager.WriteErrorLog("Erreur récupération du résultat " + errScan.Error())
				noError = false
			}
		}
		row.Close()
	}
	if deputy != (models.DeputyModel{}) {
		return &deputy, noError
	} else {
		return nil, noError
	}
}

func (dr *DeputyRepository) InsertDeputy(deputy *models.DeputyModel) (int64, bool) {
	db := dr.InitDB()
	var lid int64
	noError := true

	if deputy == nil {
		dr.LogManager.WriteErrorLog("No Data send to insert")
	} else {
		queryDeputy := "INSERT INTO PROCESSDEPUTES.Deputy(StartDate, EndDate, RefDeputy, MandateId) VALUES (?,?,?,?)"
		stmt, errPrepare := db.Prepare(queryDeputy)
		if errPrepare != nil {
			dr.LogManager.WriteErrorLog("Erreur récupération du résultat " + errPrepare.Error())
			noError = false
		} else {
			res, errExec := stmt.Exec(deputy.StartDate, deputy.EndDate, deputy.RefDeputy, deputy.MandateId)
			if errExec != nil {
				dr.LogManager.WriteErrorLog("Deputy Repository : Erreur exécution requête " + errExec.Error())
				noError = false
			} else {
				var errGetLastId error
				lid, errGetLastId = res.LastInsertId()
				if errGetLastId != nil {
					dr.LogManager.WriteErrorLog("Erreur lors de la récupération de l'id enregistré " + errGetLastId.Error())
				}
			}
		}
	}
	defer db.Close()

	return lid, noError
}

func (dr *DeputyRepository) UpdateDeputy(deputy *models.DeputyModel, id int) bool {
	db := dr.InitDB()
	noError := true

	queryDeputy := "UPDATE  PROCESSDEPUTES.Deputy SET StartDate=?, EndDate=?, RefDeputy=?, MandateId=? WHERE DeputyId = ?"
	stmt, errPrepare := db.Prepare(queryDeputy)
	if errPrepare != nil {
		dr.LogManager.WriteErrorLog("Erreur récupération du résultat " + errPrepare.Error())
		noError = false
	} else {
		_, errExec := stmt.Exec(deputy.StartDate, deputy.EndDate, deputy.RefDeputy, deputy.MandateId, id)
		if errExec != nil {
			dr.LogManager.WriteErrorLog("Congressman Repository : Erreur exécution requête " + errExec.Error())
			noError = false
		}
	}

	defer db.Close()
	return noError
}

func (dr *DeputyRepository) DeleteDeputy(id int) (int64, bool) {
	var nbDelete int64
	db := dr.InitDB()
	noError := true

	queryDeputy := "DELETE FROM PROCESSDEPUTES.Deputy WHERE DeputyId = ?"
	stmt, errPrepare := db.Prepare(queryDeputy)
	if errPrepare != nil {
		dr.LogManager.WriteErrorLog("Erreur récupération du résultat " + errPrepare.Error())
		noError = false
	} else {
		result, errExec := stmt.Exec(id)
		if errExec != nil {
			dr.LogManager.WriteErrorLog("Congressman Repository : Erreur exécution requête " + errExec.Error())
			noError = false
		}
		nbDelete, _ = result.RowsAffected()
	}
	defer db.Close()

	return nbDelete, noError
}
