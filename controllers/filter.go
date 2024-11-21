package controllers

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"groupie-tracker/models"
)

func Filter(r *http.Request, api *models.API) []models.Artist {
	useCreationDateFilter := r.FormValue("filter-cd") != ""
	useFirstAlbumDateFilter := r.FormValue("filter-fad") != ""
	useMemberCountFilter := r.FormValue("filter-nom") != ""
	useLocationFilter := r.FormValue("filter-loc") != ""

	startDate, _ := strconv.Atoi(r.FormValue("start-date1"))
	endDate, _ := strconv.Atoi(r.FormValue("end-date1"))
	startFirstAlbumDate, _ := strconv.Atoi(r.FormValue("start-date2"))
	endFirstAlbumDate, _ := strconv.Atoi(r.FormValue("end-date2"))

	memberCounts := r.Form["number"]
	location := r.FormValue("location")

	var filteredArtists []models.Artist
	for _, artist := range api.Artists {
		if useCreationDateFilter && (artist.CreationDate < startDate || artist.CreationDate > endDate) {
			continue
		}
		if useFirstAlbumDateFilter {
			dateParts := strings.Split(artist.FirstAlbum, "-")
			var firstAlbumYear int
			if len(dateParts) == 3 {
				firstAlbumYear, _ = strconv.Atoi(dateParts[2])
			}
			if firstAlbumYear < startFirstAlbumDate || firstAlbumYear > endFirstAlbumDate {
				continue
			}
		}

		memberCountMatches := !useMemberCountFilter
		if useMemberCountFilter {
			for _, count := range memberCounts {
				if strconv.Itoa(len(artist.Members)) == count {
					memberCountMatches = true
					break
				}
			}
		}
		if !memberCountMatches {
			continue
		}

		locationMatches := !useLocationFilter
		if useLocationFilter {
			for _, loc := range api.Locations.Index {
				if loc.ID == artist.ID {
					for _, artistLocation := range loc.Location {
						if strings.Contains(strings.ToLower(artistLocation), strings.ToLower(location)) {
							locationMatches = true
							break
						}
					}
					break
				}
			}
		}
		if !locationMatches {
			continue
		}
		filteredArtists = append(filteredArtists, artist)
	}
	return filteredArtists
}

func HandleFilter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		Err405(w)
		return
	}

	t, err := template.ParseFiles("views/filter.html", "views/header.html")
	if err != nil {
		Err500(w)
		return
	}

	api, err := models.FetchAllData()
	if err != nil {
		Err500(w)
		return
	}

	filter := Filter(r, api)

	if err := t.Execute(w, filter); err != nil {
		Err500(w)
	}
}
