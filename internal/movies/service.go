package movies

type MovieService struct {
	movieRepository MovieRepository
}

func NewService(Repository *MovieRepository) *MovieService {
	return &MovieService{
		movieRepository: *Repository,
	}
}

func (service *MovieService) GetAll() ([]Movie, error) {
	return service.movieRepository.GetAll()
}

func (service *MovieService) GetByID(id string) (Movie, error) {
	return service.movieRepository.GetByID(id)
}

func (service *MovieService) Create(movie Movie) error {
	return service.movieRepository.Create(movie)
}

func (service *MovieService) Delete(id string) error {
	return service.movieRepository.Delete(id)
}
