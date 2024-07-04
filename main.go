package main

import (
	"groupie-tracker/handlers"
	"groupie-tracker/utils"
	"log"
	"net/http"
)

const RED = "\033[31;1m"
const GREEN = "\033[32;1m"
const YELLOW = "\033[33;1m"
const NONE = "\033[0m"

func main() {
	log.SetFlags(log.Ltime)
	log.SetPrefix("groupie-tracker:")

	utils.ReadJson("json/artists.json", &utils.ArtistsJson)
	utils.ReadJson("json/locations.json", &utils.LocationsJson)
	utils.ReadJson("json/relation.json", &utils.RelationJson)

	scripts := http.FileServer(http.Dir("./templates/scripts"))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", scripts))

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/searchbar", handlers.SearchBarHandler)
	http.HandleFunc("/artist", handlers.ArtistHandler)
	http.HandleFunc("/artists", handlers.ArtistsHandler)
	http.HandleFunc("/filters", handlers.FiltersHandler)

	log.Println(GREEN, "Server started at http://localhost:8080", NONE)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
