package datatypes

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

type Location struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Date struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type SearchRequest struct {
	Search string
}

type SearchReponse struct {
	Artists []Artist         `json:"artists"`
	Members []MemberResponse `json:"members"`
}

type MemberResponse struct {
	Name   string `json:"name"`
	Artist Artist `json:"artist"`
	Type   string `json:"type"`
}

type ArtistResponse struct {
	Artist    Artist              `json:"artist"`
	Relations Relation            `json:"relations"`
	Locations map[string][]string `json:"locations"`
}

type Relation struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type ArtistsPage struct {
	Artists   []Artist
	Locations []string
}

type FiltersRequest struct {
	FromCreationYear   int    `json:"fromCreationYear"`
	ToCreationYear     int    `json:"toCreationYear"`
	FromFirstAlbumYear int    `json:"fromAlbumYear"`
	ToFirstAlbumYear   int    `json:"toAlbumYear"`
	Members            int    `json:"members"`
	Location           string `json:"location"`
}

type ErrorPage struct {
	ErrStatus string
	ErrMsg    string
}
