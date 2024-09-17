package ratings

type Rating struct {
	ID        int `json:"id"`
	Score     int `json:"score"`
	Movie_ID  int `json:"movie_id"`
	User_ID   int `json:"user_id"`
	CreatedAt int `json:"created_at"`
	UpdatedAt int `json:"updated_at"`
}

type RatingReturn struct {
	Score    int `json:"score"`
	Movie_ID int `json:"movie_id"`
	User_ID  int `json:"user_id"`
}

type RatingUpdate struct {
	Score int `json:"score"`
}
