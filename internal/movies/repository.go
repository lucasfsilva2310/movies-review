package movies

import "database/sql"

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
