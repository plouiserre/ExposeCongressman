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
type IModel interface {
	QueryGetById() string
	RowsScanGetById(rows *sql.Rows, logManager *manager.LogManager) (EntityModel, bool)
	IsEntityFill(entity EntityModel, logManager *manager.LogManager) bool
	QueryCreate() string
	ExecuteCreateQuery(stmt *sql.Stmt, model EntityModel) (sql.Result, string, error)
	QueryUpdate() string
	ExecuteUpdateQuery(stmt *sql.Stmt, model EntityModel, id int64) (string, error)
	QueryDelete() string
}
