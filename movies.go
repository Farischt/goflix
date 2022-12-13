package main

import (
	"goflix/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// TODO: Extract controller & service in separated files
// Controller layer
func (s Server) handleGetMovies() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movies, err := s.Store.GetMovies()
		if err != nil {
			customErr := NewError("internal server error while retrieving movies", http.StatusInternalServerError)
			s.respondWithError(w, r, customErr, log.Error)
			return
		}

		if len(movies) == 0 {
			customErr := NewError("movies not found", http.StatusNotFound)
			s.respondWithError(w, r, customErr, log.Warn)
			return
		}
		// decode movies
		var jsonMovies []models.JsonMovie
		for _, movie := range movies {
			jsonMovies = append(jsonMovies, models.MapMovieToJSON(*movie))
		}

		// send response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		s.respondWithJSON(w, r, jsonMovies, http.StatusOK)
	}
}

func (s Server) handleGetMovie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			customErr := NewError("missing id parameter", http.StatusBadRequest)
			s.respondWithError(w, r, customErr, log.Warn)
			return
		}

		// Error omitted here since err = nil is equal to movie == nil
		movie, _ := s.Store.GetMovie(id)
		// if err != nil {
		// 	customErr := NewError("internal server error while retrieving the movie", http.StatusInternalServerError)
		// 	s.respondWithError(w, r, customErr, log.Error)
		// 	return
		// }

		if movie == nil {
			customErr := NewError("movie not found", http.StatusNotFound)
			s.respondWithError(w, r, customErr, log.Warn)
			return
		}

		jsonMovie := models.MapMovieToJSON(*movie)
		s.respondWithJSON(w, r, jsonMovie, http.StatusOK)
	}
}

// Service layer
func (store *DatabaseStore) GetMovie(id int) (*models.Movie, error) {
	var movie models.Movie
	err := store.db.Get(&movie, "SELECT * FROM movies WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &movie, nil
}

func (store *DatabaseStore) GetMovies() ([]*models.Movie, error) {
	var movies []*models.Movie
	err := store.db.Select(&movies, "SELECT * FROM movies")
	return movies, err
}
