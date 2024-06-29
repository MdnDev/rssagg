package main

import "net/http"

func handlerErr(responseWriter http.ResponseWriter, r *http.Request){
	respondWithError(responseWriter, 400, "Something went wrong")
}