package Repository

import (
	"database/sql"

	models "github.com/plouiserre/exposecongressman/Models"
)

//TODO si OK mettre toutes les méthodes communes dedans
type IRepository interface {
	InitDB() (db *sql.DB)
	//TODO modifier le retour de AllMandates
	GetAll() (*models.EntityModel, bool)
	//InitRepository() (IRepository, Manager.LogManager)
}
