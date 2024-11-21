package main

import (
	"log"
	"net/http"
	"strings"

	"groupie-tracker/controllers"
	"groupie-tracker/models"
)

func main() {
	if _, err := models.FetchAllData(); err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	mux.Handle("/dynamic/", http.StripPrefix("/dynamic/", http.FileServer(http.Dir("./dynamic"))))

	mux.HandleFunc("/", controllers.HandleIndex)
	mux.HandleFunc("/details/", controllers.HandleDetails)
	mux.HandleFunc("/artists", controllers.HandleData)
	mux.HandleFunc("/locations", controllers.HandleData)
	mux.HandleFunc("/dates", controllers.HandleData)
	mux.HandleFunc("/relations", controllers.HandleData)
	mux.HandleFunc("/members", controllers.HandleData)
	mux.HandleFunc("/search", controllers.SearchHandler)
	mux.HandleFunc("/filter", controllers.HandleFilter)

	wrappedMux := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		validPaths := map[string]bool{
			"/": true, "/artists": true, "/locations": true, "/dates": true,
			"/relations": true, "/members": true, "/search": true, "/filter": true,
		}
		if !validPaths[r.URL.Path] &&
			!strings.HasPrefix(r.URL.Path, "/static/") &&
			!strings.HasPrefix(r.URL.Path, "/dynamic/") {
			controllers.Err404(w)
			return
		}
		mux.ServeHTTP(w, r)
	})

	log.Println("Server starting on port 3030...")
	log.Println("http://localhost:3030")
	if err := http.ListenAndServe(":3030", wrappedMux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
