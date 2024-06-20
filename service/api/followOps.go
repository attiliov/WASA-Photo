package api

import (
	"encoding/json"
	"net/http"

	"github.com/attiliov/WASA-Photo/service/structs"
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

	// Check authorization (bearer token not banned from user)
	beaerToken, err := getBearerToken(r)
	if err != nil {
		// If there was an error getting the bearer token, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	banned, err := rt.db.IsBanned(userID, beaerToken)
	if err != nil || banned {
		// If there was an error checking if the user is banned, return a 500 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get the followers of the specified user
	followers, err := rt.db.GetFollowersList(userID)
	if err != nil {
		// If there was an error getting the followers, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create a response object
	response := structs.UserCollection{Users: followers}

	// Set the header and write the response body
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		// If there was an error encoding the response, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (rt *_router) getFollowingsList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get the user ID from the URL
	userID := ps.ByName("userId")
	// Check authorization (bearer token not banned from user)
	beaerToken, err := getBearerToken(r)
	if err != nil {
		// If there was an error getting the bearer token, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	banned, err := rt.db.IsBanned(userID, beaerToken)
	if err != nil || banned {
		// If there was an error checking if the user is banned, return a 500 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get the followings of the specified user
	followings, err := rt.db.GetFollowingsList(userID)
	if err != nil {
		rt.baseLogger.Println("err:", err)
		// If there was an error getting the followings, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create a response object
	response := structs.UserCollection{Users: followings}

	// Set the header and write the response body
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		// If there was an error encoding the response, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Get the user ID from the URL
	userID := ps.ByName("userId")

	// Get the following ID from the URL
	followingID := ps.ByName("followingId")

	// Check authorization
	beaerToken, err := getBearerToken(r)
	if err != nil || beaerToken != userID {
		// If there was an error getting the bearer token, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Follow the specified user
	err = rt.db.FollowUser(userID, followingID)
	if err != nil {
		rt.baseLogger.Println("err:", err)
		// If there was an error following the user, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create a response object
	response := structs.Success{Message: "Successfully followed user"}

	// Set the header and write the response body
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		// If there was an error encoding the response, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Get the user ID from the URL
	userID := ps.ByName("userId")

	// Get the following ID from the URL
	followingID := ps.ByName("followingId")

	// Check authorization
	beaerToken, err := getBearerToken(r)
	if err != nil || beaerToken != userID {
		// If there was an error getting the bearer token, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Unfollow the specified user
	err = rt.db.UnfollowUser(userID, followingID)
	if err != nil {
		rt.baseLogger.Println("err:", err)
		// If there was an error unfollowing the user, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create a response object
	response := structs.Success{Message: "Successfully unfollowed user"}

	// Set the header and write the response body
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		// If there was an error encoding the response, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
