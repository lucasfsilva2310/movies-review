package movies

import (
	"database/sql"
	"encoding/json"

	Configuration "github.com/lucasfsilva2310/movies-review/internal/config"
	"github.com/lucasfsilva2310/movies-review/internal/errorHandlers"
)

type MovieRepository struct {
	Repo *Configuration.Repository
}

func NewMovieRepository(repo *Configuration.Repository) *MovieRepository {
	return &MovieRepository{
		Repo: repo,
	}
}

func (movieRepo *MovieRepository) GetAll() ([]MovieReturn, error) {
	var movies []MovieReturn

	rows, err := movieRepo.Repo.DB.Query("SELECT * FROM movies WHERE active = true")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tagsJSON, platformsJSON []byte

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	for rows.Next() {
		var movie MovieReturn
		if err := rows.Scan(&movie.Title, &movie.Description, &movie.ReleaseDate, &tagsJSON, &platformsJSON); err != nil {
			return nil, err
		}

		if err := json.Unmarshal(tagsJSON, &movie.Tags); err != nil {
			return nil, err
		}

		if err := json.Unmarshal(platformsJSON, &movie.Platforms); err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	return movies, nil
}

func (movieRepo *MovieRepository) GetByID(id int) (MovieReturn, error) {
	var movie MovieReturn

	var tagsJSON, platformsJSON []byte

	err := movieRepo.Repo.DB.QueryRow(`
	SELECT
		title,
		description,
		release_date,
		tags,
		platforms
	FROM movies
	WHERE id = $1
	AND active = true
	`, id).Scan(
		&movie.Title,
		&movie.Description,
		&movie.ReleaseDate,
		&tagsJSON,
		&platformsJSON,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return MovieReturn{}, &errorHandlers.EntityNotFound{
				Entity: "Movie",
			}
		}
		return MovieReturn{}, err
	}

	if err := json.Unmarshal(tagsJSON, &movie.Tags); err != nil {
		return MovieReturn{}, err
	}

	if err := json.Unmarshal(platformsJSON, &movie.Platforms); err != nil {
		return MovieReturn{}, err
	}

	return movie, nil
}

func (movieRepo *MovieRepository) Create(movie MovieReturn) error {
	var existingMovie Movie

	err := movieRepo.Repo.DB.QueryRow("SELECT id FROM movies WHERE LOWER(title) = LOWER($1)", movie.Title).Scan(&existingMovie.ID)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if existingMovie.ID != 0 {
		return &errorHandlers.AlreadyExistsError{
			Entity: "Movie",
		}
	}

	tagsJSON, err := json.Marshal(movie.Tags)
	if err != nil {
		return err
	}

	platformsJSON, err := json.Marshal(movie.Platforms)
	if err != nil {
		return err
	}

	_, err = movieRepo.Repo.DB.Exec(`
		INSERT INTO movies
	(title, description, release_date, tags, platforms)
	  	VALUES 
	($1, $2, $3, $4, $5)`, movie.Title, movie.Description, movie.ReleaseDate, tagsJSON, platformsJSON)

	if err != nil {
		return err
	}

	return nil
}

func (movieRepo *MovieRepository) Delete(id int) error {
	_, err := movieRepo.Repo.DB.Exec("UPDATE movies SET active = false WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
