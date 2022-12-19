package models

import "fmt"

type Movie struct {
	ID          int64  `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	ReleaseDate string `db:"release_date"`
	Duration    int64  `db:"duration"`
	TrailerURL  string `db:"trailer_url"`
}

type JsonMovie struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ReleaseDate string `json:"releaseDate"`
	Duration    int64  `json:"duration"`
	TrailerURL  string `json:"trailerUrl"`
}

func (m Movie) String() string {
	return fmt.Sprintf("Movie (%d): %s", m.ID, m.Title)
}

func MapMovieToJSON(m Movie) JsonMovie {
	return JsonMovie(m)
}

// Service layer
func (store *DatabaseStore) GetMovie(id int) (*Movie, error) {
	var movie Movie
	err := store.db.Get(&movie, "SELECT * FROM movies WHERE id = $1", id)

	if err != nil {
		return nil, err
	}
	return &movie, nil
}

func (store *DatabaseStore) GetMovies() ([]*Movie, error) {
	var movies []*Movie
	err := store.db.Select(&movies, "SELECT * FROM movies")
	return movies, err
}
