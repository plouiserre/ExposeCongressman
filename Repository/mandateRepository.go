package Repository

import (
	_ "github.com/go-sql-driver/mysql"
	manager "github.com/plouiserre/exposecongressman/Manager"
)

type MandateRepository struct {
	LogManager *manager.LogManager
}
