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
