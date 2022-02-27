package Models

import (
	"database/sql"

	manager "github.com/plouiserre/exposecongressman/Manager"
)

//TODO changer le nom de cette interface et le nom des méthodes
type IModels interface {
	GetQuery(db *sql.DB) (*sql.Rows, error)
	RowsScanGetEntities(rows *sql.Rows, logManager *manager.LogManager) (IModels, bool)
}

type IModel interface {
	QueryGetById() string
	RowsScanGetById(rows *sql.Rows, logManager *manager.LogManager) (IModel, bool)
	IsEntityFill(model IModel, logManager *manager.LogManager) bool
	QueryCreate() string
	ExecuteCreateQuery(model IModel, stmt *sql.Stmt) (sql.Result, string, error)
	QueryUpdate() string
	ExecuteUpdateQuery(stmt *sql.Stmt, id int64) (string, error)
	QueryDelete() string
}
