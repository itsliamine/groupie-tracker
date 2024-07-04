package handlers

import (
	"html/template"
	"net/http"
)

func badRequestHandler(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	t, _ := template.ParseFiles("templates/400.html")
	t.Execute(w, nil)
}
