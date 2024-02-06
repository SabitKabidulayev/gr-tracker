package internal

import (
	"log"
	"net/http"
)

func Server() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", Home)
	mux.HandleFunc("/artist/", ArtistPage)

	fileServer := http.FileServer(http.Dir("frontend/css"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting server on http://localhost:8000")

	http.ListenAndServe(":8000", mux)
}
