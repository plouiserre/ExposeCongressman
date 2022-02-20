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

func (rb RepositoryBase) GetAll(model models.IModels) (*models.EntityModel, bool) {
	entities := models.EntityModel{}
	db := rb.InitDB()
	noError := true

	rows, err := model.GetQuery(db)

	if err != nil {
		rb.LogManager.WriteErrorLog("Erreur requête " + err.Error())
		noError = false
	} else {
		defer rows.Close()

		//TODO changer le renvoi de rowscangetentities plus tard
		entities, noError = model.RowsScanGetEntities(rows, rb.LogManager)
	}

	return &entities, noError
}

func (rb RepositoryBase) GetById(model models.IGetByIdEntity, id int) (*models.EntityModel, bool) {
	var entity models.EntityModel
	db := rb.InitDB()
	noError := true
	isEmpty := false

	row, err := model.QueryGetById(db, id)

	if err != nil {
		rb.LogManager.WriteErrorLog("Erreur requête " + err.Error())
		noError = false
	} else {
		if row.Next() {
			entity, noError = model.RowsScanGetById(row, rb.LogManager)
		} else {
			isEmpty = true
		}
		row.Close()
	}
	if !isEmpty {
		return &entity, noError
	} else {
		return nil, noError
	}
}

func (rb RepositoryBase) CreateEntity(model models.ICreateEntity, entity *models.EntityModel) (int64, bool) {
	db := rb.InitDB()
	var lid int64
	noError := true

	isEntityFill := model.IsEntityFill(*entity, rb.LogManager)

	if isEntityFill {
		stmt, noError := model.PrepareCreateQuery(db, rb.LogManager)

		if noError {
			res, nameRepository, errExec := model.ExecuteCreateQuery(stmt, *entity)
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

func (rb RepositoryBase) UpdateEntity(entity *models.EntityModel, id int) bool {
	return false
}

func (rb RepositoryBase) DeleteEntity(id int) (int64, bool) {
	return 0, false
}
