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

func (rb RepositoryBase) GetById(id int) (*models.EntityModel, bool) {
	return nil, false
}

func (rb RepositoryBase) CreateEntity(entity *models.EntityModel) (int64, bool) {
	return 0, false
}

func (rb RepositoryBase) UpdateEntity(entity *models.EntityModel, id int) bool {
	return false
}

func (rb RepositoryBase) DeleteEntity(id int) (int64, bool) {
	return 0, false
}
