package api

import (
	"encoding/json"
	"net/http"

	"github.com/attiliov/WASA-Photo/service/structs"
	"github.com/julienschmidt/httprouter"
)

/*
	This file contains the handlers for the API endpoint that returns the user feed
	i.e. the following endpoint:
		- GET /users/:userId/feed
*/

func (rt *_router) getFeed(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	// Get the user ID from the URL
	userID := ps.ByName("userId")

	// Check authorization (a user can request only their own feed)
	beaerToken, err := getBearerToken(r)
	if err != nil || userID != beaerToken {
		// If there was an error getting the bearer token, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get the user feed
	feed, err := rt.db.GetUserFeed(userID)
	if err != nil {
		rt.baseLogger.Println("Error getting user feed:", err)
		// If there was an error getting the user feed, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create a response object
	response := structs.PostStream{Posts: feed}

	// Set the header and write the response body
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}