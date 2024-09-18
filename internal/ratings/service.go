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

func (service *RatingService) UpdateRating(id int, username string, rating RatingUpdate) error {
	return service.ratingRepository.Update(id, username, rating)
}

func (service *RatingService) Delete(id int, username string) error {
	return service.ratingRepository.Delete(id, username)
}
