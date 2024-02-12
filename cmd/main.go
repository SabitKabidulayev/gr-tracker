package main

import (
	"groupie-tracker/backend/handlers"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("frontend/css"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", handlers.IndexPage)
	mux.HandleFunc("/artist", handlers.ArtistPage)
	log.Println("Starting server on http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
