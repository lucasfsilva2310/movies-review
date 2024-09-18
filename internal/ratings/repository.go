package ratings

import (
	"database/sql"

	"github.com/lib/pq"
	Configuration "github.com/lucasfsilva2310/movies-review/internal/config"
	"github.com/lucasfsilva2310/movies-review/internal/errorHandlers"
)

type RatingRepository struct {
	Repo *Configuration.Repository
}

func NewRatingRepository(repo *Configuration.Repository) *RatingRepository {
	return &RatingRepository{
		Repo: repo,
	}
}

func (ratingRepo *RatingRepository) Create(rating Rating) error {
	rowsMovies, errorMovies := ratingRepo.Repo.DB.Query("SELECT id FROM movies WHERE id = $1", rating.Movie_ID)

	if errorMovies != nil {
		return errorMovies
	}

	defer rowsMovies.Close()

	if !rowsMovies.Next() {
		return &errorHandlers.EntityNotFound{
			Entity: "Movie",
		}
	}

	rowsUser, errorUser := ratingRepo.Repo.DB.Query("SELECT id FROM users WHERE id = $1", rating.User_ID)

	if errorUser != nil {
		return errorUser
	}

	defer rowsUser.Close()

	if !rowsUser.Next() {
		return &errorHandlers.EntityNotFound{
			Entity: "User",
		}
	}

	_, err := ratingRepo.Repo.DB.Exec(
		`INSERT INTO ratings
		(
			score,
			movie_id,
			user_id
		) VALUES
		($1, $2, $3)
		`,
		rating.Score,
		rating.Movie_ID,
		rating.User_ID,
	)

	if err != nil {
		if err, ok := err.(*pq.Error); ok && err.Code.Name() == "unique_violation" {
			return &errorHandlers.AlreadyExistsError{
				Entity: "Rating",
			}
		}
		return err
	}
	return nil
}

func (ratingRepo *RatingRepository) GetAll() ([]RatingReturn, error) {
	var ratings []RatingReturn

	rows, err := ratingRepo.Repo.DB.Query("SELECT score, movie_id, user_id FROM ratings")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var rating RatingReturn
		if err := rows.Scan(&rating.Score, &rating.Movie_ID, &rating.User_ID); err != nil {
			return nil, err
		}
		ratings = append(ratings, rating)
	}
	return ratings, nil
}

func (ratingRepo *RatingRepository) GetAllByMovieID(id int) ([]RatingReturn, error) {
	var ratings []RatingReturn

	rows, err := ratingRepo.Repo.DB.Query(`
	SELECT
	 score, 
	 movie_id, 
	 user_id 
	FROM ratings
		WHERE movie_id = $1
	`, id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var rating RatingReturn
		if err := rows.Scan(&rating.Score, &rating.Movie_ID, &rating.User_ID); err != nil {
			return nil, err
		}
		ratings = append(ratings, rating)
	}
	return ratings, nil
}

func (ratingRepo *RatingRepository) GetAllByUserID(id int) ([]RatingReturn, error) {
	var ratings []RatingReturn

	rows, err := ratingRepo.Repo.DB.Query(`
	SELECT
	 score, 
	 movie_id, 
	 user_id 
	FROM ratings
		WHERE user_id = $1
	`, id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var rating RatingReturn
		if err := rows.Scan(&rating.Score, &rating.Movie_ID, &rating.User_ID); err != nil {
			return nil, err
		}
		ratings = append(ratings, rating)
	}
	return ratings, nil
}

func (ratingRepo *RatingRepository) Update(id int, username string, rating RatingUpdate) error {
	var dbRating = struct {
		ID       int
		User_ID  int
		Username string
	}{}

	err := ratingRepo.Repo.DB.QueryRow(`
	SELECT
		ratings.id,
		user_id,
		users.username
	FROM ratings
	JOIN users on users.id = ratings.user_id
	WHERE ratings.id = $1
	`, id).Scan(&dbRating.ID, &dbRating.User_ID, &dbRating.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			return &errorHandlers.EntityNotFound{
				Entity: "rating",
			}
		}
		return err
	}

	if dbRating.Username != username {
		return &errorHandlers.NotSameUser{}
	}

	_, err = ratingRepo.Repo.DB.Exec(`
		UPDATE ratings
		SET score = $1 
		WHERE id = $2
	`, rating.Score, id)

	if err != nil {
		return err
	}

	return nil
}

func (ratingRepo *RatingRepository) Delete(id int, username string) error {
	var dbRating = struct {
		ID       int
		User_ID  int
		Username string
	}{}

	errScanning := ratingRepo.Repo.DB.QueryRow(`
	SELECT
		ratings.id,
		user_id,
		users.username
	FROM ratings
	JOIN users on users.id = ratings.user_id
	WHERE ratings.id = $1
	`, id).Scan(&dbRating.ID, &dbRating.User_ID, &dbRating.Username)

	if errScanning != nil {
		if errScanning == sql.ErrNoRows {
			return &errorHandlers.EntityNotFound{
				Entity: "rating",
			}
		}
		return errScanning
	}

	if dbRating.Username != username {
		return &errorHandlers.NotSameUser{}
	}

	_, err := ratingRepo.Repo.DB.Exec("DELETE FROM ratings WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
