package movie

import "time"

type Tags string

const (
	Comedy  Tags = "Comedy"
	Action  Tags = "Action"
	Drama   Tags = "Drama"
	Horror  Tags = "Horror"
	SciFi   Tags = "Sci-Fi"
	Romance Tags = "Romance"
	Musical Tags = "Musical"
	Sports  Tags = "Sports"
	Fantasy Tags = "Fantasy"
)

type Movie struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ReleaseDate string    `json:"releaseDate"`
	Tags        []Tags    `json:"tags"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
