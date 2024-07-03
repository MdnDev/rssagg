package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/MdnDev/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(responseWriter http.ResponseWriter, r *http.Request, user database.User){
	type parameters struct {
		Name string `json:"name"`
		URL string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(responseWriter, 400, fmt.Sprintf("Error parsing jSON err: %s", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:  		uuid.New(),
		CreatedAt:	time.Now().UTC(),
		UpdatedAt: 	time.Now().UTC(),
		Name:		params.Name,
		Url:		params.URL,
		UserID: 	user.ID,
	})

	if err != nil {
		respondWithError(responseWriter, 400, fmt.Sprintf("Could not create feed : %s", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:  		uuid.New(),
		CreatedAt:	time.Now().UTC(),
		UpdatedAt: 	time.Now().UTC(),
		UserID: 	user.ID,
		FeedID:		feed.ID,
	})

	if err != nil {
		respondWithError(responseWriter, 400, fmt.Sprintf("Could not create feed follow : %s", err))
		return
	}


	respondWithJSON(responseWriter, 201, struct{
		Feed 		Feed `json:"feed"`
		FeedFollow	FeedFollow `json:"feedFollow"`
	} {
		Feed: databaseFeedToFeed(feed),
		FeedFollow: databaseFeedFollowToFeedFollow(feedFollow),
	})
}

func (apiCfg *apiConfig) handlerGetFeeds(responseWriter http.ResponseWriter, r *http.Request){
	feeds, err := apiCfg.DB.GetFeeeds(r.Context())

	if err != nil {
		respondWithError(responseWriter, 400, fmt.Sprintf("Could not get feeds : %s", err))
		return
	}


	respondWithJSON(responseWriter, 201, databaseFeedsToFeeds(feeds))
}