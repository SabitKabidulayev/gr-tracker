package handlers

import (
	"groupie-tracker/backend/data"
	"groupie-tracker/backend/utilities"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrHandler(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	if r.URL.Path != "/artist/" {
		ErrHandler(w, r, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	id := r.URL.Query().Get("id")

	if !utilities.IsValid(id) {
		ErrHandler(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	if utilities.StartsWithZero(id) {
		ErrHandler(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	idd, err := strconv.Atoi(id)

	err = data.FetchDataFromJSON(&data.Artists, "https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		ErrHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	if !utilities.IsInRange(idd) {
		ErrHandler(w, r, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	err = data.GetDataForArtist(idd)
	if err != nil {
		ErrHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	t, err := template.ParseFiles("frontend/html/artist.html")
	if err != nil {
		log.Println(err)
		ErrHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	err = t.Execute(w, data.Artists[idd-1])
	if err != nil {
		http.Error(w, "Error executin file", http.StatusInternalServerError)
		return
	}
}
