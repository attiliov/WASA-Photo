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
			User.id,
			User.username,
			User.signup_date,
			User.last_seen,
			User.bio,
			User.profile_image_id,
			User.followers_count,
			User.following_count
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
		err = rows.Scan(&follower.UserID, &follower.Username, &follower.SignUpDate, &follower.LastSeenDate, &follower.Bio, &follower.ProfileImage, &follower.Followers, &follower.Following)
		if err != nil {
			return followers, fmt.Errorf("scanning follower: %w", err)
		}
		followers = append(followers, follower)
	}
	if err := rows.Err(); err != nil {
		return followers, fmt.Errorf("iterating over followers: %w", err)
	}
	return followers, nil
}

// GetFollowingsList returns all the users followed by the user with the given userID
func (db *appdbimpl) GetFollowingsList(userID string) ([]structs.User, error) {
	var followings []structs.User
	rows, err := db.c.Query(`
		SELECT 
			User.id,
			User.username,
			User.signup_date,
			User.last_seen,
			User.bio,
			User.profile_image_id,
			User.followers_count,
			User.following_count
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
		err = rows.Scan(&following.UserID, &following.Username, &following.SignUpDate, &following.LastSeenDate, &following.Bio, &following.ProfileImage, &following.Followers, &following.Following)
		if err != nil {
			return followings, fmt.Errorf("scanning following: %w", err)
		}
		followings = append(followings, following)
	}
	if err := rows.Err(); err != nil {
		return followings, fmt.Errorf("iterating over followings: %w", err)
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
		SET following_count = following_count + 1
		WHERE id = ?`, userID)
	if err != nil {
		return fmt.Errorf("updating following counter: %w", err)
	}

	// Update the followers counter
	_, err = db.c.Exec(`
		UPDATE User
		SET followers_count = followers_count + 1
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
		SET following_count = following_count - 1
		WHERE id = ?`, userID)
	if err != nil {
		return fmt.Errorf("updating following counter: %w", err)
	}
	// Update the followers counter
	_, err = db.c.Exec(`
		UPDATE User
		SET followers_count = followers_count - 1
		WHERE id = ?`, followingID)
	if err != nil {
		return fmt.Errorf("updating followers counter: %w", err)
	}

	return nil
}
