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

//TODO fusionner IGetByIdEntity ICreateEntity IUpdateEntity et IDeleteEntity dans une seule interface
type IGetByIdEntity interface {
	QueryGetById() string
	RowsScanGetById(rows *sql.Rows, logManager *manager.LogManager) (EntityModel, bool)
}

type ICreateEntity interface {
	IsEntityFill(entity EntityModel, logManager *manager.LogManager) bool
	QueryCreate() string
	ExecuteCreateQuery(stmt *sql.Stmt, model EntityModel) (sql.Result, string, error)
}

type IUpdateEntity interface {
	QueryUpdate() string
	ExecuteUpdateQuery(stmt *sql.Stmt, model EntityModel, id int64) (string, error)
}

type IDeleteEntity interface {
	QueryDelete() string
}
