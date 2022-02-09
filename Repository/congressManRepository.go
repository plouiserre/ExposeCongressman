package Repository

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/plouiserre/exposecongressman/Manager"
	manager "github.com/plouiserre/exposecongressman/Manager"
	models "github.com/plouiserre/exposecongressman/Models"
)

type CongressmanRepository struct {
	LogManager *manager.LogManager
}

func (cr CongressmanRepository) InitDB() (db *sql.DB) {
	db, err := sql.Open("mysql", "ProcessDeputesData:ASimpleP@ssW0rd@/PROCESSDEPUTES")

	if err != nil {
		cr.LogManager.WriteErrorLog("Erreur connexion " + err.Error())
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return db
}

/*func (cr *CongressmanRepository) AllCongressMans() (*models.CongressmansModel, bool) {
	var congressMans models.CongressmansModel
	db := cr.InitDB()
	noError := true

	rows, err := db.Query("select * FROM PROCESSDEPUTES.CongressMan;")

	if err != nil {
		cr.LogManager.WriteErrorLog("Erreur requête " + err.Error())
		noError = false
	} else {
		defer rows.Close()

		for rows.Next() {
			var congressman models.CongressmanModel
			err := rows.Scan(&congressman.Id, &congressman.Uid, &congressman.Civility, &congressman.FirstName,
				&congressman.LastName, &congressman.Alpha, &congressman.Trigramme, &congressman.BirthDate,
				&congressman.BirthCity, &congressman.BirthDepartment, &congressman.BirthCountry,
				&congressman.Jobtitle, &congressman.CatSocPro, &congressman.FamSocPro)

			if err != nil {
				cr.LogManager.WriteErrorLog("Erreur récupération du résultat " + err.Error())
				noError = false
			}

			congressMans = append(congressMans, congressman)
		}
	}

	return &congressMans, noError
}*/

func (cr CongressmanRepository) GetAll() (*models.EntityModel, bool) {
	var congressmans models.CongressmansModel
	var entities models.EntityModel
	db := cr.InitDB()
	noError := true

	rows, err := db.Query("select * FROM PROCESSDEPUTES.CongressMan;")

	if err != nil {
		cr.LogManager.WriteErrorLog("Erreur requête " + err.Error())
		noError = false
	} else {
		defer rows.Close()

		for rows.Next() {
			var congressman models.CongressmanModel
			err := rows.Scan(&congressman.Id, &congressman.Uid, &congressman.Civility, &congressman.FirstName,
				&congressman.LastName, &congressman.Alpha, &congressman.Trigramme, &congressman.BirthDate,
				&congressman.BirthCity, &congressman.BirthDepartment, &congressman.BirthCountry,
				&congressman.Jobtitle, &congressman.CatSocPro, &congressman.FamSocPro)

			if err != nil {
				cr.LogManager.WriteErrorLog("Erreur récupération du résultat " + err.Error())
				noError = false
			}

			congressmans = append(congressmans, congressman)
		}
		entities.Congressmans = congressmans
	}

	return &entities, noError
}

func (cr *CongressmanRepository) GetCongressman(id int) (*models.CongressmanModel, bool) {
	var congressman models.CongressmanModel
	db := cr.InitDB()
	noError := true

	row, err := db.Query("select * FROM PROCESSDEPUTES.CongressMan where CongressManId=?;", id)

	if err != nil {
		cr.LogManager.WriteErrorLog("Erreur requête " + err.Error())
		noError = false
	} else {
		if row.Next() {
			errScan := row.Scan(&congressman.Id, &congressman.Uid, &congressman.Civility, &congressman.FirstName,
				&congressman.LastName, &congressman.Alpha, &congressman.Trigramme, &congressman.BirthDate,
				&congressman.BirthCity, &congressman.BirthDepartment, &congressman.BirthCountry,
				&congressman.Jobtitle, &congressman.CatSocPro, &congressman.FamSocPro)

			if errScan != nil {
				cr.LogManager.WriteErrorLog("Erreur récupération du résultat " + errScan.Error())
				noError = false
			}
		}
		row.Close()
	}
	if congressman != (models.CongressmanModel{}) {
		return &congressman, noError
	} else {
		return nil, noError
	}
}

func (cr CongressmanRepository) GetById(id int) (*models.EntityModel, bool) {
	var entity models.EntityModel

	congressman, noError := cr.GetCongressman(id)

	if congressman != nil {
		entity.Congressman = *congressman
		return &entity, noError
	} else {
		return nil, noError
	}
}

func (cr *CongressmanRepository) InsertCongressMan(congressman *models.CongressmanModel) (int64, bool) {
	db := cr.InitDB()
	var lid int64
	noError := true

	if congressman == nil {
		cr.LogManager.WriteErrorLog("No Data send to insert")
	} else {
		queryCongressMan := "INSERT INTO PROCESSDEPUTES.Congressman(CongressManUid, Civility, FirstName, LastName, Alpha, Trigramme, BirthDate, BirthCity, BirthDepartment, BirthCountry, JobTitle, CatSocPro, FamSocPro) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)"
		stmt, errPrepare := db.Prepare(queryCongressMan)
		if errPrepare != nil {
			cr.LogManager.WriteErrorLog("Erreur récupération du résultat " + errPrepare.Error())
			noError = false
		} else {
			res, errExec := stmt.Exec(congressman.Uid, congressman.Civility, congressman.FirstName,
				congressman.LastName, congressman.Alpha, congressman.Trigramme, congressman.BirthDate,
				congressman.BirthCity, congressman.BirthDepartment, congressman.BirthCountry,
				congressman.Jobtitle, congressman.CatSocPro, congressman.FamSocPro)
			if errExec != nil {
				cr.LogManager.WriteErrorLog("Congressman Repository : Erreur exécution requête " + errExec.Error())
				noError = false
			} else {
				var errGetLastId error
				lid, errGetLastId = res.LastInsertId()
				if errGetLastId != nil {
					cr.LogManager.WriteErrorLog("Erreur lors de la récupération de l'id enregistré " + errGetLastId.Error())
				}
			}
		}
	}
	defer db.Close()

	return lid, noError
}

func (cr CongressmanRepository) CreateEntity(entity *models.EntityModel) (int64, bool) {
	lid, noError := cr.InsertCongressMan(&entity.Congressman)
	return lid, noError
}

func (cr *CongressmanRepository) UpdateCongressMan(congressman *models.CongressmanModel, id int) bool {
	db := cr.InitDB()
	noError := true

	queryCongressMan := "UPDATE  PROCESSDEPUTES.Congressman SET Civility=?, FirstName=?, LastName=?, Alpha=?, Trigramme=?, BirthDate=?, BirthCity=?, BirthDepartment=?, BirthCountry=?, JobTitle=?, CatSocPro=?, FamSocPro=? WHERE CongressManId = ?"
	stmt, errPrepare := db.Prepare(queryCongressMan)
	if errPrepare != nil {
		cr.LogManager.WriteErrorLog("Erreur récupération du résultat " + errPrepare.Error())
		noError = false
	} else {
		_, errExec := stmt.Exec(congressman.Civility, congressman.FirstName, congressman.LastName,
			congressman.Alpha, congressman.Trigramme, congressman.BirthDate,
			congressman.BirthCity, congressman.BirthDepartment, congressman.BirthCountry,
			congressman.Jobtitle, congressman.CatSocPro, congressman.FamSocPro, id)
		if errExec != nil {
			cr.LogManager.WriteErrorLog("Congressman Repository : Erreur exécution requête " + errExec.Error())
			noError = false
		}
	}

	defer db.Close()
	return noError
}

func (cr *CongressmanRepository) DeleteCongressMan(id int) (int64, bool) {
	var nbDelete int64
	db := cr.InitDB()
	noError := true

	queryCongressMan := "DELETE FROM PROCESSDEPUTES.Congressman WHERE CongressManId = ?"
	stmt, errPrepare := db.Prepare(queryCongressMan)
	if errPrepare != nil {
		cr.LogManager.WriteErrorLog("Erreur récupération du résultat " + errPrepare.Error())
		noError = false
	} else {
		result, errExec := stmt.Exec(id)
		if errExec != nil {
			cr.LogManager.WriteErrorLog("Congressman Repository : Erreur exécution requête " + errExec.Error())
			noError = false
		}
		nbDelete, _ = result.RowsAffected()
	}
	defer db.Close()

	return nbDelete, noError
}

func (cr CongressmanRepository) InitRepository() (IRepository, Manager.LogManager) {
	return nil, manager.LogManager{}
}
