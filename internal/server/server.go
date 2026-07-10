package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mzeahmed/goflix/internal/store"
)

const JWT_APP_KEY = "training.go"

type Server struct {
	router *mux.Router
	Store  store.Store
}

func NewServer() *Server {
	s := &Server{
		router: mux.NewRouter(),
	}

	s.routes()
	return s
}

func (s *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	logRequestMiddleware(s.router.ServeHTTP).ServeHTTP(writer, request)
}

// Reponse par défaut du serveur
func (s *Server) respond(w http.ResponseWriter, _ *http.Request, data interface{}, status int) {
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
func (s *Server) decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
