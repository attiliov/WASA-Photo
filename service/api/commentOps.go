package api

import (
	"encoding/json"
	"net/http"
	"github.com/attiliov/WASA-Photo/service/structs"
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

	// Get the user ID from the URL
	userID := ps.ByName("userId")

	// Check that the user requesting the comments is not banned from the post owner
	requesterId, err := getBearerToken(r)
	if err != nil {
		// If there was an error getting the bearer token, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	banned, err := rt.db.IsBanned(userID, requesterId)
	if err != nil || banned {
		// If there was an error checking if the user is banned, return a 500 status
		rt.baseLogger.Println("err: ", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get the comments of the specified post
	comments, err := rt.db.GetPostComments(postID)
	if err != nil {
		// If there was an error getting the comments, return a 500 status
		rt.baseLogger.Println("err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	

	// Create a response object
	response := structs.CommentStream{Comments: comments}

	// Set the header and write the response body
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (rt *_router) createComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	// Get the post ID from the URL
	postID := ps.ByName("postId")

	// Get the user ID from the URL
	//userID := ps.ByName("userId")

	// Get the comment from the request body
	var comment structs.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		// If there was an error decoding the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check authorization
	beaerToken, err := getBearerToken(r)
	if err != nil || beaerToken != comment.AuthorID {
		// If there was an error getting the bearer token, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Check that authorId is the same as bearer token
	if string(comment.AuthorID) != beaerToken {
		// If the authorId is not the same as the bearer token, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Create the comment
	err = rt.db.CreateComment(postID ,comment)
	if err != nil {
		rt.baseLogger.Println("err: ", err)
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

	// Get the user ID from the URL
	userID := ps.ByName("userId")

	// Check that the user requesting the comments is not banned from the post owner
	requesterId, err := getBearerToken(r)
	if err != nil {
		// If there was an error getting the bearer token, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	banned, err := rt.db.IsBanned(userID, requesterId)
	if err != nil || banned {
		// If there was an error checking if the user is banned, return a 500 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	
	// Get the comment
	comment, err := rt.db.GetComment(commentID)
	if err != nil {
		// If there was an error getting the comment, return a 500 status
		//rt.baseLogger.Println("err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the header and write the response body
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comment)
}

func (rt *_router) editComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	// Get the comment ID from the URL
	commentID := ps.ByName("commentId")

	// Get the user ID from the URL
	//userID := ps.ByName("userId")

	// Get the comment from the request body
	var comment structs.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		// If there was an error decoding the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check authorization
	beaerToken, err := getBearerToken(r)
	if err != nil || beaerToken != comment.AuthorID {
		// If there was an error getting the bearer token, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Edit the comment
	err = rt.db.EditComment(commentID, comment)
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
	//userID := ps.ByName("userId")

	// Get comment author id
	comment, err := rt.db.GetComment(commentID)
	if err != nil {
		// If there was an error getting the comment, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Check authorization
	beaerToken, err := getBearerToken(r)
	if err != nil || beaerToken != comment.AuthorID {
		// If there was an error getting the bearer token, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Delete the comment
	err = rt.db.DeleteComment(commentID)
	if err != nil {
		// If there was an error deleting the comment, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the header and write the response body
	w.WriteHeader(http.StatusOK)
}