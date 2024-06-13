package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
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

type SearchRequest struct {
	Search string
}

type SearchReponse struct {
	Artist Artist
	Type   string
}

var artistsJson []Artist

func main() {
	log.SetFlags(log.Ltime)
	log.SetPrefix("groupie-tracker:")

	f, err := os.ReadFile("json/artists.json")
	if err != nil {
		log.Fatal("Error when reading json file", err)
		return
	}

	json.Unmarshal(f, &artistsJson)

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/searchbar", searchBarHandler)
	http.HandleFunc("/artist", artistHandler)

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

func searchBarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		badRequestHandler(w)
	}

	var s SearchRequest

	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var response []SearchReponse
	for _, v := range artistsJson {
		if strings.Contains(strings.ToLower(v.Name), strings.ToLower(s.Search)) {
			r := SearchReponse{
				Artist: v,
				Type:   "artist/band",
			}
			response = append(response, r)
		}
		for _, member := range v.Members {
			if strings.Contains(strings.ToLower(member), strings.ToLower(s.Search)) {
				r := SearchReponse{
					Artist: v,
					Type:   "member",
				}
				response = append(response, r)
			}
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func artistHandler(w http.ResponseWriter, r *http.Request) {
	id := 1
	var a Artist
	for _, v := range artistsJson {
		if v.Id == id {
			a = v
		}
	}
	t, _ := template.ParseFiles("templates/artist.html")
	t.Execute(w, a)
}

func badRequestHandler(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	t, _ := template.ParseFiles("templates/400.html")
	t.Execute(w, nil)
}

// func internalServerErrorHandler(w http.ResponseWriter, err error) {
// 	w.WriteHeader(http.StatusInternalServerError)
// 	t, _ := template.ParseFiles("templates/500.html")
// 	t.Execute(w, err)
// }
