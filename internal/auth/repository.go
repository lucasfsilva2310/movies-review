package auth

import (
	"database/sql"
	"errors"

	Configuration "github.com/lucasfsilva2310/movies-review/internal/config"
)

type AuthRepository struct {
	Repo *Configuration.Repository
}

func NewAuthRepository(repo *Configuration.Repository) *AuthRepository {
	return &AuthRepository{
		Repo: repo,
	}
}

func (authRepo *AuthRepository) GetUserByUsername(username string) (string, string, error) {
	var storedPassword string
	var storedUsername string

	err := authRepo.Repo.DB.QueryRow("SELECT username, password FROM users WHERE username = $1", username).Scan(&storedUsername, &storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", "", errors.New("user not found")
		}
		return "", "", err
	}

	return storedUsername, storedPassword, nil
}
