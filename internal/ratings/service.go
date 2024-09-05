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
