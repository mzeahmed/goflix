package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Store interface {
	Open() error
	Close() error

	GetMovies() ([]*Movie, error)
	GetMovieById(id int64) (*Movie, error)
}

type dbStore struct {
	db *sqlx.DB
}

// Schema de la base de données
var schema = `
CREATE TABLE IF NOT EXISTS movie
(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT,
	release_date TEXT,
	duration INTEGER,
	trailer_url TEXT
)
`

func (store *dbStore) Open() error {
	db, err := sqlx.Connect("sqlite3", "goflix.db")
	if err != nil {
		return err
	}
	log.Println("Connected to DB")
	// Création des tables
	db.MustExec(schema)
	store.db = db
	return nil
}

func (store *dbStore) Close() error {
	return store.db.Close()
}

func (store *dbStore) GetMovies() ([]*Movie, error) {
	var movies []*Movie
	err := store.db.Select(&movies, "SELECT * FROM movie")
	if err != nil {
		return movies, err
	}
	return movies, nil
}

func (store *dbStore) getMovieById(id int64) (*Movie, error) {
	var movie = &Movie{}
	err := store.db.Get(movie, "SELECT * FROM movie WHERE id=$1", id)
	if err != nil {
		return movie, nil
	}

	return movie, nil
}
