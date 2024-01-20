package api

import (
	"encoding/json"
	"github.com/attiliov/WASA-Photo/service/structs"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

/*
	This file contains the handlers for the API endpoints that are used to interact with the post database.
		i.e. the following endpoints:
		- GET /users/userId/posts
		- POST /users/userId/posts
		- PUT /users/userId/posts/postId
		- DELETE /users/userId/posts/postId
		- GET /users/userId/posts/postId
*/

func (rt *_router) getUserPosts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get the user ID from the URL
	userID := ps.ByName("userId")

	// Check that that the bearer is not banned from post owner
	beaerToken, err := getBearerToken(r)
	if err != nil {
		// If there was an error getting the bearer token, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	banned, err := rt.db.IsBanned(userID, beaerToken) // isBanned(userID, beaerToken) returns true if the bearer is banned from the user with the given ID
	if err != nil || banned {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get the posts of the specified user
	posts_id, err := rt.db.GetUserPosts(userID) // getUserPosts(userID) returns the list of post IDs of the given user
	if err != nil {
		// If there was an error getting the posts, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create a response object
	response := structs.PostStream{Posts: posts_id}

	// Set the header and write the response body
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (rt *_router) createPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get the user ID from the URL
	userID := ps.ByName("userId")

	// Parse and decode the request body into a Post object
	var post structs.UserPost
	err := json.NewDecoder(r.Body).Decode(&post)
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

	// Create a new post in the database
	post_id, err := rt.db.AddPost(post) // createPost(userID, post) returns the post ID of the created post
	if err != nil {
		// If there was an error creating the post, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create a response object
	response := structs.Success{Message: "Post created successfully", Body: post_id}

	// Set the header and write the response body
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (rt *_router) getPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get the user ID and post ID from the URL
	userID := ps.ByName("userId")
	postID := ps.ByName("postId")

	// Check that that the bearer is not banned from post owner
	beaerToken, err := getBearerToken(r)
	if err != nil {
		// If there was an error getting the bearer token, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	banned, err := rt.db.IsBanned(userID, beaerToken) // isBanned(userID, beaerToken) returns true if the bearer is banned from the user with the given ID
	if err != nil || banned {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get the post with the specified ID
	post, err := rt.db.GetPost(postID) // getPost(userID, postID) returns the post with the given ID
	if err != nil {
		// If there was an error getting the post, return a 404 status
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Set the header and write the response body
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func (rt *_router) editPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get the user ID and post ID from the URL
	userID := ps.ByName("userId")
	postID := ps.ByName("postId")

	// Parse and decode the request body into a Post object
	var post structs.UserPost
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		rt.baseLogger.Println("error decoding request body", err)
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

	// Update the post with the specified ID
	err = rt.db.UpdatePost(postID, post) // updatePost(postID, post) returns an error if the post does not exist
	if err != nil {
		// If there was an error updating the post, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create a response object
	response := structs.Success{Message: "Post updated successfully"}

	// Set the header and write the response body
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (rt *_router) deletePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get the user ID and post ID from the URL
	userID := ps.ByName("userId")
	postID := ps.ByName("postId")

	// Check that the beaer in the body matches the user ID in the URL (authorized operation)
	beaerToken, err := getBearerToken(r)
	if err != nil || beaerToken != userID {
		// If there was an error getting the bearer token, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Delete the post with the specified ID
	err = rt.db.DeletePost(postID) // deletePost(postID) returns an error if the post does not exist
	if err != nil {
		// If there was an error deleting the post, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create a response object
	response := structs.Success{Message: "Post deleted successfully"}

	// Set the header and write the response body
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
