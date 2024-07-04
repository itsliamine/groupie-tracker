package handlers

import (
	"encoding/json"
	"groupie-tracker/datatypes"
	"groupie-tracker/utils"
	"net/http"
	"strings"
)

func SearchBarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		badRequestHandler(w)
	}

	var s datatypes.SearchRequest

	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var response datatypes.SearchReponse

	for _, v := range utils.ArtistsJson {
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
