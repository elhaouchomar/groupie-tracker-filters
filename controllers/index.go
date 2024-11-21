package controllers

import (
	"groupie-tracker/models"
	"html/template"
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		Err405(w)
		return
	}

	t, err := template.ParseFiles("views/index.html", "views/header.html")
	if err != nil {
		Err500(w)
		return
	}

	api, err := models.FetchAllData()
	if err != nil {
		Err500(w)
		return
	}

	if err := t.Execute(w, api.Artists); err != nil {
		Err500(w)
	}
}
