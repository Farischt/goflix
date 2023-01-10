package main

import (
	"encoding/json"
	"fmt"
	"goflix/middlewares"
	"goflix/models"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type ResponseError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Date    string `json:"date"`
}

func (e ResponseError) String() string {
	return fmt.Sprintf("[HTTP] %d  %s", e.Status, e.Message)
}

func NewError(m string, s int) ResponseError {
	return ResponseError{
		Message: m,
		Status:  s,
		Date:    time.Now().Format(time.RFC3339),
	}
}

type Server struct {
	Store  models.Store
	router *mux.Router
}

func NewServer() *Server {
	s := &Server{
		router: mux.NewRouter(),
	}
	s.router.Use(middlewares.LoggingMiddleware)
	return s
}

func (s *Server) serveHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) decode(w http.ResponseWriter, r *http.Request, dataToDecode io.Reader, v interface{}) error {
	return json.NewDecoder(dataToDecode).Decode(v)
}

func (s *Server) respondWithJSON(w http.ResponseWriter, _ *http.Request, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Error while JSON encoding data: %s", err.Error())
	}
}

func (s *Server) respondWithError(w http.ResponseWriter, r *http.Request, err ResponseError, log func(args ...interface{})) {
	log(err)
	s.respondWithJSON(w, r, err, err.Status)
}
