package controllers

import (
    "encoding/json"
    "html/template"
    "net/http"
    "strings"
    "strconv"
    "groupie-tracker/models"
)

type SearchResult struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Type string `json:"type"`
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        Err405(w)
        return
    }

    query := strings.ToLower(r.URL.Query().Get("q"))
    if query == "" {
        Err400(w)
        return
    }

    if r.URL.Query().Get("suggestion") == "true" {
        handleSuggestions(w, query)
    } else {
        handleSearchResults(w, query)
    }
}

func handleSuggestions(w http.ResponseWriter, query string) {
    suggestions := []SearchResult{}
    seen := make(map[int]bool)
    for _, artist := range models.AllArtists {
        addSuggestions(&suggestions, artist, query, seen)
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(suggestions)
}

func addSuggestions(suggestions *[]SearchResult, artist models.Artist, query string, seen map[int]bool) {
    if strings.Contains(strings.ToLower(artist.Name), query) && !seen[artist.ID] {
        *suggestions = append(*suggestions, SearchResult{artist.ID, artist.Name, "artist/band name"})
        seen[artist.ID] = true
    }
    if strings.Contains(strings.ToLower(artist.FirstAlbum), query) && !seen[artist.ID] {
        *suggestions = append(*suggestions, SearchResult{artist.ID, artist.Name + " - " + artist.FirstAlbum, "first album date"})
        seen[artist.ID] = true
    }
    if strings.Contains(strconv.Itoa(artist.CreationDate), query) && !seen[artist.ID] {
        *suggestions = append(*suggestions, SearchResult{artist.ID, artist.Name + " - " + strconv.Itoa(artist.CreationDate), "creation date"})
        seen[artist.ID] = true
    }
    for _, member := range artist.Members {
        if strings.Contains(strings.ToLower(member), query) && !seen[artist.ID] {
            *suggestions = append(*suggestions, SearchResult{artist.ID, artist.Name + " - " + member, "member"})
            seen[artist.ID] = true
        }
    }
    for _, loc := range models.AllLocations.Index {
        if loc.ID == artist.ID {
            for _, location := range loc.Location {
                if strings.Contains(strings.ToLower(location), query) && !seen[artist.ID] {
                    *suggestions = append(*suggestions, SearchResult{artist.ID, artist.Name + " - " + location, "location"})
                    seen[artist.ID] = true
                }
            }
        }
    }
}

func handleSearchResults(w http.ResponseWriter, query string) {
    results := []models.Artist{}
    seen := make(map[int]bool)
    for _, artist := range models.AllArtists {
        if matchesQuery(artist, query, seen) {
            results = append(results, artist)
        }
    }

    t, err := template.ParseFiles("views/search.html", "views/header.html")
    if err != nil {
        Err500(w)
        return
    }

    if len(results) == 0 {
        results = []models.Artist{}
    }

    if err := t.Execute(w, results); err != nil {
        Err500(w)
    }
}

func matchesQuery(artist models.Artist, query string, seen map[int]bool) bool {
    if !seen[artist.ID] {
        if strings.Contains(strings.ToLower(artist.Name), query) ||
            strings.Contains(strings.ToLower(artist.FirstAlbum), query) ||
            strings.Contains(strconv.Itoa(artist.CreationDate), query) {
            seen[artist.ID] = true
            return true
        }
        for _, member := range artist.Members {
            if strings.Contains(strings.ToLower(member), query) {
                seen[artist.ID] = true
                return true
            }
        }
        for _, loc := range models.AllLocations.Index {
            if loc.ID == artist.ID {
                for _, location := range loc.Location {
                    if strings.Contains(strings.ToLower(location), query) {
                        seen[artist.ID] = true
                        return true
                    }
                }
            }
        }
    }
    return false
}
