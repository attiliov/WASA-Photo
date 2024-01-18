package api

import (
	"errors"
	"net/http"
	"strings"
)


func getBearerToken(r *http.Request) (string, error) {

	// Get the Authorization header value
	authHeader := r.Header.Get("Authorization")

    // Split the Authorization header into "Bearer" and the token
    parts := strings.Split(authHeader, " ")
    if len(parts) < 2 || parts[0] != "Bearer" {
        // If the Authorization header is not well-formed, return a 401 status
        return "", errors.New("Invalid Authorization header")
    }

    // The second part of the Authorization header is the bearer token
    bearerToken := parts[1]
	return bearerToken, nil
}
