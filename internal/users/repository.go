package users

import (
	"database/sql"

	"github.com/lib/pq"
	Configuration "github.com/lucasfsilva2310/movies-review/internal/config"
	"github.com/lucasfsilva2310/movies-review/internal/errorHandlers"
)

type UserRepository struct {
	Repo *Configuration.Repository
}

func NewUserRepository(repo *Configuration.Repository) *UserRepository {
	return &UserRepository{
		Repo: repo,
	}
}

func (userRepo *UserRepository) Create(user User) error {
	_, err := userRepo.Repo.DB.Exec(
		`INSERT INTO users
		 (
			username, 
			password, 
			full_name, 
			email
		) VALUES 
		 ($1, $2, $3, $4)`,
		user.Username,
		user.Password,
		user.FullName,
		user.Email)

	if err != nil {
		if err, ok := err.(*pq.Error); ok && err.Code.Name() == "unique_violation" {
			return &errorHandlers.AlreadyExistsError{
				Entity: "User",
			}
		}
		return err
	}
	return nil
}

func (userRepo *UserRepository) GetById(id int) (UserReturn, error) {
	var user UserReturn

	err := userRepo.Repo.DB.QueryRow(`
	SELECT 
		username, 
		full_name, 
		email 
	FROM users 
	WHERE id = $1`, id).Scan(
		&user.Username,
		&user.FullName,
		&user.Email,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return UserReturn{}, &errorHandlers.EntityNotFound{
				Entity: "User",
			}
		}
		return UserReturn{}, err
	}

	return user, nil
}

func (userRepo *UserRepository) GetByUsername(username string) (UserReturn, error) {
	var user UserReturn

	err := userRepo.Repo.DB.QueryRow(`
	SELECT 
		username, 
		full_name, 
		email 
	FROM users 
	WHERE username = $1`, username).Scan(
		&user.Username,
		&user.FullName,
		&user.Email,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return UserReturn{}, &errorHandlers.EntityNotFound{
				Entity: "User",
			}
		}
		return UserReturn{}, err
	}

	return user, nil
}

func (userRepo *UserRepository) DeleteById(id int) error {
	_, err := userRepo.Repo.DB.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (userRepo *UserRepository) DeleteByUsername(username string) error {
	_, err := userRepo.Repo.DB.Exec("DELETE FROM users WHERE username = $1", username)
	if err != nil {
		return err
	}
	return nil
}
