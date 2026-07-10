package server

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/mzeahmed/goflix/internal/store"
)

type jsonMovie struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	ReleaseDate string `json:"release_date"`
	Duration    int64  `json:"duration"`
	TrailerUrl  string `json:"trailer_url"`
}

// Récuperation de la liste des films
func (s *Server) handleMovieList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movies, err := s.Store.GetMovies()
		if err != nil {
			log.Printf("Cannot load movies, err=%v \n", err)
			s.respond(w, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = make([]jsonMovie, len(movies))
		for i, m := range movies {
			resp[i] = mapMovieToJson(m)
		}

		s.respond(w, r, resp, http.StatusOK)
	}
}

// Recuperation des details d'un film
func (s *Server) handleMovieDetail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			log.Printf("Cannot parse id to int. err=%v", err)
			s.respond(w, r, nil, http.StatusBadRequest)
			return
		}

		m, err := s.Store.GetMovieById(id)
		if err != nil {
			log.Printf("Cannot load movie. err=%v", err)
			s.respond(w, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = mapMovieToJson(m)
		s.respond(w, r, resp, http.StatusOK)
	}
}

// Creation d'un film
func (s *Server) handleMovieCreate() http.HandlerFunc {
	type request struct {
		Title       string `json:"title"`
		ReleaseDate string `json:"release_date"`
		Duration    int64  `json:"duration"`
		TrailerUrl  string `json:"trailer_url"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := request{}
		err := s.decode(w, r, &req)
		if err != nil {
			log.Printf("Cannot parse movie body. err=%v", err)
			s.respond(w, r, nil, http.StatusBadRequest)
		}

		m := &store.Movie{
			ID:          0,
			Title:       req.Title,
			ReleaseDate: req.ReleaseDate,
			Duration:    req.Duration,
			TrailerUrl:  req.TrailerUrl,
		}

		err = s.Store.CreateMovie(m)
		if err != nil {
			log.Printf("Cannot create movie in DB. err=%v", err)
			s.respond(w, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = mapMovieToJson(m)
		s.respond(w, r, resp, http.StatusOK)
	}
}

func mapMovieToJson(m *store.Movie) jsonMovie {
	return jsonMovie{
		ID:          m.ID,
		Title:       m.Title,
		ReleaseDate: m.ReleaseDate,
		Duration:    m.Duration,
		TrailerUrl:  m.TrailerUrl,
	}
}
