package movies

import "database/sql"

type Repository struct {
	db *sql.DB
}
