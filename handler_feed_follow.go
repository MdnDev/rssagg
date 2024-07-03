package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/MdnDev/rssagg/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(responseWriter http.ResponseWriter, r *http.Request, user database.User){
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(responseWriter, 400, fmt.Sprintf("Error parsing JSON err: %s", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:  		uuid.New(),
		CreatedAt:	time.Now().UTC(),
		UpdatedAt: 	time.Now().UTC(),
		UserID: 	user.ID,
		FeedID:		params.FeedID,
	})

	if err != nil {
		respondWithError(responseWriter, 400, fmt.Sprintf("Could not create feed follow : %s", err))
		return
	}


	respondWithJSON(responseWriter, 201, databaseFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollows(responseWriter http.ResponseWriter, r *http.Request, user database.User){
	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil{
		respondWithError(responseWriter, 400, fmt.Sprintf("Could not get feed follows %v", err))
		return
	}

	respondWithJSON(responseWriter, 201, databaseFeedFollowsToFeedFollows(feedFollows))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(responseWriter http.ResponseWriter, r *http.Request, user database.User){
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		respondWithError(responseWriter, 400, fmt.Sprintf("Could not parse feed follow id: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID: feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(responseWriter, 400, fmt.Sprintf("Could not delete feed follow: %v", err))
		return
	}

	respondWithJSON(responseWriter, 200, struct{}{})
}


