package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/MdnDev/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(responseWriter http.ResponseWriter, r *http.Request){
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(responseWriter, 400, fmt.Sprintf("Error parsing jSON err: %s", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:  		uuid.New(),
		CreatedAt:	time.Now().UTC(),
		UpdatedAt: 	time.Now().UTC(),
		Name:		params.Name,
	})

	if err != nil {
		respondWithError(responseWriter, 400, fmt.Sprintf("Could not create user : %s", err))
		return
	}


	respondWithJSON(responseWriter, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(responseWriter http.ResponseWriter, r *http.Request, user database.User){
	respondWithJSON(responseWriter, 200, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetPostsForUser(responseWriter http.ResponseWriter, r *http.Request, user database.User){
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: 10,
	})

	if err != nil {
		respondWithError(responseWriter, 400, fmt.Sprintf("Could not get posts: %v", err))
	}

	respondWithJSON(responseWriter, 200, databasePostsToPosts(posts))
}