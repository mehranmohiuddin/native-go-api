package main

import (
	"log"
	"net/http"

	"github.com/mehranmohiuddin/native-go-api/handlers"
)

func main() {
	http.HandleFunc("/", handlers.DefaultHandler)
	http.HandleFunc("/movies", handlers.GetMovieHandler)
	log.Default().Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
