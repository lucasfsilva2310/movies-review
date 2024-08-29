package movies

import (
	"database/sql"
	"encoding/json"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (repo *Repository) GetAll() ([]Movie, error) {
	rows, err := repo.db.Query("SELECT * FROM movies")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var movies []Movie
	var tagsJSON, platformsJSON []byte

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

func (repo *Repository) GetByID(id string) (Movie, error) {
	var movie Movie

	err := repo.db.QueryRow("SELECT * FROM movies WHERE id = $1", id).Scan(
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

func (repo *Repository) CreateMovie(movie Movie) (Movie, error) {
	tagsJSON, err := json.Marshal(movie.Tags)
	if err != nil {
		return Movie{}, err
	}

	platformsJSON, err := json.Marshal(movie.Platforms)
	if err != nil {
		return Movie{}, err
	}

	_, err = repo.db.Exec(`
		INSERT INTO movies
	(title, description, release_date, tags, platforms, created_at, updated_at)
	  	VALUES 
	($1, $2, $3, $4, $5, $6, $7)`, movie.Title, movie.Description, movie.ReleaseDate, tagsJSON, platformsJSON, movie.CreatedAt, movie.UpdatedAt)

	if err != nil {
		return Movie{}, err
	}

	return movie, nil
}
