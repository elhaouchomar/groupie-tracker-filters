package controllers

import (
	"encoding/json"
	"net/http"
	"log"
)

type DynamicResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func DynamicActionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		Err405(w)
		return
	}

	var requestData map[string]string
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		Err400(w)
		return
	}

	action, ok := requestData["action"]
	if !ok {
		Err400(w)
		return
	}

	var response DynamicResponse
	switch action {
	case "greet":
		name := requestData["name"]
		response = DynamicResponse{
			Message: "Hello, " + name + "!",
			Status:  "success",
		}
	case "goodbye":
		name := requestData["name"]
		response = DynamicResponse{
			Message: "Goodbye, " + name + "!",
			Status:  "success",
		}
	default:
		response = DynamicResponse{
			Message: "Unknown action",
			Status:  "error",
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
		Err500(w)
	}
}
