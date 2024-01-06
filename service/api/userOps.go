package api

import (
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

/*
	This file contains the handlers for the API endpoints that are used to interact with the user database.
   	i.e. the following endpoints:
	   - GET /users
	   - GET /users/:userId
	   - PUT /users/:userId
	   - DELETE /users/:userId
*/


func (rt *_router) searchUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Parse and decode the request body into a string
	var username string
	err := json.NewDecoder(r.Body).Decode(&username)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get users with similar usernames
	users, err := rt.db.searchUsername(username)
	if err != nil {
		// If there was an error getting the users, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create a response object
	response := UserCollection{Users: users}

	// Set the header and write the response body
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}


func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get the user ID from the URL
	userID := ps.ByName("userId")

	// Get the user with the specified ID
	user, err := rt.db.getUser(userID)
	if err != nil {
		// User not found, return a 404 status
		w.WriteHeader(http.StatusNotFound)
		return		
	}

	// Set the header and write the response body
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}


func (rt *_router) updateUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get the user ID from the URL
	userID := ps.ByName("userId")

	// Parse and decode the request body into a User object
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check that the beaer in the body matches the user ID in the URL (authorized operation)
	beaerToken, err := getBearerToken(r)
	if err != nil || beaerToken != userID {
		// If there was an error getting the bearer token, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}



	// Update the user with the specified ID
	err = rt.db.updateUser(userID, user)
	if err != nil {
		// If there was an error updating the user, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the header and write the response body
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}


func (rt *_router) deleteUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get the user ID from the URL
	userID := ps.ByName("userId")

	// Check that the beaer in the body matches the user ID in the URL (authorized operation)
	beaerToken, err := getBearerToken(r)
	if err != nil || beaerToken != userID {
		// If there was an error getting the bearer token, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Delete the user with the specified ID
	err = rt.db.deleteUser(userID)
	if err != nil {
		// If there was an error deleting the user, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 200 status
	w.WriteHeader(http.StatusOK)
}