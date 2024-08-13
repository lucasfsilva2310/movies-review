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
