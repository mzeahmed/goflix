package main

import (
	"fmt"
	"net/http"
)

func (s *server) handleIndex() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		//fmt.Fprintf(writer, "Welcome to Goflix")
		_, _ = fmt.Fprintf(writer, "Welcome to Goflix")
	}
}
