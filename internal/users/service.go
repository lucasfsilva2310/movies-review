package users

type UserService struct {
	userRepository UserRepository
}

func NewUserService(Repository *UserRepository) *UserService {
	return &UserService{
		userRepository: *Repository,
	}
}

func (service *UserService) Create(user User) error {
	return service.userRepository.Create(user)
}

func (service *UserService) GetById(id int) (UserReturn, error) {
	return service.userRepository.GetById(id)
}

func (service *UserService) GetByUsername(username string) (UserReturn, error) {
	return service.userRepository.GetByUsername(username)
}

func (service *UserService) DeleteById(id int) error {
	return service.userRepository.DeleteById(id)
}

func (service *UserService) DeleteByUsername(username string) error {
	return service.userRepository.DeleteByUsername(username)
}
