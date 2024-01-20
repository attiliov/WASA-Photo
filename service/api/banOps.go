package api

import (
	"encoding/json"
	"net/http"

	"github.com/attiliov/WASA-Photo/service/structs"
	"github.com/julienschmidt/httprouter"
)

/*
	This file contains the handlers for the API endpoints that are used to interact with the ban database
	i.e. the following endpoints:
		- GET /users/:userId/banned
		- PUT /users/:userId/banned/:bannedId
		- DELETE /users/:userId/banned/:bannedId
*/

func (rt *_router) getUserBanList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Get the user ID from the URL
	userID := ps.ByName("userId")

	// Check authorization
	beaerToken, err := getBearerToken(r)
	if err != nil || beaerToken != userID {
		// If there was an error getting the bearer token, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get the banned users of the specified user
	bannedUsers, err := rt.db.GetUserBanList(userID)
	if err != nil {
		rt.baseLogger.Println(err)
		// If there was an error getting the banned users, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create a response object
	response := structs.UserCollection{Users: bannedUsers}

	// Set the header and write the response body
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Get the user ID from the URL
	userID := ps.ByName("userId")

	// Get the banned user ID from the URL
	bannedID := ps.ByName("bannedId")

	// Check authorization
	beaerToken, err := getBearerToken(r)
	if err != nil || beaerToken != userID {
		// If there was an error getting the bearer token, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Ban the user
	err = rt.db.BanUser(userID, bannedID)
	if err != nil {
		// If there was an error banning the user, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the header and write the response body
	w.WriteHeader(http.StatusOK)
}

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Get the user ID from the URL
	userID := ps.ByName("userId")

	// Get the banned user ID from the URL
	bannedID := ps.ByName("bannedId")

	// Check authorization
	beaerToken, err := getBearerToken(r)
	if err != nil || beaerToken != userID {
		// If there was an error getting the bearer token, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Unban the user
	err = rt.db.UnbanUser(userID, bannedID)
	if err != nil {
		rt.baseLogger.Println(err)
		// If there was an error unbanning the user, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the header and write the response body
	w.WriteHeader(http.StatusOK)
}
