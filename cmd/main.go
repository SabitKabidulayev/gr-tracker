package main

import (
	"groupie-tracker/backend/handlers"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.IndexPage)
	mux.HandleFunc("/artist/", handlers.ArtistPage)

	fileServer := http.FileServer(http.Dir("frontend/css"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting server on http://localhost:8000")

	http.ListenAndServe(":8000", mux)
}
