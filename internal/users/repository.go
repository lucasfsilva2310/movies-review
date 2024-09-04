package users

import (
	Configuration "github.com/lucasfsilva2310/movies-review/internal/config"
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
		return err
	}
	return nil
}
