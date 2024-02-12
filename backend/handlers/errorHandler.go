package handlers

import (
	"log"
	"net/http"
	"text/template"
)

type Err struct {
	StatusCode int
	StatusText string
}

func ErrorPage(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
	tmpl, err := template.ParseFiles("./frontend/templates/errorPage.html")
	if err != nil {
		http.Error(w, http.StatusText(statusCode), statusCode)
		return
	}
	er := Err{statusCode, http.StatusText(statusCode)}
	err = tmpl.Execute(w, er)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Error:", http.StatusInternalServerError)
	}
}
