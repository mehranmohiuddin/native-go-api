package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

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
		getMovie(w, r)
	case "PUT":
		updateMovie(w, r)
	case "POST":
		createMovie(w)
	case "DELETE":
		deleteMovie(w)
	default:
		returnJsonResponse(w, http.StatusNotFound, "Method not found", "false")
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	moviesByteArray, err := ioutil.ReadFile("./data/movies.json")
	if err != nil {
		log.Fatal("Error reading file")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(moviesByteArray)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid request URL", http.StatusBadRequest)
		return
	}
	id := parts[2]
	movieId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	fmt.Println("your path is:", movieId)
	returnJsonResponse(w, http.StatusOK, "Successfully called update movie handler", "true")
}

func createMovie(w http.ResponseWriter) {
	returnJsonResponse(w, http.StatusOK, "Successfully called create movie handler", "true")
}

func deleteMovie(w http.ResponseWriter) {
	returnJsonResponse(w, http.StatusOK, "Successfully called delete movie handler", "true")
}
