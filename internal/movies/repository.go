package movies

import (
	"database/sql"
	"encoding/json"

	Configuration "github.com/lucasfsilva2310/movies-review/internal/config"
)

type MovieRepository struct {
	Repo *Configuration.Repository
}

func NewMovieRepository(repo *Configuration.Repository) *MovieRepository {
	return &MovieRepository{
		Repo: repo,
	}
}

type AlreadyExistsError struct{}

func (alreadyExistsError *AlreadyExistsError) Error() string {
	return "Movie Already Exists."
}

func (movieRepo *MovieRepository) GetAll() ([]Movie, error) {
	var movies []Movie

	rows, err := movieRepo.Repo.DB.Query("SELECT * FROM movies")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tagsJSON, platformsJSON []byte

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	for rows.Next() {
		var movie Movie
		if err := rows.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.ReleaseDate, &tagsJSON, &platformsJSON, &movie.CreatedAt, &movie.UpdatedAt); err != nil {
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

func (movieRepo *MovieRepository) GetByID(id string) (Movie, error) {
	var movie Movie

	err := movieRepo.Repo.DB.QueryRow("SELECT * FROM movies WHERE id = $1", id).Scan(
		&movie.ID,
		&movie.Title,
		&movie.Description,
		&movie.ReleaseDate,
		&movie.Tags,
		&movie.Platforms,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return Movie{}, nil
		}
		return Movie{}, err
	}

	return movie, nil
}

func (movieRepo *MovieRepository) Create(movie Movie) error {
	var existingMovie Movie

	err := movieRepo.Repo.DB.QueryRow("SELECT id FROM movies WHERE LOWER(title) = LOWER($1)", movie.Title).Scan(&existingMovie.ID)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if existingMovie.ID != "" {
		return &AlreadyExistsError{}
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
	(title, description, release_date, tags, platforms, created_at, updated_at)
	  	VALUES 
	($1, $2, $3, $4, $5, $6, $7)`, movie.Title, movie.Description, movie.ReleaseDate, tagsJSON, platformsJSON, movie.CreatedAt, movie.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (movieRepo *MovieRepository) Delete(id string) error {
	_, err := movieRepo.Repo.DB.Exec("DELETE FROM movies WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
