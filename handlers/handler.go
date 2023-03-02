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
		getMovie((w))
	case "PUT":
		updateMovie(w)
	case "POST":
		createMovie(w)
	case "DELETE":
		deleteMovie(w)
	default:
		returnJsonResponse(w, http.StatusNotFound, "Method not found", "false")
	}
}

func getMovie(w http.ResponseWriter) {
	returnJsonResponse(w, http.StatusOK, "Successfully called get movie handler", "true")
}

func updateMovie(w http.ResponseWriter) {
	returnJsonResponse(w, http.StatusOK, "Successfully called update movie handler", "true")
}

func createMovie(w http.ResponseWriter) {
	returnJsonResponse(w, http.StatusOK, "Successfully called create movie handler", "true")
}

func deleteMovie(w http.ResponseWriter) {
	returnJsonResponse(w, http.StatusOK, "Successfully called delete movie handler", "true")
}
