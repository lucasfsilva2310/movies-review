package watchedMovies

type WatchedMovieService struct {
	watchedMovieRepository WatchedMovieRepository
}

func NewWatchedMovieService(Repository *WatchedMovieRepository) *WatchedMovieService {
	return &WatchedMovieService{
		watchedMovieRepository: *Repository,
	}
}

func (service *WatchedMovieService) Create(watchedMovie WatchedMovie) error {
	return service.watchedMovieRepository.Create(watchedMovie)
}

func (service *WatchedMovieService) GetAll() ([]WatchedMovieReturn, error) {
	return service.watchedMovieRepository.GetAll()
}

func (service *WatchedMovieService) GetAllByUserID(id int) ([]WatchedMovieReturn, error) {
	return service.watchedMovieRepository.GetAllByUserID(id)
}

func (service *WatchedMovieService) GetAllByMovieID(id int) ([]WatchedMovieReturn, error) {
	return service.watchedMovieRepository.GetAllByMovieID(id)
}
