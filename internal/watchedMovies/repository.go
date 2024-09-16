package watchedMovies

import Configuration "github.com/lucasfsilva2310/movies-review/internal/config"

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
		WHERE active = true
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
		AND active = true
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
		AND active = true
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

func (watchedMovieRepo *WatchedMovieRepository) Delete(id int) error {
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
