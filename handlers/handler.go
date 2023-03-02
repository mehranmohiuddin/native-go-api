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

func MoviesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getMovieHandler((w))
	default:
		returnJsonResponse(w, http.StatusNotFound, "Method not found", "false")
	}
}

func getMovieHandler(w http.ResponseWriter) {
	returnJsonResponse(w, http.StatusOK, "Successfully called movie handler", "true")
}
