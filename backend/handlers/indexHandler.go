package handlers

import (
	"groupie-tracker/backend/data"
	"html/template"
	"net/http"
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorPage(w, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		ErrorPage(w, http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("./frontend/templates/indexPage.html")
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError)
		return
	}

	data, err := data.GetData()
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError)
		return
	}
}
