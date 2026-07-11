package store

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

type Store interface {
	Open() error
	Close() error

	GetMovies() ([]*Movie, error)
	GetMovieById(id int64) (*Movie, error)
	CreateMovie(m *Movie) error
}

type DBStore struct {
	db *sqlx.DB
}

// Schema de la base de données
var schema = `
CREATE TABLE IF NOT EXISTS movie
(
	id INT AUTO_INCREMENT PRIMARY KEY,
	title TEXT,
	release_date TEXT,
	duration INTEGER,
	trailer_url TEXT
)
`

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func dsn() string {
	host := getenv("DB_HOST", "127.0.0.1")
	port := getenv("DB_PORT", "3306")
	user := getenv("DB_USER", "goflix")
	password := getenv("DB_PASSWORD", "goflix")
	name := getenv("DB_NAME", "goflix")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, name)
}

func (store *DBStore) Open() error {
	db, err := sqlx.Connect("mysql", dsn())
	if err != nil {
		return err
	}
	log.Println("Connected to DB")
	// Création des tables
	db.MustExec(schema)
	store.db = db
	return nil
}

func (store *DBStore) Close() error {
	return store.db.Close()
}

func (store *DBStore) GetMovies() ([]*Movie, error) {
	var movies []*Movie
	err := store.db.Select(&movies, "SELECT * FROM movie")
	if err != nil {
		return movies, err
	}
	return movies, nil
}

func (store *DBStore) GetMovieById(id int64) (*Movie, error) {
	var movie = &Movie{}
	err := store.db.Get(movie, "SELECT * FROM movie WHERE id=?", id)
	if err != nil {
		return movie, nil
	}

	return movie, nil
}

func (store *DBStore) CreateMovie(m *Movie) error {
	res, err := store.db.Exec("INSERT INTO movie (title, release_date, duration, trailer_url) VALUES (?, ?, ?, ?)",
		m.Title, m.ReleaseDate, m.Duration, m.TrailerUrl)
	if err != nil {
		return err
	}

	m.ID, err = res.LastInsertId()
	return err
}
