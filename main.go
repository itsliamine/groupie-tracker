package main

import (
	"html/template"
	"log"
	"net/http"
)

const RED = "\033[31;1m"
const GREEN = "\033[32;1m"
const YELLOW = "\033[33;1m"
const NONE = "\033[0m"

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Locations struct {
	Id    int      `json:"id"`
	Locs  []string `json:"locations"`
	Dates string   `json:"dates"`
}

type Date struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

func main() {
	log.SetFlags(log.Ltime)
	log.SetPrefix("groupie-tracker:")

	http.HandleFunc("/", homeHandler)

	log.Println(GREEN, "Server started at http://localhost:8080", NONE)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		badRequestHandler(w)
	}
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
}

func badRequestHandler(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	t, _ := template.ParseFiles("templates/400.html")
	t.Execute(w, nil)
}
