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

func (rt *_router) 