package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const JWT_APP_KEY = "training.go"

type server struct {
	router *mux.Router
	store  Store
}

func newServer() *server {
	s := &server{
		router: mux.NewRouter(),
	}

	s.routes()
	return s
}

func (s *server) serveHTTP(writer http.ResponseWriter, request *http.Request) {
	logRequestMiddleware(s.router.ServeHTTP).ServeHTTP(writer, request)
}

// Reponse par défaut du serveur
func (s *server) respond(w http.ResponseWriter, _ *http.Request, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("cannot format json. err=%v \n", err)
	}
}

// Décodage des données envoyées par le client
func (s *server) decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
