package movies

type MovieService struct {
	movieRepository MovieRepository
}

func NewMovieService(Repository *MovieRepository) *MovieService {
	return &MovieService{
		movieRepository: *Repository,
	}
}

func (service *MovieService) GetAll() ([]Movie, error) {
	return service.movieRepository.GetAll()
}

func (service *MovieService) GetByID(id int) (MovieReturn, error) {
	return service.movieRepository.GetByID(id)
}

func (service *MovieService) Create(movie MovieReturn) error {
	return service.movieRepository.Create(movie)
}

func (service *MovieService) Delete(id int) error {
	return service.movieRepository.Delete(id)
}
