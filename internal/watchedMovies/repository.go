package watchedMovies

import (
	"database/sql"

	Configuration "github.com/lucasfsilva2310/movies-review/internal/config"
	"github.com/lucasfsilva2310/movies-review/internal/errorHandlers"
)

type WatchedMovieRepository struct {
	Repo *Configuration.Repository
}

func NewWatchedMovieRepository(repo *Configuration.Repository) *WatchedMovieRepository {
	return &WatchedMovieRepository{
		Repo: repo,
	}
}

func (watchedMovieRepo *WatchedMovieRepository) Create(watchedMovie WatchedMovie) error {
	_, err := watchedMovieRepo.Repo.DB.Exec(
		`
		INSERT INTO watched_movies
			(
				movie_id,
				user_id,		
			)
		`,
		watchedMovie.Movie_ID,
		watchedMovie.User_ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (watchedMovieRepo *WatchedMovieRepository) GetAll() ([]WatchedMovieReturn, error) {
	var watchedMovies []WatchedMovieReturn

	rows, err := watchedMovieRepo.Repo.DB.Query(`
		SELECT 
			movie_id,
			user_id,
			watched
		FROM watched_movies
	`)
	if err != nil {
		return watchedMovies, err
	}

	for rows.Next() {
		var watchedMovie WatchedMovieReturn
		if err := rows.Scan(
			&watchedMovie.Movie_ID,
			&watchedMovie.User_ID,
			&watchedMovie.Watched,
		); err != nil {
			return watchedMovies, err
		}
		watchedMovies = append(watchedMovies, watchedMovie)
	}

	return watchedMovies, nil
}

func (watchedMovieRepo *WatchedMovieRepository) GetAllByUserID(id int) ([]WatchedMovieReturn, error) {
	var watchedMovies []WatchedMovieReturn

	rows, err := watchedMovieRepo.Repo.DB.Query(`
		SELECT 
			movie_id,
			user_id,
			watched
		FROM watched_movies
		WHERE user_id = $1
	`, id)
	if err != nil {
		return watchedMovies, err
	}

	for rows.Next() {
		var watchedMovie WatchedMovieReturn
		if err := rows.Scan(
			&watchedMovie.Movie_ID,
			&watchedMovie.User_ID,
			&watchedMovie.Watched,
		); err != nil {
			return watchedMovies, err
		}
		watchedMovies = append(watchedMovies, watchedMovie)
	}

	return watchedMovies, nil
}

func (watchedMovieRepo *WatchedMovieRepository) GetAllByMovieID(id int) ([]WatchedMovieReturn, error) {
	var watchedMovies []WatchedMovieReturn

	rows, err := watchedMovieRepo.Repo.DB.Query(`
		SELECT 
			movie_id,
			user_id,
			watched
		FROM watched_movies
		WHERE movie_id = $1
	`, id)
	if err != nil {
		return watchedMovies, err
	}

	for rows.Next() {
		var watchedMovie WatchedMovieReturn
		if err := rows.Scan(
			&watchedMovie.Movie_ID,
			&watchedMovie.User_ID,
			&watchedMovie.Watched,
		); err != nil {
			return watchedMovies, err
		}
		watchedMovies = append(watchedMovies, watchedMovie)
	}

	return watchedMovies, nil
}

func (watchedMovieRepo *WatchedMovieRepository) Update(id int, username string) error {
	var dbWatchedMovie = struct {
		ID       int
		watched  bool
		User_ID  int
		Username string
	}{}

	err := watchedMovieRepo.Repo.DB.QueryRow(`
	SELECT
		watched_movies.id,
		watched,
		user_id,
		users.username
	FROM watched_movies
	JOIN users on users.id = watched_movies.user_id
	WHERE watched_movies.id = $1
	`, id).Scan(&dbWatchedMovie.ID, &dbWatchedMovie.watched, &dbWatchedMovie.User_ID, &dbWatchedMovie.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			return &errorHandlers.EntityNotFound{
				Entity: "WatchedMovie",
			}
		}
		return err
	}

	if dbWatchedMovie.Username != username {
		return &errorHandlers.NotSameUser{}
	}

	_, err = watchedMovieRepo.Repo.DB.Exec(`
		UPDATE watched_movies
		SET watched = $1 
		WHERE id = $2
	`, !dbWatchedMovie.watched, id)

	if err != nil {
		return err
	}

	return nil
}

func (watchedMovieRepo *WatchedMovieRepository) Delete(id int, username string) error {
	var dbWatchedMovie = struct {
		ID       int
		watched  bool
		User_ID  int
		Username string
	}{}

	errorScanning := watchedMovieRepo.Repo.DB.QueryRow(`
	SELECT
		watched_movies.id,
		watched,
		user_id,
		users.username
	FROM watched_movies
	JOIN users on users.id = watched_movies.user_id
	WHERE watched_movies.id = $1
	`, id).Scan(&dbWatchedMovie.ID, &dbWatchedMovie.watched, &dbWatchedMovie.User_ID, &dbWatchedMovie.Username)

	if errorScanning != nil {
		if errorScanning == sql.ErrNoRows {
			return &errorHandlers.EntityNotFound{
				Entity: "WatchedMovie",
			}
		}
		return errorScanning
	}

	if dbWatchedMovie.Username != username {
		return &errorHandlers.NotSameUser{}
	}

	_, err := watchedMovieRepo.Repo.DB.Exec(
		`
		DELETE FROM watched_movies
		WHERE id = $1
		`,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}
