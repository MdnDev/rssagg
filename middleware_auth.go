package main

import (
	"fmt"
	"net/http"

	"github.com/MdnDev/rssagg/internal/auth"
	"github.com/MdnDev/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, r *http.Request){
		apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(responseWriter, 403, fmt.Sprintf("Auth error: %v", err))
		return
	}

	user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(responseWriter, 400, fmt.Sprintf("Couldn't get user: %v", err))
		return
	}

	handler(responseWriter, r, user)
	}
}