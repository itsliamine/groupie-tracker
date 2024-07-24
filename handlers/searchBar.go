package handlers

import (
	"encoding/json"
	"errors"
	"groupie-tracker/datatypes"
	"groupie-tracker/utils"
	"net/http"
	"strconv"
	"strings"
)

func SearchBarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ErrorHandler(w, errors.New("400 | Wrong method"))
	}

	var s datatypes.SearchRequest

	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		ErrorHandler(w, errors.New("500 | Internal server error: Could not decode json"))
	}

	var response datatypes.SearchReponse

	for _, v := range utils.ArtistsJson {
		// Artist name search
		if strings.Contains(strings.ToLower(v.Name), strings.ToLower(s.Search)) {
			response.Artists = append(response.Artists, v)
		}

		// Member search
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

		// Location search
		locations := utils.GetLocations(v.Id)
		for _, location := range locations {
			if strings.Contains(strings.ToLower(location), strings.ToLower(s.Search)) {
				response.Locations = append(response.Locations, datatypes.LocationResponse{
					Name:     location,
					Artist:   v.Name,
					ArtistId: v.Id,
				})
			}
		}

		// First album date search
		if strings.Contains(v.FirstAlbum, s.Search) {
			response.FirstAlbumDates = append(response.FirstAlbumDates, datatypes.DateResponse{
				Date:     v.FirstAlbum,
				Artist:   v.Name,
				ArtistId: v.Id,
			})
		}

		// Creation date search
		if strconv.Itoa(v.CreationDate) == s.Search {
			response.CreationDates = append(response.CreationDates, datatypes.DateResponse{
				Date:     strconv.Itoa(v.CreationDate),
				Artist:   v.Name,
				ArtistId: v.Id,
			})
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
