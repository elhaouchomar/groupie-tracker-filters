package controllers

import (
	"html/template"
	"net/http"

	"groupie-tracker/models"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		Err405(w)
		return
	}

	tmplPath := "views/index.html"
	t, err := template.ParseFiles(tmplPath, "views/header.html")
	if err != nil {
		Err500(w)
		return
	}

	api, err := models.FetchAllData()
	if err != nil {
		Err500(w)
		return
	}
	Filter(r, api)
	if err := t.Execute(w, api.Artists); err != nil {
		Err500(w)
	}
}
