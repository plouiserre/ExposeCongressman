package Repository

import (
	"database/sql"
	"time"

	manager "github.com/plouiserre/exposecongressman/Manager"
	models "github.com/plouiserre/exposecongressman/Models"
)

type RepositoryBase struct {
	LogManager *manager.LogManager
}

func (rb RepositoryBase) InitDB() (db *sql.DB) {
	db, err := sql.Open("mysql", "ProcessDeputesData:ASimpleP@ssW0rd@/PROCESSDEPUTES")

	if err != nil {
		rb.LogManager.WriteErrorLog("Erreur connexion " + err.Error())
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return db
}

func (rb RepositoryBase) GetAll(model models.IModels) (*models.IModels, bool) {
	db := rb.InitDB()
	noError := true
	var entities models.IModels

	rows, err := model.GetQuery(db)

	if err != nil {
		rb.LogManager.WriteErrorLog("Erreur requête " + err.Error())
		noError = false
	} else {
		defer rows.Close()

		entities, noError = model.RowsScanGetEntities(rows, rb.LogManager)
	}

	return &entities, noError
}

func (rb RepositoryBase) GetById(model models.IModel, id int) (*models.IModel, bool) {
	db := rb.InitDB()
	noError := true
	isEmpty := false

	query := model.QueryGetById()

	row, err := db.Query(query, id)

	if err != nil {
		rb.LogManager.WriteErrorLog("Erreur requête " + err.Error())
		noError = false
	} else {
		if row.Next() {
			model, noError = model.RowsScanGetById(row, rb.LogManager)
		} else {
			isEmpty = true
		}
		row.Close()
	}
	if !isEmpty {
		return &model, noError
	} else {
		return nil, noError
	}
}

func (rb RepositoryBase) CreateEntity(entity *models.IModel) (int64, bool) {
	db := rb.InitDB()
	model := *entity
	var lid int64
	noError := true

	isEntityFill := model.IsEntityFill(*entity, rb.LogManager)

	if isEntityFill {
		queryCreate := model.QueryCreate()
		stmt, errPrepare := db.Prepare(queryCreate)
		if errPrepare != nil {
			rb.LogManager.WriteErrorLog("Erreur récupération du résultat " + errPrepare.Error())
			noError = false
		} else {
			res, nameRepository, errExec := model.ExecuteCreateQuery(*entity, stmt)
			if errExec != nil {
				rb.LogManager.WriteErrorLog(nameRepository + " Repository : Erreur exécution requête " + errExec.Error())
				noError = false
			} else {
				var errGetLastId error
				lid, errGetLastId = res.LastInsertId()
				if errGetLastId != nil {
					rb.LogManager.WriteErrorLog("Erreur lors de la récupération de l'id enregistré " + errGetLastId.Error())
				}
			}
		}
	}
	defer db.Close()

	return lid, noError
}

func (rb RepositoryBase) UpdateEntity(entity *models.IModel, id int64) (int64, bool) {
	db := rb.InitDB()
	model := *entity
	noError := true

	queryUpdate := model.QueryUpdate()
	stmt, errPrepare := db.Prepare(queryUpdate)
	if errPrepare != nil {
		rb.LogManager.WriteErrorLog("Erreur récupération du résultat " + errPrepare.Error())
		noError = false
	} else {
		nameRepository, errExec := model.ExecuteUpdateQuery(stmt, id)
		if errExec != nil {
			rb.LogManager.WriteErrorLog(nameRepository + "Repository : Erreur exécution requête " + errExec.Error())
			noError = false
		}
	}

	defer db.Close()
	return id, noError
}

func (rb RepositoryBase) DeleteEntity(model models.IModel, id int) (int64, bool) {
	var nbDelete int64
	db := rb.InitDB()
	noError := true

	queryDelete := model.QueryDelete()
	stmt, errPrepare := db.Prepare(queryDelete)
	if errPrepare != nil {
		rb.LogManager.WriteErrorLog("Erreur récupération du résultat " + errPrepare.Error())
		noError = false
	} else {
		result, errExec := stmt.Exec(id)
		if errExec != nil {
			//TODO corriger ce log
			rb.LogManager.WriteErrorLog("Congressman Repository : Erreur exécution requête " + errExec.Error())
			noError = false
		}
		nbDelete, _ = result.RowsAffected()
	}
	defer db.Close()

	return nbDelete, noError
}
