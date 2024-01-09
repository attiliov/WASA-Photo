package database

import (
	"database/sql"
	"errors"
	"fmt"
)

/*
	This file contains the implementation of every function used to interact with the user table
*/


func (db *appdbimpl) GetUserFromDB(username string) (User, error) {
	var user User
	err := db.c.QueryRow(`SELECT * FROM users WHERE username=?;`, username).Scan(&user.UserID, &user.Username)
	if err != nil {
		return User{}, fmt.Errorf("error getting user from database: %w", err)
	}
	return user, nil
}

func (db *appdbimpl) CreateUserInDB(username string) (User, error) {
	// Create a new user
	user := User{
		UserID:   ResourceID{Value: "1"},
		Username: Username{Value: username},
	}
	// Insert the user into the database
	_, err := db.c.Exec(`INSERT INTO users (userId, username) VALUES (?, ?);`, user.UserID.Value, user.Username.Value)
	if err != nil {
		return User{}, fmt.Errorf("error inserting user into database: %w", err)
	}
	return user, nil
}