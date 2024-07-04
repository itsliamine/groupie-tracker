package handlers

import (
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		badRequestHandler(w)
	}
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
}
