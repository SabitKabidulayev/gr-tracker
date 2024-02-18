package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

type ErrorStruct struct {
	StatusCode string
	StatusText string
}

func ErrorPage(w http.ResponseWriter, r *http.Request, statusCode int, statusText string) {
	data := ErrorStruct{
		StatusCode: fmt.Sprint(statusCode),
		StatusText: statusText,
	}
	t, err := template.ParseFiles("./frontend/html/errpage.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}
	w.WriteHeader(statusCode)
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Error when executing", http.StatusInternalServerError)
		return
	}
}
