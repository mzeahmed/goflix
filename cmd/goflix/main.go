package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mzeahmed/goflix/internal/server"
	"github.com/mzeahmed/goflix/internal/store"
)

func main() {
	fmt.Println("Goflix")

	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s \n", err)
		os.Exit(1)
	}
}

func run() error {
	srv := server.NewServer()
	srv.Store = &store.DBStore{}
	err := srv.Store.Open()
	if err != nil {
		return err
	}
	defer func(s store.Store) {
		_ = s.Close()
	}(srv.Store)

	http.HandleFunc("/", srv.ServeHTTP)
	log.Printf("Serving HTTP on port 9000")
	err = http.ListenAndServe(":9000", nil)
	if err != nil {
		return err
	}

	return nil
}
