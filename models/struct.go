package models

import (
	"encoding/json"
	"net/http"
)

var (
	AllArtists  []Artist
	AllLocations Locations
)

type API struct {
	Artists   []Artist
	Locations Locations
	Dates     Dates
	Relation  Relation
}

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Locations struct {
	Index []struct {
		ID       int      `json:"id"`
		Location []string `json:"locations"`
	} `json:"index"`
}

type Dates struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

type Relation struct {
	Index []struct {
		ID             int64               `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

func FetchData(url string, v interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(v)
}

func FetchAllData() (*API, error) {
	var api API
	urls := map[string]interface{}{
		"https://groupietrackers.herokuapp.com/api/artists":   &api.Artists,
		"https://groupietrackers.herokuapp.com/api/locations": &api.Locations,
		"https://groupietrackers.herokuapp.com/api/dates":     &api.Dates,
		"https://groupietrackers.herokuapp.com/api/relation":  &api.Relation,
	}

	for url, target := range urls {
		if err := FetchData(url, target); err != nil {
			return nil, err
		}
	}

	AllArtists = api.Artists
	AllLocations = api.Locations
	return &api, nil
}
