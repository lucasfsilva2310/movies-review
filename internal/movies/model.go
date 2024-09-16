package movies

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

type Platforms string

const (
	Netflix    Platforms = "Netflix"
	Amazon     Platforms = "Amazon"
	Hulu       Platforms = "Hulu"
	Apple      Platforms = "Apple"
	DisneyPlus Platforms = "Disney+"
	Other      Platforms = "Other"
)

type Movie struct {
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	ReleaseDate string      `json:"releaseDate"`
	Tags        []Tags      `json:"tags"`
	Platforms   []Platforms `json:"platforms"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
	Active      bool        `json:"active"`
}

type MovieReturn struct {
	Title       string      `json:"title"`
	Description string      `json:"description"`
	ReleaseDate string      `json:"releaseDate"`
	Tags        []Tags      `json:"tags"`
	Platforms   []Platforms `json:"platforms"`
}
