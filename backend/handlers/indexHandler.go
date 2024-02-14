package handlers

import (
	"groupie-tracker/backend/data"
	"log"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errHandler(w, r, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	if r.Method != http.MethodGet {
		errHandler(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	t, err := template.ParseFiles("frontend/html/index.html")
	if err != nil {
		log.Println(err)
		errHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	err = data.FetchDataFromJSON(&data.Artists, "https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		errHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	err = t.Execute(w, data.Artists)
	if err != nil {
		http.Error(w, "Error when executing", http.StatusInternalServerError)
		return
	}
}
