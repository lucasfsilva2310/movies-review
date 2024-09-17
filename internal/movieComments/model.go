package movieComments

type MovieComment struct {
	ID        int    `json:"id"`
	Movie_ID  int    `json:"movie_id"`
	User_ID   int    `json:"user_id"`
	Comment   string `json:"comment"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
}

type MovieCommentReturn struct {
	Comment  string `json:"comment"`
	User_ID  int    `json:"user_id"`
	Movie_ID int    `json:"movie_id"`
}

type MovieCommentUpdate struct {
	Comment string `json:"comment"`
}
