package handlers

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/datatypes"
	"groupie-tracker/utils"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func FiltersHandler(w http.ResponseWriter, r *http.Request) {
	request := datatypes.FiltersRequest{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Println(request)

	var filtered []datatypes.Artist

	for _, artist := range utils.ArtistsJson {
		match := true

		if artist.CreationDate < request.FromCreationYear {
			match = false
		}

		if artist.CreationDate > request.ToCreationYear {
			match = false
		}

		firstAlbum, _ := strconv.Atoi(strings.Split(artist.FirstAlbum, "-")[2])

		if firstAlbum < request.FromFirstAlbumYear {
			match = false
		}

		if firstAlbum > request.ToFirstAlbumYear {
			match = false
		}

		if request.Members != 0 && len(artist.Members) != request.Members {
			match = false
		}

		if match {
			filtered = append(filtered, artist)
		}
	}

	t, _ := template.ParseFiles("templates/artist_block.html")
	t.ExecuteTemplate(w, "artist_block", filtered)
}
