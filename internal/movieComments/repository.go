package movieComments

import Configuration "github.com/lucasfsilva2310/movies-review/internal/config"

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

func (movieCommentRepo *MovieCommentRepository) Delete(id int) error {
	_, err := movieCommentRepo.Repo.DB.Exec(`
		DELETE FROM movie_comments
		WHERE id = $1
	`, id)

	if err != nil {
		return err
	}
	return nil
}
