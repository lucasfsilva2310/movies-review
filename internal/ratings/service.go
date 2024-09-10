package ratings

type RatingService struct {
	ratingRepository RatingRepository
}

func NewRatingService(Repository *RatingRepository) *RatingService {
	return &RatingService{
		ratingRepository: *Repository,
	}
}

func (service *RatingService) Create(rating Rating) error {
	return service.ratingRepository.Create(rating)
}

func (service *RatingService) GetAll() ([]RatingReturn, error) {
	return service.ratingRepository.GetAll()
}

func (service *RatingService) GetAllByMovieID(id int) ([]RatingReturn, error) {
	return service.ratingRepository.GetAllByMovieID(id)
}

func (service *RatingService) GetAllByUserID(id int) ([]RatingReturn, error) {
	return service.ratingRepository.GetAllByUserID(id)
}
