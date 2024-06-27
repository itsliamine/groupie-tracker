package main

import (
	"encoding/json"
	"groupie-tracker/datatypes"
	"groupie-tracker/utils"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const RED = "\033[31;1m"
const GREEN = "\033[32;1m"
const YELLOW = "\033[33;1m"
const NONE = "\033[0m"

var artistsJson []datatypes.Artist
var locationsJson []datatypes.Location
var relationJson []datatypes.Relation

func main() {
	log.SetFlags(log.Ltime)
	log.SetPrefix("groupie-tracker:")

	f, err := os.ReadFile("json/artists.json")
	if err != nil {
		log.Fatal("Error when reading json file", err)
		return
	}
	json.Unmarshal(f, &artistsJson)

	f, err = os.ReadFile("json/locations.json")
	if err != nil {
		log.Fatal("Error when reading json file", err)
		return
	}
	json.Unmarshal(f, &locationsJson)

	f, err = os.ReadFile("json/relation.json")
	if err != nil {
		log.Fatal("Error when reading json file", err)
		return
	}
	json.Unmarshal(f, &relationJson)

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

	var s datatypes.SearchRequest

	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var response datatypes.SearchReponse

	for _, v := range artistsJson {
		if strings.Contains(strings.ToLower(v.Name), strings.ToLower(s.Search)) {
			response.Artists = append(response.Artists, v)
		}
		for _, member := range v.Members {
			if strings.Contains(strings.ToLower(member), strings.ToLower(s.Search)) {
				t := "member of "
				if len(v.Members) == 1 {
					t = "real name of "
				}
				d := datatypes.MemberResponse{
					Name:   member,
					Artist: v,
					Type:   t,
				}
				response.Members = append(response.Members, d)
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func artistHandler(w http.ResponseWriter, r *http.Request) {
	requestID := r.URL.Query().Get("id")

	var a datatypes.ArtistResponse
	artistID, _ := strconv.Atoi(requestID)
	a.Artist = utils.GetArtist(artistID, artistsJson)
	a.Relations = utils.GetRelations(artistID, relationJson)

	t, _ := template.ParseFiles("templates/artist.html")
	t.Execute(w, a)
}

func badRequestHandler(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	t, _ := template.ParseFiles("templates/400.html")
	t.Execute(w, nil)
}
