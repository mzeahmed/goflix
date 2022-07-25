package main

import (
	"log"
	"net/http"
)

// Affichage de log lors d'une requête
func logRequestMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("[%v] %v", request.Method, request.RequestURI)
		next.ServeHTTP(writer, request)
	}
}
