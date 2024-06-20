package api

import (
	"encoding/json"
	"net/http"

	"github.com/attiliov/WASA-Photo/service/structs"
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

func (rt *_router) searchUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Get requesterId
	requesterId, err := getBearerToken(r)
	if err != nil {
		// If there was an error getting the bearer token, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Parse the request parameter into a Username object
	var username structs.Username
	username.Username = r.URL.Query().Get("username")

	// Check that the username is not empty
	if username.Username == "" {
		// If the username is empty, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get users with similar usernames
	users, err := rt.db.SearchUsername(username.Username)
	if err != nil {
		// If there was an error getting the users, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Filter users who banned the requester using rt.db.IsBanned
	filteredUsers := make([]structs.User, 0)
	for _, user := range users {
		isBanned, err := rt.db.IsBanned(user.UserID, requesterId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if !isBanned {
			filteredUsers = append(filteredUsers, user)
		}
	}
	users = filteredUsers

	// Create a response object
	response := structs.UserCollection{Users: users}

	// Set the header and write the response body
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		// If there was an error encoding the response, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get the user ID from the URL
	userID := ps.ByName("userId")

	// Check that thwe reqeuster is not banned from the userID
	beaerToken, err := getBearerToken(r)
	if err != nil {
		// If there was an error getting the bearer token, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	banned, err := rt.db.IsBanned(userID, beaerToken) // isBanned(userID, beaerToken) returns true if the bearer is banned from the user with the given ID
	if err != nil || banned {
		// If there was an error getting the banned status, return unauthorized
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get the user with the specified ID
	user, err := rt.db.GetUser(userID)
	if err != nil {
		// User not found, return a 404 status
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Set the header and write the response body
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		// If there was an error encoding the response, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (rt *_router) updateUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get the user ID from the URL
	userID := ps.ByName("userId")

	// Parse and decode the request body into a User object
	var user structs.User
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
	err = rt.db.UpdateUser(userID, user)
	if err != nil {
		// If there was an error updating the user, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the header and write the response body
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		// If there was an error encoding the response, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
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
	err = rt.db.DeleteUser(userID)
	if err != nil {
		// If there was an error deleting the user, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 200 status
	w.WriteHeader(http.StatusOK)
}
