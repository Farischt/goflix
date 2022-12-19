package main

import (
	"fmt"
	"goflix/config"
	"goflix/models"
	"log"
	"net/http"
)

func run() error {
	// Init server and db
	s := NewServer()
	s.Store = models.NewDatabaseStore()
	err := s.Store.Open()
	if err != nil {
		return err
	}
	defer s.Store.Close()

	// Start server
	s.routes()
	http.HandleFunc("/", s.serveHTTP)
	log.Printf("Server is running on port %d", config.PORT)
	err = http.ListenAndServe(fmt.Sprintf(":%d", config.PORT), nil)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
