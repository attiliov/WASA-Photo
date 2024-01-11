package database

import (
	"fmt"
	"github.com/attiliov/WASA-Photo/service/structs"
)

/* This file contains the implementation of every function used to interact with the like tables
   i.e. the follwoing functions
   	GetFollowersList(userID string) ([]structs.User, error)
	GetFollowingsList(userID string) ([]structs.User, error)
	FollowUser(userID string, followingID string) error
	UnfollowUser(userID string, followingID string) error
*/

// GetFollowersList returns all the followers of the user with the given userID
func (db *appdbimpl) GetFollowersList(userID string) ([]structs.User, error) {
	var followers []structs.User
	rows, err := db.c.Query(`
		SELECT 
			User.*
		FROM 
			Follow JOIN User
		ON
			Follow.follower = User.id
		WHERE
			Follow.following = ?`, userID)
	if err != nil {
		return followers, fmt.Errorf("querying followers: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var follower structs.User
		err = rows.Scan(&follower.UserID.Value, &follower.Username.Value, &follower.SignUpDate.Value, &follower.LastSeenDate.Value, &follower.Bio.Value, &follower.ProfileImage.URI, &follower.Followers.Value, &follower.Following.Value)
		if err != nil {
			return followers, fmt.Errorf("scanning follower: %w", err)
		}
		followers = append(followers, follower)
	}
	return followers, nil
}

// GetFollowingsList returns all the users followed by the user with the given userID
func (db *appdbimpl) GetFollowingsList(userID string) ([]structs.User, error) {
	var followings []structs.User
	rows, err := db.c.Query(`
		SELECT 
			User.*
		FROM 
			Follow JOIN User
		ON
			Follow.following = User.id
		WHERE
			Follow.follower = ?`, userID)
	if err != nil {
		return followings, fmt.Errorf("querying followings: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var following structs.User
		err = rows.Scan(&following.UserID.Value, &following.Username.Value, &following.SignUpDate.Value, &following.LastSeenDate.Value, &following.Bio.Value, &following.ProfileImage.URI, &following.Followers.Value, &following.Following.Value)
		if err != nil {
			return followings, fmt.Errorf("scanning following: %w", err)
		}
		followings = append(followings, following)
	}
	return followings, nil
}

// FollowUser adds a new entry in the Follow table
func (db *appdbimpl) FollowUser(userID string, followingID string) error {

	// Check if the the follow already exists
	var followExists bool
	err := db.c.QueryRow(`
		SELECT EXISTS(
			SELECT 1
			FROM Follow
			WHERE follower = ? AND following = ?
		)`, userID, followingID).Scan(&followExists)
	if err != nil {
		return fmt.Errorf("checking if follow exists: %w", err)
	}
	if followExists {
		return nil
	}

	// Insert the follow
	_, err = db.c.Exec(`
		INSERT INTO Follow (follower, following)
		VALUES (?, ?)`, userID, followingID)
	if err != nil {
		return fmt.Errorf("inserting follow: %w", err)
	}

	// Update the following counter
	_, err = db.c.Exec(`
		UPDATE User
		SET following = following + 1
		WHERE id = ?`, userID)
	if err != nil {
		return fmt.Errorf("updating following counter: %w", err)
	}

	// Update the followers counter
	_, err = db.c.Exec(`
		UPDATE User
		SET followers = followers + 1
		WHERE id = ?`, followingID)
	if err != nil {
		return fmt.Errorf("updating followers counter: %w", err)
	}
	
	return nil
}

// UnfollowUser removes an entry from the Follow table
func (db *appdbimpl) UnfollowUser(userID string, followingID string) error {

	// Check if the the follow exists
	var followExists bool
	err := db.c.QueryRow(`
		SELECT EXISTS(
			SELECT 1
			FROM Follow
			WHERE follower = ? AND following = ?
		)`, userID, followingID).Scan(&followExists)
	if err != nil {
		return fmt.Errorf("checking if follow exists: %w", err)
	}
	if !followExists {
		return nil
	}


	// Delete the follow
	_, err = db.c.Exec(`
		DELETE FROM Follow
		WHERE follower = ? AND following = ?`, userID, followingID)
	if err != nil {
		return fmt.Errorf("deleting follow: %w", err)
	}

	// Update the following counter
	_, err = db.c.Exec(`
		UPDATE User
		SET following = following - 1
		WHERE id = ?`, userID)
	if err != nil {
		return fmt.Errorf("updating following counter: %w", err)
	}
	// Update the followers counter
	_, err = db.c.Exec(`
		UPDATE User
		SET followers = followers - 1
		WHERE id = ?`, followingID)
	if err != nil {
		return fmt.Errorf("updating followers counter: %w", err)
	}

	return nil
}
