package main

import (
	"fmt"
	"goflix/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func (s Server) handleCreateMovie() http.HandlerFunc {
	type CreateMovieDTO struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		ReleaseDate string `json:"releaseDate"`
		Duration    int64  `json:"duration"`
		TrailerURL  string `json:"trailerUrl"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var movieBody CreateMovieDTO

		err := s.decode(w, r, r.Body, &movieBody)
		if err != nil {
			customErr := NewError("couldn't parse movie body", http.StatusBadRequest)
			s.respondWithError(w, r, customErr, log.Error)
			return
		}

		// run validation
		if movieBody.Title == "" {
			customErr := NewError("movie title cannot be empty", http.StatusBadRequest)
			s.respondWithError(w, r, customErr, log.Error)
			return
		} else if movieBody.Description == "" {
			customErr := NewError("movie description cannot be empty", http.StatusBadRequest)
			s.respondWithError(w, r, customErr, log.Error)
			return
		} else if movieBody.ReleaseDate == "" {
			customErr := NewError("movie release date cannot be empty", http.StatusBadRequest)
			s.respondWithError(w, r, customErr, log.Error)
			return
		} else if movieBody.Duration == 0 {
			customErr := NewError("movie duration cannot be equal to 0", http.StatusBadRequest)
			s.respondWithError(w, r, customErr, log.Error)
			return
		} else if movieBody.TrailerURL == "" {
			customErr := NewError("movie trailer url cannot be empty", http.StatusBadRequest)
			s.respondWithError(w, r, customErr, log.Error)
			return
		}

		m := &models.Movie{
			ID:          0,
			Title:       movieBody.Title,
			Description: movieBody.Description,
			ReleaseDate: movieBody.ReleaseDate,
			Duration:    movieBody.Duration,
			TrailerURL:  movieBody.TrailerURL,
		}

		// Check if a movie with the same name already exists
		if s.Store.MovieExists(m.Title) {
			customErr := NewError("movie already exists", http.StatusConflict)
			s.respondWithError(w, r, customErr, log.Error)
			return
		}

		err = s.Store.CreateMovie(m)
		if err != nil {
			fmt.Println(err)
			customErr := NewError("couldn't create the movie", http.StatusInternalServerError)
			s.respondWithError(w, r, customErr, log.Error)
			return
		}
		s.respondWithJSON(w, r, m.Title, http.StatusAccepted)
	}
}

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

		movie, _ := s.Store.GetMovie(id)

		if movie == nil {
			customErr := NewError("movie not found", http.StatusNotFound)
			s.respondWithError(w, r, customErr, log.Warn)
			return
		}

		jsonMovie := models.MapMovieToJSON(*movie)
		s.respondWithJSON(w, r, jsonMovie, http.StatusOK)
	}
}
