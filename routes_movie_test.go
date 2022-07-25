package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testSore struct {
	movieId int64
	movies  []*Movie
}

func (t testSore) Open() error {
	return nil
}

func (t testSore) Close() error {
	return nil
}

func (t testSore) GetMovies() ([]*Movie, error) {
	return t.movies, nil
}

func (t testSore) GetMovieById(id int64) (*Movie, error) {
	for _, m := range t.movies {
		if m.ID == id {
			return m, nil
		}
	}
	return nil, nil
}

func (t testSore) CreateMovie(m *Movie) error {
	t.movieId++
	m.ID = t.movieId
	t.movies = append(t.movies, m)
	return nil
}

func TestMovieCreateUnit(t *testing.T) {
	// Création du server avec test de base de donnée
	srv := newServer()
	srv.store = &testSore{}

	// Préparation du BODY JSON
	p := struct {
		Title       string `json:"title"`
		ReleaseDate string `json:"release_date"`
		Duration    int64  `json:"duration"`
		TrailerUrl  string `json:"trailer_url"`
	}{
		Title:       "Incepion",
		ReleaseDate: "2010-07-18",
		Duration:    148,
		TrailerUrl:  "http://url",
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(p)

	assert.Nil(t, err)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/api/movies/", &buf)

	srv.handleMovieCreate()(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestMovieCreateIntegration(t *testing.T) {
	// Création du server avec test de base de donnée
	srv := newServer()
	srv.store = &testSore{}

	// Préparation du BODY JSON
	p := struct {
		Title       string `json:"title"`
		ReleaseDate string `json:"release_date"`
		Duration    int64  `json:"duration"`
		TrailerUrl  string `json:"trailer_url"`
	}{
		Title:       "Incepion",
		ReleaseDate: "2010-07-18",
		Duration:    148,
		TrailerUrl:  "http://url",
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(p)

	assert.Nil(t, err)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/api/movies/", &buf)

	//srv.handleMovieCreate()(w, r)
	srv.serveHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}
