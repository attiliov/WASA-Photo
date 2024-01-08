package api

import (
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

/*
	This file contains the handlers for the API endpoints that are used to interact with the comments database
	i.e. the following endpoints:
		- GET /users/:userId/posts/postId/comments
		- POST /users/:userId/posts/postId/comments
		- GET /users/:userId/posts/postId/comments/:commentId
		- PUT /users/:userId/posts/postId/comments/commentId
		- DELETE /users/:userId/posts/postId/comments/:commentId
*/


func (rt *_router) getPostComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Get the post ID from the URL
	postID := ps.ByName("postId")

	// Get the comments of the specified post
	comments, err := rt.db.getPostComments(postID)
	if err != nil {
		// If there was an error getting the comments, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create a response object
	response := CommentStream{Comments: comments}

	// Set the header and write the response body
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (rt *_router) createComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	// Get the post ID from the URL
	postID := ps.ByName("postId")

	// Get the user ID from the URL
	userID := ps.ByName("userId")

	// Check authorization
	beaerToken, err := getBearerToken(r)
	if err != nil || beaerToken != userID {
		// If there was an error getting the bearer token, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get the comment from the request body
	var comment Comment
	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		// If there was an error decoding the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create the comment
	err = rt.db.createComment(userID, postID, comment)
	if err != nil {
		// If there was an error creating the comment, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the header and write the response body
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (rt *_router) getComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Get the comment ID from the URL
	commentID := ps.ByName("commentId")

	// Get the comment
	comment, err := rt.db.getComment(commentID)
	if err != nil {
		// If there was an error getting the comment, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create a response object
	response := Success{Message: "Comment retrieved successfully", Body: comment}

	// Set the header and write the response body
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (rt *_router) editComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	// Get the comment ID from the URL
	commentID := ps.ByName("commentId")

	// Get the user ID from the URL
	userID := ps.ByName("userId")

	// Check authorization
	beaerToken, err := getBearerToken(r)
	if err != nil || beaerToken != userID {
		// If there was an error getting the bearer token, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get the comment from the request body
	var comment Comment
	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		// If there was an error decoding the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Edit the comment
	err = rt.db.editComment(commentID, comment)
	if err != nil {
		// If there was an error editing the comment, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the header and write the response body
	w.WriteHeader(http.StatusOK)
} 

func (rt *_router) deleteComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	// Get the comment ID from the URL
	commentID := ps.ByName("commentId")

	// Get the user ID from the URL
	userID := ps.ByName("userId")

	// Check authorization
	beaerToken, err := getBearerToken(r)
	if err != nil || beaerToken != userID {
		// If there was an error getting the bearer token, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Delete the comment
	err = rt.db.deleteComment(commentID)
	if err != nil {
		// If there was an error deleting the comment, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the header and write the response body
	w.WriteHeader(http.StatusOK)
}