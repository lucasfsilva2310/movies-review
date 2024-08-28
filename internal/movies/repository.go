package movies

import (
	"database/sql"

	"github.com/lib/pq"
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

	for rows.Next() {
		var movie Movie
		if err := rows.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.Tags, &movie.Platforms, &movie.CreatedAt, &movie.UpdatedAt); err != nil {
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
	_, err := repo.db.Exec(`
		INSERT INTO movies
	(title, description, release_date, tags, platforms, created_at, updated_at)
	  	VALUES 
	($1, $2, $3, $4, $5, $6, $7)`, movie.Title, movie.Description, movie.ReleaseDate, pq.Array(movie.Tags), pq.Array(movie.Platforms), movie.CreatedAt, movie.UpdatedAt)

	if err != nil {
		return Movie{}, err
	}

	return movie, nil
}
