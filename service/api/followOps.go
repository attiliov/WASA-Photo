package api

import (
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

/*
	This file contains the handlers for the API endpoints that are used to interact with the following database
	i.e. the following endpoints:
		- GET /users/:userId/follower
		- GET /user/:userId/following
		- PUT /users/:userId/following/:followingId
		- DELETE /users/:userId/following/followingId
*/

func (rt *_router) getFollowersList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	// Get the user ID from the URL
	userID := ps.ByName("userId")

	// Get the followers of the specified user
	followers, err := rt.db.getFollowersList(userID)
	if err != nil {
		// If there was an error getting the followers, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create a response object
	response := UserCollection{Users: followers}

	// Set the header and write the response body
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (rt *_router) getFollowingsList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	// Get the user ID from the URL
	userID := ps.ByName("userId")

	// Get the followings of the specified user
	followings, err := rt.db.getFollowingsList(userID)
	if err != nil {
		// If there was an error getting the followings, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create a response object
	response := UserCollection{Users: followings}

	// Set the header and write the response body
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	// Get the user ID from the URL
	userID := ps.ByName("userId")

	// Get the following ID from the URL
	followingID := ps.ByName("followingId")

	// Check authorization
	beaerToken, err := getBearerToken(r)
	if err != nil || beaerToken != followingID {
		// If there was an error getting the bearer token, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Follow the specified user
	err = rt.db.followUser(userID, followingID)
	if err != nil {
		// If there was an error following the user, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create a response object
	response := Success{Message: "Successfully followed user"}

	// Set the header and write the response body
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	// Get the user ID from the URL
	userID := ps.ByName("userId")

	// Get the following ID from the URL
	followingID := ps.ByName("followingId")

	// Check authorization
	beaerToken, err := getBearerToken(r)
	if err != nil || beaerToken != followingID {
		// If there was an error getting the bearer token, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Unfollow the specified user
	err = rt.db.unfollowUser(userID, followingID)
	if err != nil {
		// If there was an error unfollowing the user, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create a response object
	response := Success{Message: "Successfully unfollowed user"}

	// Set the header and write the response body
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}