package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/attiliov/WASA-Photo/service/globaltime"
	"github.com/attiliov/WASA-Photo/service/structs"
	"github.com/gofrs/uuid"
)

/*
	This file contains the implementation of every function used to interact with the user table
	i.e. the follwoing functions
	GetUser(username string) (structs.User, error)
	CreateUser(username string) (structs.User, error)
	SearchUsername(username string) ([]structs.User, error)
	UpdateUser(userID string, user structs.User) error
	DeleteUser(userID string) error
*/

// GetUser returns the user with the given username or id
func (db *appdbimpl) GetUser(param string) (structs.User, error) {
	var user structs.User
	err := db.c.QueryRow(`
    SELECT 
        id, 
		username, 
        signup_date, 
        last_seen, 
        bio, 
        profile_image_id, 
        followers_count, 
        following_count 
    FROM 
        User 
    WHERE 
        username = ? OR id = ? `,
		param, param).Scan(&user.UserID, &user.Username, &user.SignUpDate, &user.LastSeenDate, &user.Bio, &user.ProfileImage, &user.Followers, &user.Following)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, fmt.Errorf("user not found: %w", err)
		}
		return user, fmt.Errorf("error getting user: %w", err)
	}
	return user, nil
}

func (db *appdbimpl) CreateUser(username string) (structs.User, error) {
	var user structs.User

	// Generate a new UUID v4
	userID, err := uuid.NewV4()
	if err != nil {
		return user, fmt.Errorf("error generating UUID: %w", err)
	}

	// Set signup date and last seen date to the current time
	signupDate := globaltime.Now().String()
	lastSeenDate := globaltime.Now().String()

	err = db.c.QueryRow(`
    INSERT INTO 
        User (id, username, signup_date, last_seen, followers_count, following_count) 
    VALUES 
        (?, ?, ?, ?, ?, ?) 
    RETURNING 
        id, 
        username, 
        signup_date, 
        last_seen, 
        bio, 
        profile_image_id, 
        followers_count, 
        following_count`,
		userID.String(), username, signupDate, lastSeenDate, 0, 0).Scan(&user.UserID, &user.Username, &user.SignUpDate, &user.LastSeenDate, &user.Bio, &user.ProfileImage, &user.Followers, &user.Following)
	if err != nil {
		return user, fmt.Errorf("error creating user: %w", err)
	}
	return user, nil
}

// SearchUsername returns a list of users whose username is similar to the given one
func (db *appdbimpl) SearchUsername(username string) ([]structs.User, error) {
	var users []structs.User
	rows, err := db.c.Query(`
	SELECT 
		id, 
		username, 
		signup_date, 
		last_seen, 
		bio, 
		profile_image_id, 
		followers_count, 
		following_count 
	FROM 
		User 
	WHERE 
		username LIKE ?`,
		"%"+username+"%")

	if err != nil {
		return users, fmt.Errorf("error searching username: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var user structs.User
		err = rows.Scan(&user.UserID, &user.Username, &user.SignUpDate, &user.LastSeenDate, &user.Bio, &user.ProfileImage, &user.Followers, &user.Following)
		if err != nil {
			return users, fmt.Errorf("error scanning user: %w", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return users, fmt.Errorf("error iterating over users: %w", err)
	}
	return users, nil
}

// UpdateUser updates the user with the given userID with all the new values in the user struct
func (db *appdbimpl) UpdateUser(userID string, user structs.User) error {
	_, err := db.c.Exec(`
	UPDATE 
		User 
	SET 
		username = ?, 
		signup_date = ?, 
		last_seen = ?, 
		bio = ?, 
		profile_image_id = ?, 
		followers_count = ?, 
		following_count = ? 
	WHERE 
		id = ?`,
		user.Username, user.SignUpDate, user.LastSeenDate, user.Bio, user.ProfileImage, user.Followers, user.Following, userID)
	if err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}
	return nil
}

// DeleteUser deletes the user with the given userID
func (db *appdbimpl) DeleteUser(userID string) error {
	_, err := db.c.Exec(`
	DELETE FROM 
		User 
	WHERE 
		id = ?`,
		userID)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}
	return nil
}
