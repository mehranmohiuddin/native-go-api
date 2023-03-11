package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
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
	case "POST":
		createMovie(w, r)
	case "DELETE":
		deleteMovie(w)
	default:
		returnJsonResponse(w, http.StatusNotFound, "Method not found", "false")
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	moviesByteArray, err := ioutil.ReadFile("./data/movies.json")
	if err != nil {
		http.Error(w, "Error reading movies file", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(moviesByteArray)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	var m models.Movie
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, "Error decoding request", http.StatusBadRequest)
		return
	}

	validate := validator.New()
	if err := validate.Struct(m); err != nil {
		http.Error(w, "Error validating request", http.StatusBadRequest)
		return
	}

	moviesByteArray, err := ioutil.ReadFile("./data/movies.json")
	if err != nil {
		http.Error(w, "Error reading movies file", http.StatusInternalServerError)
	}

	var movies []models.Movie

	json.Unmarshal(moviesByteArray, &movies)

	newId := len(movies) + 1

	newMovies := models.Movie{
		ID:       strconv.Itoa(newId),
		Name:     m.Name,
		Year:     m.Year,
		Director: m.Director,
	}

	movies = append(movies, newMovies)

	moviesJson, err := json.Marshal(movies)
	if err != nil {
		http.Error(w, "Error marshalling new movies", http.StatusInternalServerError)
	}

	_ = ioutil.WriteFile("./data/movies.json", moviesJson, 0644)

	returnJsonResponse(w, http.StatusOK, "Successfully called create movie handler", "true")
}

func deleteMovie(w http.ResponseWriter) {
	returnJsonResponse(w, http.StatusOK, "Successfully called delete movie handler", "true")
}
