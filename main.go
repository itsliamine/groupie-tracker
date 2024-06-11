package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Artist struct {
	id           int    `json: data`
	image        string `json: data`
	name         string `json: data`
	members      string `json: data`
	creationDate string `json: data`
	firstAlbum   string `json: data`
	locations    string `json: data`
	concertDates string `json: data`
	relations    string `json: data`
}

func main() {
	urls := make(map[string]string)
	urls["artists"] = "https://groupietrackers.herokuapp.com/api/artists"
	urls["locations"] = "https://groupietrackers.herokuapp.com/api/locations"
	urls["dates"] = "https://groupietrackers.herokuapp.com/api/dates"
	urls["relations"] = "https://groupietrackers.herokuapp.com/api/relation"
	resp, err := http.Get(urls["artists"])
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var response []Artist
	err = json.Unmarshal(body, &response)
	fmt.Println(response)
	if err != nil {
		panic(err)
	}
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
