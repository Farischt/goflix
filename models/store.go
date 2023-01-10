package models

import (
	"goflix/config"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store interface {
	Open() error
	Close() error
	GetMovies() ([]*Movie, error)
	GetMovie(id int) (*Movie, error)
	CreateMovie(m *Movie) error
	MovieExists(title string) bool
}

type DatabaseStore struct {
	db *sqlx.DB
}

func NewDatabaseStore() *DatabaseStore {
	return &DatabaseStore{}
}

func (store *DatabaseStore) Open() error {
	var err error
	store.db, err = sqlx.Open("postgres", config.PSQL_URL)
	if err != nil {
		return err
	}
	log.Println("Connected to database")

	// Create tables if not exists
	store.db.MustExec(Schema)
	return err
}

func (store *DatabaseStore) Close() error {
	return store.db.Close()
}
