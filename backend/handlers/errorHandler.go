package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

type ErrorPage struct {
	StatusCode string
	StatusText string
}

func ErrHandler(w http.ResponseWriter, r *http.Request, statusCode int, statusText string) {
	data := ErrorPage{
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
