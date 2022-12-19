package main

import "net/http"

func (s *Server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	}
}

func (s *Server) routes() {
	s.router.HandleFunc("/", s.handleIndex()).Methods("GET")
	s.router.HandleFunc("/api/movies/{id}", s.handleGetMovie()).Methods("GET")
	s.router.HandleFunc("/api/movies", s.handleGetMovies()).Methods("GET")
	s.router.HandleFunc("/api/movies", s.handleCreateMovie()).Methods("POST")
}
