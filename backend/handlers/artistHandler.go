package handlers

import (
	"errors"
	"fmt"
	"groupie-tracker/backend/data"
	"groupie-tracker/backend/models"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorPage(w, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/artist" {
		ErrorPage(w, http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("./frontend/templates/artistPage.html")
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError)
		return
	}

	data, err := GetID(w, r)
	if err != nil {
		if err.Error() == "404" {
			ErrorPage(w, http.StatusNotFound)
			return
		}
		ErrorPage(w, http.StatusBadRequest)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError)
		return
	}
}

func GetID(w http.ResponseWriter, r *http.Request) (models.Artist, error) {
	var artist models.Artist

	id, err := CustomAtoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		return artist, fmt.Errorf("400")
	}
	artists, err := data.GetData()
	if err != nil {
		return artist, fmt.Errorf("404")
	}
	if id > 0 && id <= len(artists) {
		artist = artists[id-1]
		return artist, err
	}

	return artist, fmt.Errorf("404")
}

func CustomAtoi(s string) (int, error) {
	if strings.TrimLeft(s, "0") != s {
		return 0, errors.New("Error")
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return n, fmt.Errorf("Error")
	}
	return n, nil
}
