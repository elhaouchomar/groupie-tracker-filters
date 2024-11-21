package controllers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"groupie-tracker/models"
)

func HandleData(w http.ResponseWriter, r *http.Request) {
	api, err := models.FetchAllData() 
	if err != nil {
		log.Printf("Error fetching data: %v", err)
		Err500(w)
	}
	tmplPath := filepath.Join("views", filepath.Base(r.URL.Path)+".html")
	t, err := template.ParseFiles(tmplPath, "views/header.html")
	if err != nil {
		Err500(w)
		return
	}
	if err := t.Execute(w, api); err != nil {
		Err500(w)
	}
}
