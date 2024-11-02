package controllers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"groupie-tracker/models"
)

func HandleData(w http.ResponseWriter, r *http.Request) {
	dataChan := make(chan interface{})
	errChan := make(chan error)

	go func() {
		if api, err := models.FetchAllData(); err != nil {
			errChan <- err
		} else {
			dataChan <- api
		}
	}()

	select {
	case api := <-dataChan:
		tmplPath := filepath.Join("views", filepath.Base(r.URL.Path)+".html")

		t, err := template.ParseFiles(tmplPath, "views/header.html")
		if err != nil {
			Err500(w)
			return
		}

		if err := t.Execute(w, api); err != nil {
			Err500(w)
		}

	case err := <-errChan:
		log.Printf("Error fetching data: %v", err)
		Err500(w)
	}
}
