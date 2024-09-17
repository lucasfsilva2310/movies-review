package movieComments

type MovieCommentService struct {
	movieCommentRepository MovieCommentRepository
}

func NewMovieCommentService(Repository *MovieCommentRepository) *MovieCommentService {
	return &MovieCommentService{
		movieCommentRepository: *Repository,
	}
}

func (service *MovieCommentService) Create(movieComment MovieComment) error {
	return service.movieCommentRepository.Create(movieComment)
}

func (service *MovieCommentService) GetAll() ([]MovieCommentReturn, error) {
	return service.movieCommentRepository.GetAll()
}

func (service *MovieCommentService) GetAllByUserID(id int) ([]MovieCommentReturn, error) {
	return service.movieCommentRepository.GetAllByUserID(id)
}

func (service *MovieCommentService) GetAllByMovieID(id int) ([]MovieCommentReturn, error) {
	return service.movieCommentRepository.GetAllByMovieID(id)
}

func (service *MovieCommentService) UpdateMovieComment(id int, username string, movieComment MovieCommentUpdate) error {
	return service.movieCommentRepository.Update(id, username, movieComment)
}

func (service *MovieCommentService) Delete(id int) error {
	return service.movieCommentRepository.Delete(id)
}
