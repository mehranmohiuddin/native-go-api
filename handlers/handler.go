package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mehranmohiuddin/native-go-api/models"
)

func returnJsonResponse(res http.ResponseWriter, httpCode int, message string, status string) {
	responseStruct := models.Response{
		Message: message,
		Success: status,
	}

	byteResponse, err := json.Marshal(responseStruct)
	if err != nil {
		log.Fatal("Error in marshalling response struct", err)
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(httpCode)
	res.Write(byteResponse)
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	returnJsonResponse(w, http.StatusOK, "Invalid URL", "false")
}

func GetMovieHandler(w http.ResponseWriter, r *http.Request) {
	movieID := r.URL.Path[len("/movies/"):]
	if movieID == "" {
		returnJsonResponse(w, http.StatusNotFound, "Error getting movie id", "false")
	} else {
		returnJsonResponse(w, http.StatusNotFound, "Resource retrieved successfully", "true")
	}
}
