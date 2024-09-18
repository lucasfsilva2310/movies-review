package movieComments

import (
	"database/sql"

	Configuration "github.com/lucasfsilva2310/movies-review/internal/config"
	"github.com/lucasfsilva2310/movies-review/internal/errorHandlers"
)

type MovieCommentRepository struct {
	Repo *Configuration.Repository
}

func NewMovieCommentRepository(repo *Configuration.Repository) *MovieCommentRepository {
	return &MovieCommentRepository{
		Repo: repo,
	}
}

func (movieCommentRepo *MovieCommentRepository) Create(movieComment MovieComment) error {
	_, err := movieCommentRepo.Repo.DB.Exec(
		`
		INSERT INTO movie_comments
			(
				movie_id,
				user_id,		
			)
		`,
		movieComment.Movie_ID,
		movieComment.User_ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (movieCommentRepo *MovieCommentRepository) GetAll() ([]MovieCommentReturn, error) {
	var movieComments []MovieCommentReturn

	rows, err := movieCommentRepo.Repo.DB.Query(`
		SELECT 
			movie_id,
			user_id,
			comment
		FROM movie_comments
	`)
	if err != nil {
		return movieComments, err
	}

	for rows.Next() {
		var movieComment MovieCommentReturn
		if err := rows.Scan(
			&movieComment.Movie_ID,
			&movieComment.User_ID,
			&movieComment.Comment,
		); err != nil {
			return movieComments, err
		}
		movieComments = append(movieComments, movieComment)
	}

	return movieComments, nil
}

func (movieCommentRepo *MovieCommentRepository) GetAllByUserID(id int) ([]MovieCommentReturn, error) {
	var movieComments []MovieCommentReturn

	rows, err := movieCommentRepo.Repo.DB.Query(`
		SELECT 
			movie_id,
			user_id,
			comment
		FROM movie_comments
		WHERE user_id = $1
	`, id)
	if err != nil {
		return movieComments, err
	}

	for rows.Next() {
		var movieComment MovieCommentReturn
		if err := rows.Scan(
			&movieComment.Movie_ID,
			&movieComment.User_ID,
			&movieComment.Comment,
		); err != nil {
			return movieComments, err
		}
		movieComments = append(movieComments, movieComment)
	}

	return movieComments, nil
}

func (movieCommentRepo *MovieCommentRepository) GetAllByMovieID(id int) ([]MovieCommentReturn, error) {
	var movieComments []MovieCommentReturn

	rows, err := movieCommentRepo.Repo.DB.Query(`
		SELECT 
			movie_id,
			user_id,
			comment
		FROM movie_comments
		WHERE movie_id = $1
	`, id)
	if err != nil {
		return movieComments, err
	}

	for rows.Next() {
		var movieComment MovieCommentReturn
		if err := rows.Scan(
			&movieComment.Movie_ID,
			&movieComment.User_ID,
			&movieComment.Comment,
		); err != nil {
			return movieComments, err
		}
		movieComments = append(movieComments, movieComment)
	}

	return movieComments, nil
}

func (movieCommentRepo *MovieCommentRepository) Update(id int, username string, movieComment MovieCommentUpdate) error {
	var dbMovieComment = struct {
		ID       int
		User_ID  int
		Username string
	}{}

	err := movieCommentRepo.Repo.DB.QueryRow(`
	SELECT
		movie_comments.id,
		user_id,
		users.username
	FROM movie_comments
	JOIN users on users.id = movie_comments.user_id
	WHERE movie_comments.id = $1
	`, id).Scan(&dbMovieComment.ID, &dbMovieComment.User_ID, &dbMovieComment.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			return &errorHandlers.EntityNotFound{
				Entity: "MovieComment",
			}
		}
		return err
	}

	if dbMovieComment.Username != username {
		return &errorHandlers.NotSameUser{}
	}

	_, err = movieCommentRepo.Repo.DB.Exec(`
		UPDATE movie_comments
		SET comment = $1 
		WHERE id = $2
	`, movieComment.Comment, id)

	if err != nil {
		return err
	}

	return nil
}

func (movieCommentRepo *MovieCommentRepository) Delete(id int, username string) error {
	var dbMovieComment = struct {
		ID       int
		User_ID  int
		Username string
	}{}

	errorScanning := movieCommentRepo.Repo.DB.QueryRow(`
	SELECT
		movie_comments.id,
		user_id,
		users.username
	FROM movie_comments
	JOIN users on users.id = movie_comments.user_id
	WHERE movie_comments.id = $1
	`, id).Scan(&dbMovieComment.ID, &dbMovieComment.User_ID, &dbMovieComment.Username)

	if errorScanning != nil {
		if errorScanning == sql.ErrNoRows {
			return &errorHandlers.EntityNotFound{
				Entity: "MovieComment",
			}
		}
		return errorScanning
	}

	if dbMovieComment.Username != username {
		return &errorHandlers.NotSameUser{}
	}

	_, err := movieCommentRepo.Repo.DB.Exec(`
		DELETE FROM movie_comments
		WHERE id = $1
	`, id)

	if err != nil {
		return err
	}
	return nil
}
