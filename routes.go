package main

// Mapping des différentes routes
func (s *server) routes() {
	s.router.HandleFunc("/", s.handleIndex()).Methods("GET")
	s.router.HandleFunc("/api/movies/", s.handleMovieList()).Methods("GET")
}
