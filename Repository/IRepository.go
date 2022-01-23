package Repository

import "database/sql"

//TODO si OK mettre toutes les méthodes communes dedans
type IRepository interface {
	InitDB() (db *sql.DB)
}
