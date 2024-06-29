package main

import (
	"encoding/json"
	"net/http"
	"log"
)

func respondWithError(responseWriter http.ResponseWriter, code int, msg string){
	if code > 499 {
		log.Println("Responding woth 500 level error: ", msg)
		
		
	}

	type errResponse struct{
		Error string `json:"error"`
	}
	respondWithJSON(responseWriter, code, errResponse{
		Error: msg,
	})
}

func respondWithJSON(responseWriter http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed to Marshal JSON response: %v", payload)
		responseWriter.WriteHeader(500)
		return
	}

	responseWriter.Header().Add("Content-type", "application/json")
	responseWriter.WriteHeader(code)
	responseWriter.Write(data)
}