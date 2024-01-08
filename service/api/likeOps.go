package api

import (
	"encoding/json"
	"net/http"
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
	// Get the user ID from the URL
	userID := ps.ByName("userId")

	// Get the post ID from the URL
	postID := ps.ByName("postId")

	// Get the likes of the specified post
	likes, err := rt.db.getPostLikes(postID)
	if err != nil {
		// If there was an error getting the likes, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create a response object
	response := LikeCollection{Likes: likes}

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
	err = rt.db.likePost(postID, likerID)
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
	err = rt.db.unlikePost(postID, likerID)
	if err != nil {
		// If there was an error liking the post, return a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the header and write the response body
	w.WriteHeader(http.StatusOK)
	return
}
