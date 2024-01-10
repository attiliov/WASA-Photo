package api

import (
	"encoding/json"
	"net/http"

	"github.com/attiliov/WASA-Photo/service/structs"
	"github.com/julienschmidt/httprouter"
)

/*
	This file contains the handlers for the API endpoints that are used to interact with the likes database
	i.e. the following endpoints:
		- GET /users/:userId/posts/postId/likes
		- PUT /users/:userId/posts/:postId/likes/:likeId
		- DELETE /users/:userId/posts/:postId/likes/:likeId
		- GET /users/:userId/posts/postId/comments/:commentId/likes
		- PUT /users/:userId/posts/postId/comments/:commentId/likes/likeId
		- DELETE /users/:userId/posts/postId/comments/:commentId/likes/likeId
*/

func (rt *_router) getPostLikes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Get the post ID from the URL
	postID := ps.ByName("postId")

	// Get the user ID from the URL
	userID := ps.ByName("userId")

	// Check authorization (bearer token not banned from post owner)
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

	// Get the likes of the specified post
	likes, err := rt.db.GetPostLikes(postID)
	if err != nil {
		// If there was an error getting the likes, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create a response object
	response := structs.LikeCollection{Likes: likes}

	// Set the header and write the response body
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}


func (rt *_router) likePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Get the post ID from the URL
	postID := ps.ByName("postId")

	// Get the like ID from the URL (the user id of the liker)
	likerID := ps.ByName("likeId")

	// Check authorization i.e. likerID == bearerToken
	beaerToken, err := getBearerToken(r)
	if err != nil || beaerToken != likerID {
		// If there was an error getting the bearer token, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Like the post
	err = rt.db.LikePost(postID, likerID)
	if err != nil {
		// If there was an error liking the post, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the header and write the response body
	w.WriteHeader(http.StatusOK)
	return
}

func (rt *_router) unlikePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Get the post ID from the URL
	postID := ps.ByName("postId")

	// Get the like ID from the URL (the user id of the liker)
	likerID := ps.ByName("likeId")

	// Check authorization i.e. likerID == bearerToken
	beaerToken, err := getBearerToken(r)
	if err != nil || beaerToken != likerID {
		// If there was an error getting the bearer token, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Like the post
	err = rt.db.UnlikePost(postID, likerID)
	if err != nil {
		// If there was an error liking the post, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the header and write the response body
	w.WriteHeader(http.StatusOK)
	return
}

func (rt *_router) getCommentLikes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Get the comment ID from the URL
	commentID := ps.ByName("commentId")

	// Get the user ID from the URL
	userID := ps.ByName("userId")

	// Check authorization (bearer token not banned from comment owner)
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

	// Get the likes of the specified post
	likes, err := rt.db.GetCommentLikes(commentID)
	if err != nil {
		// If there was an error getting the likes, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create a response object
	response := structs.LikeCollection{Likes: likes}

	// Set the header and write the response body
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (rt *_router) likeComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Get the comment ID from the URL
	commentID := ps.ByName("commentId")

	// Get the like ID from the URL (the user id of the liker)
	likerID := ps.ByName("likeId")

	// Check authorization i.e. likerID == bearerToken
	beaerToken, err := getBearerToken(r)
	if err != nil || beaerToken != likerID {
		// If there was an error getting the bearer token, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Like the post
	err = rt.db.LikeComment(commentID, likerID)
	if err != nil {
		// If there was an error liking the post, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the header and write the response body
	w.WriteHeader(http.StatusOK)
	return

}

func (rt *_router) unlikeComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	// Get the comment ID from the URL
	commentID := ps.ByName("commentId")

	// Get the like ID from the URL (the user id of the liker)
	likerID := ps.ByName("likeId")

	// Check authorization i.e. likerID == bearerToken
	beaerToken, err := getBearerToken(r)
	if err != nil || beaerToken != likerID {
		// If there was an error getting the bearer token, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Like the post
	err = rt.db.UnlikeComment(commentID, likerID)
	if err != nil {
		// If there was an error liking the post, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the header and write the response body
	w.WriteHeader(http.StatusOK)
	return
}

