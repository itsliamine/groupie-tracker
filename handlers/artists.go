package handlers

import (
	"groupie-tracker/datatypes"
	"groupie-tracker/utils"
	"html/template"
	"net/http"
)

func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	data := datatypes.ArtistsPage{
		Artists:   utils.ArtistsJson,
		Locations: utils.GetAllLocations(utils.LocationsJson),
	}

	t, _ := template.ParseFiles("templates/artists.html", "templates/artist_block.html", "templates/slider.html")
	t.Execute(w, data)
}
