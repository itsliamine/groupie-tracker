package handlers

import (
	"errors"
	"html/template"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, e error) {
	w.WriteHeader(http.StatusBadRequest)
	t, err := template.ParseFiles("templates/400.html")
	if err != nil {
		ErrorHandler(w, errors.New("500 | Internal server error: Could not parse template"))
	}
	t.Execute(w, nil)
}
