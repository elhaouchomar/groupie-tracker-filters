package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "strings"
    "groupie-tracker/models"
)

type SearchResult struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Type  string `json:"type"`
    Image string `json:"image,omitempty"`
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
    query := strings.ToLower(r.URL.Query().Get("q"))
    suggestions := make([]SearchResult, 0)
    artistMap := make(map[int]models.Artist)
    for _, artist := range models.AllArtists {
        artistMap[artist.ID] = artist
    }
    for _, artist := range models.AllArtists {
        if strings.Contains(strings.ToLower(artist.Name), query) {
            suggestions = append(suggestions, SearchResult{ID: artist.ID, Name: artist.Name, Type: "artist/band", Image: artist.Image})
        }
        for _, member := range artist.Members {
            if strings.Contains(strings.ToLower(member), query) {
                suggestions = append(suggestions, SearchResult{ID: artist.ID, Name: artist.Name + " - " + member, Type: "artist - member", Image: artist.Image})
            }
        }
        if strings.Contains(strings.ToLower(artist.FirstAlbum), query) {
            suggestions = append(suggestions, SearchResult{ID: artist.ID, Name: artist.Name + " - " + artist.FirstAlbum, Type: "artist - first album date", Image: artist.Image})
        }
        if strings.Contains(strconv.Itoa(artist.CreationDate), query) {
            suggestions = append(suggestions, SearchResult{ID: artist.ID, Name: artist.Name + " - " + strconv.Itoa(artist.CreationDate), Type: "artist - creation date", Image: artist.Image})
        }
        for _, locationData := range models.AllLocations.Index {
            if locationData.ID == artist.ID {
                for _, location := range locationData.Location {
                    if strings.Contains(strings.ToLower(location), query) {
                        suggestions = append(suggestions, SearchResult{ID: artist.ID, Name: artist.Name + " - " + location, Type: "location", Image: artist.Image})
                    }
                }
            }
        }
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(suggestions)
}