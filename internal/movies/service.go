package movies

type Service struct {
	repository Repository
}

func NewService(Repository *Repository) *Service {
	return &Service{
		repository: *Repository,
	}
}

func (service *Service) GetAll() ([]Movie, error) {
	return service.repository.GetAll()
}

func (service *Service) GetByID(id string) (Movie, error) {
	return service.repository.GetByID(id)
}

func (service *Service) Create(movie Movie) error {
	return service.repository.Create(movie)
}

func (service *Service) Delete(id string) error {
	return service.repository.Delete(id)
}
