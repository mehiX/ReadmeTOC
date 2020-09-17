package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
	"github.com/mehiX/ReadmeTOC/internal/handlers"
)

func main() {
	if 2 != len(os.Args) {
		fmt.Printf("Usage: %s addr\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	addr := os.Args[1]

	handler := mux.NewRouter()

	handler.HandleFunc("/", handlers.Home).Methods(http.MethodGet)
	handler.HandleFunc("/query", handlers.HandleQueryParam).Methods(http.MethodGet)
	handler.HandleFunc("/json", handlers.HandleJSON).Methods(http.MethodGet)

	server := &http.Server{
		Addr:              addr,
		Handler:           handler,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	fmt.Println("Listening...")
	log.Fatal(server.ListenAndServe())

}
