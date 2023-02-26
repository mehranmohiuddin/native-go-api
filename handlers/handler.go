package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mehranmohiuddin/native-go-api/models"
)

func returnJsonResponse(res http.ResponseWriter, httpCode int, resMessage []byte) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(httpCode)
	res.Write(resMessage)
}

func GetMovieHandler(w http.ResponseWriter, r *http.Request) {
	responseStruct := models.SuccessResponse{
		Message: "Resource retrieved successfully",
		Status:  "success",
	}

	b, err := json.Marshal(responseStruct)
	if err != nil {
		log.Fatal("Error in marshalling response struct", err)
	}

	returnJsonResponse(w, http.StatusOK, b)
}
