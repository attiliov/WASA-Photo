package api

import (
	"encoding/json"
	"net/http"

	"github.com/attiliov/WASA-Photo/service/structs"

	"github.com/julienschmidt/httprouter"
)

// getToken is the handler for POST /session
// Parse the request body which should contain a username,
// and return the userId if the username is valid.
// If user is not present in the database, it will be created
// and the userId will be returned.

func (rt *_router) getAuthToken(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    // Parse and decode the request body into a string
    var username structs.Username
    err := json.NewDecoder(r.Body).Decode(&username)
    if err != nil {
        // If there is something wrong with the request body, return a 400 status
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    
    // Get the user from the database
    user, err := rt.db.GetUser(string(username.Username))

    // If the user doesn't exist, create a new user
    if err != nil {
        user, err = rt.db.CreateUser(string(username.Username))
        if err != nil {
            // If there was an error creating the user, return a 500 status
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        // If a new user was created, return a 201 status
        w.WriteHeader(http.StatusCreated)
    } else {
        // If the user already existed, return a 200 status
        w.WriteHeader(http.StatusOK)
    }

    // Return the user's ID in the response body
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user.UserID)
}