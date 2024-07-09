package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

/*
	This file contains the handlers for the API endpoints that are used to interact with the photo database
	i.e. the following endpoints:
		- POST /users/:userId/photos
		- GET /users/:userId/photos/:photoId
		- DELETE /users/:userId/photos/:photoId
*/

func (rt *_router) savePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rt.baseLogger.Println("savePhoto called")
	// Get the user ID from the URL
	userID := ps.ByName("userId")

	// Check authorization
	beaerToken, err := getBearerToken(r)
	if err != nil || beaerToken != userID {
		// If there was an error getting the bearer token, return a 401 status
		// rt.baseLogger.Println("savePhoto called: 401")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Parse the multipart form in the request
	err = r.ParseMultipartForm(10 << 20) // Max memory 10MB
	if err != nil {
		// rt.baseLogger.Println("savePhoto called: 400")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Retrieve the file from form data
	file, _, err := r.FormFile("photo") // "photo" is the key of the form data
	if err != nil {
		// rt.baseLogger.Println("savePhoto called: 400")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Upload the photo
	id, err := rt.db.SavePhoto(userID, file)
	if err != nil {
		// If there was an error uploading the photo, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		// rt.baseLogger.Println("Error saving photo: ", err)
		return
	}

	// Set the header and write the response body
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(id)
	// rt.baseLogger.Println("id: ", id)
	if err != nil {
		// If there was an error encoding the response, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		// rt.baseLogger.Println("Error encoding response: ", err)
		return
	}
}

func (rt *_router) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Get the user ID and photo ID from the URL
	userID := ps.ByName("userId")
	photoID := ps.ByName("photoId")

	// Check authorization
	// TODO: check if the user is allowed to view the photo (ban)

	// Get the photo
	photo, err := rt.db.GetPhoto(userID, photoID)
	if err != nil {
		// If there was an error getting the photo, return a 500 status
		rt.baseLogger.Println("Error getting photo: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the header and write the response body
	w.Header().Set("Content-Type", "image/jpeg")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(photo)
	if err != nil {
		// If there was an error writing the response, return a 500 status
		rt.baseLogger.Println("Error writing response: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Get the user ID and photo ID from the URL
	userID := ps.ByName("userId")
	photoID := ps.ByName("photoId")

	// Check authorization
	beaerToken, err := getBearerToken(r)
	if err != nil || beaerToken != userID {
		// If there was an error getting the bearer token, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Delete the photo
	err = rt.db.DeletePhoto(userID, photoID)
	if err != nil {
		// If there was an error deleting the photo, return a 500 status
		rt.baseLogger.Println("Error deleting photo: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the header and write the response body
	w.WriteHeader(http.StatusOK)
}
