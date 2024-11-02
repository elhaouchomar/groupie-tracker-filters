package controllers

import (
    "html/template"
    "net/http"
    "strconv"
    "strings"
    "groupie-tracker/models"
)

var detailsTemplate = template.Must(template.ParseFiles("views/details.html", "views/header.html"))

type ResponseDetails struct {
    Artist    models.Artist
    Dates     []string
    Locations []string
    Relations map[string][]string
}

func HandleDetails(w http.ResponseWriter, r *http.Request) {
    artistIDStr := strings.TrimPrefix(r.URL.Path, "/details/")
    artistID, err := strconv.Atoi(artistIDStr)
    if err != nil {
        Err404(w)
        return
    }
    api, err := models.FetchAllData()
    if err != nil {
        Err500(w)
        return
    }
    var artist models.Artist
    found := false
    for _, a := range api.Artists {
        if a.ID == artistID {
            artist = a
            found = true
            break
        }
    }
    if !found {
        Err404(w)
        return
    }
    index := -1
    for i, a := range api.Artists {
        if a.ID == artist.ID {
            index = i
            break
        }
    }
    data := &ResponseDetails{
        Artist:    artist,
        Dates:     api.Dates.Index[index].Dates,
        Locations: api.Locations.Index[index].Location,
        Relations: api.Relation.Index[index].DatesLocations,
    }
    if err := detailsTemplate.ExecuteTemplate(w, "details.html", data); err != nil {
        Err500(w)
    }
}