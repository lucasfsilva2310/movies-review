package databases

import (
	"database/sql"

	Configuration "github.com/lucasfsilva2310/movies-review/internal/config"
)

func ConnectPostgres(config *Configuration.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", config.PostgresURI)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
