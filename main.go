package main

import (
	"log"
	"net/http"
	"strings"

	"groupie-tracker/controllers"
	"groupie-tracker/models"
)

type Event struct {
	Path string
}

func eventProcessor(events <-chan Event) {
	for event := range events {
		log.Println("Processing event for path:", event.Path)
	}
}

func handleRequest(events chan<- Event, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		events <- Event{Path: r.URL.Path}
		handler(w, r)
	}
}

func main() {
	if _, err := models.FetchAllData(); err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}

	events := make(chan Event)
	go eventProcessor(events)

	mux := http.NewServeMux()

	// Serve static files from the 'static' and 'dynamic' folders
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	mux.Handle("/dynamic/", http.StripPrefix("/dynamic/", http.FileServer(http.Dir("./dynamic")))) // Add this line

	// Define route handlers
	mux.HandleFunc("/", handleRequest(events, controllers.HandleIndex))
	mux.HandleFunc("/details/", handleRequest(events, controllers.HandleDetails))
	mux.HandleFunc("/artists", handleRequest(events, controllers.HandleData))
	mux.HandleFunc("/locations", handleRequest(events, controllers.HandleData))
	mux.HandleFunc("/dates", handleRequest(events, controllers.HandleData))
	mux.HandleFunc("/relations", handleRequest(events, controllers.HandleData))
	mux.HandleFunc("/members", handleRequest(events, controllers.HandleData))
	mux.HandleFunc("/search", controllers.SearchHandler)
	mux.HandleFunc("/#filter2", controllers.HandleIndex)

	// Example route for dynamic actions if needed
	mux.HandleFunc("/dynamic/action", controllers.DynamicActionHandler) // Add this if you have dynamic actions

	// Wrap the mux for custom 404 handling
	wrappedMux := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		validPaths := map[string]bool{
			"/": true, "/artists": true, "/locations": true, "/dates": true,
			"/relations": true, "/members": true, "/search": true,
		}

		// 404 for invalid paths
		if !validPaths[r.URL.Path] &&
			!strings.HasPrefix(r.URL.Path, "/details/") &&
			!strings.HasPrefix(r.URL.Path, "/static/") &&
			!strings.HasPrefix(r.URL.Path, "/dynamic/") { // Update this line
			controllers.Err404(w)
			return
		}
		mux.ServeHTTP(w, r)
	})

	// Start the server
	log.Println("Server starting on port 3030...")
	log.Println("http://localhost:3030")
	if err := http.ListenAndServe(":3030", wrappedMux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
