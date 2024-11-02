package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"groupie-tracker/models"
)

func Filter(r *http.Request, api models.API) {
	fmt.Println(r.URL.Query().Get("start-date1"))
	if r.FormValue("start-date1") == "" {
		return
	}
	maxStr := r.FormValue("start-date1")
	max, er := strconv.Atoi(maxStr)
	if er != nil {
		log.Fatal(er)
	}
	
	minStr := r.FormValue("end-date1")
	min, er := strconv.Atoi(minStr)
	if er != nil {
		log.Fatal(er)
	}

	maxfirstAllbumStr := r.FormValue("start-date2")
	maxfirstAllbum, er := strconv.Atoi(maxfirstAllbumStr)
	if er != nil {
		log.Fatal(er)
	}

	minfirstAlbumStr := r.FormValue("end-date2")
	minfirstAlbum, er := strconv.Atoi(minfirstAlbumStr)
	if er != nil {
		log.Fatal(er)
	}

	var isValid bool = false
	for _, val := range api.Artists {
		if !(val.CreationDate > min && val.CreationDate < max) {
			isValid = true
		}
		slstr := strings.Split(val.FirstAlbum, "-")
		n, _ := strconv.Atoi(slstr[2])
		if !(n > minfirstAlbum && n < maxfirstAllbum) {
			isValid = true
		}
		if isValid {
			api.Artists = nil
		}
	}
}
