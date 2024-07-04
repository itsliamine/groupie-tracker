package handlers

import (
	"groupie-tracker/datatypes"
	"groupie-tracker/utils"
	"html/template"
	"net/http"
	"strconv"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	requestID := r.URL.Query().Get("id")

	var a datatypes.ArtistResponse
	artistID, _ := strconv.Atoi(requestID)
	a.Artist = utils.GetArtist(artistID, utils.ArtistsJson)
	a.Relations = utils.GetRelations(artistID, utils.RelationJson)

	t, _ := template.ParseFiles("templates/artist.html")
	t.Execute(w, a)
}
