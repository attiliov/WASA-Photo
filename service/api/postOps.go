package api

import (
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
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

	// Get the posts of the specified user
	posts_id, err := rt.db.getUserPosts(userID) // getUserPosts(userID) returns the list of post IDs of the given user
	if err != nil {
		// If there was an error getting the posts, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create a response object
	response := PostStream{Posts: posts_id}

	// Set the header and write the response body
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (rt *_router) createPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	// Get the user ID from the URL
	userID := ps.ByName("userId")

	// Parse and decode the request body into a Post object
	var post UserPost
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
	post_id, err := rt.db.addPost(userID, post) // createPost(userID, post) returns the post ID of the created post
	if err != nil {
		// If there was an error creating the post, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create a response object
	response := Success{Message: "Post created successfully", Body: post_id}

	// Set the header and write the response body
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}


func (rt *_router) getPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get the user ID and post ID from the URL
	userID := ps.ByName("userId")
	postID := ps.ByName("postId")

	// Get the post with the specified ID
	post, err := rt.db.getPost(userID, postID) // getPost(userID, postID) returns the post with the given ID
	if err != nil {
		// If there was an error getting the post, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// If the post does not exist, return a 404 status
	if post == nil {
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
	var post UserPost
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

	// Update the post with the specified ID
	err = rt.db.updatePost(postID, post) // updatePost(postID, post) returns an error if the post does not exist
	if err != nil {
		// If there was an error updating the post, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create a response object
	response := Success{Message: "Post updated successfully"}

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
	err = rt.db.deletePost(postID) // deletePost(postID) returns an error if the post does not exist
	if err != nil {
		// If there was an error deleting the post, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create a response object
	response := Success{Message: "Post deleted successfully"}

	// Set the header and write the response body
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}