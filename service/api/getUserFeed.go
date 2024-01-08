package api

import (
	"encoding/json"
	"net/http"
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

	// Get the user feed
	feed, err := rt.db.getUserFeed(userID)
	if err != nil {
		// If there was an error getting the user feed, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create a response object
	response := PostStream{Posts: feed}

	// Set the header and write the response body
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}