package watchedMovies

type WatchedMovie struct {
	ID        int  `json:"id"`
	Movie_ID  int  `json:"movie_id"`
	User_ID   int  `json:"user_id"`
	Watched   bool `json:"watched"`
	CreatedAt int  `json:"created_at"`
	UpdatedAt int  `json:"updated_at"`
	Active    bool `json:"active"`
}

type WatchedMovieReturn struct {
	Movie_ID int  `json:"movie_id"`
	User_ID  int  `json:"user_id"`
	Watched  bool `json:"watched"`
}
