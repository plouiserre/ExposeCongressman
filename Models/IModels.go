package Models

import (
	"database/sql"

	manager "github.com/plouiserre/exposecongressman/Manager"
)

//TODO changer le nom de cette interface et le nom des méthodes
type IModels interface {
	GetQuery(db *sql.DB) (*sql.Rows, error)
	RowsScanGetEntities(rows *sql.Rows, logManager *manager.LogManager) (EntityModel, bool)
}

type IGetByIdEntity interface {
	QueryGetById(db *sql.DB, id int) (*sql.Rows, error)
	RowsScanGetById(rows *sql.Rows, logManager *manager.LogManager) (EntityModel, bool)
}
