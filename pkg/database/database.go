package databases

import (
	"database/sql"
	"log"

	Configuration "github.com/lucasfsilva2310/movies-review/internal/config"
)

func ConnectDB(config *Configuration.DatabaseConfig) (*sql.DB, error) {
	log.Println("Connecting to Database..")

	dbConnection, err := sql.Open("postgres", config.DatabaseUrl)

	if err != nil {
		return nil, err
	}

	if err := dbConnection.Ping(); err != nil {
		return nil, err
	}
	log.Println("Connected successfully!")

	return dbConnection, nil
}
