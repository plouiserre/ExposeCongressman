package Repository

import (
	"database/sql"

	models "github.com/plouiserre/exposecongressman/Models"
)

type IRepository interface {
	InitDB() (db *sql.DB)
	GetAll() (*models.EntityModel, bool)
	GetById(id int) (*models.EntityModel, bool)
	CreateEntity(*models.EntityModel) (int64, bool)
	UpdateEntity(mandate *models.EntityModel, id int) bool
	DeleteEntity(id int) (int64, bool)
}
