package database

import (
	"fmt"
	"github.com/attiliov/WASA-Photo/service/structs"
)

/* This file contains the implementation of every function used to interact with the like tables
   i.e. the follwoing functions
	IsBanned(userID string, bannedID string) (bool, error)
	GetUserBanList(userID string) ([]structs.User, error)
	BanUser(userID string, bannedID string) error
	UnbanUser(userID string, bannedID string) error
*/

// IsBanned returns true if the user with the given userID is banned by the user with the given bannedID
func (db *appdbimpl) IsBanned(userID string, bannedID string) (bool, error) {
	var banned bool
	err := db.c.QueryRow(`
		SELECT EXISTS (
			SELECT
				1
			FROM
				Ban
			WHERE
				user_id = ? AND banned_user_id = ?
		)`, userID, bannedID).Scan(&banned)
	if err != nil {
		return banned, fmt.Errorf("querying ban: %w", err)
	}
	return banned, nil
}

// GetUserBanList returns all the users banned by the user with the given userID
func (db *appdbimpl) GetUserBanList(userID string) ([]structs.User, error) {
	var bannedUsers []structs.User

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
			Ban JOIN User
		ON
			Ban.banned_user_id = User.id
		WHERE
			Ban.user_id = ?`, userID)
	if err != nil {
		return bannedUsers, fmt.Errorf("querying banned users: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var bannedUser structs.User
		err = rows.Scan(&bannedUser.UserID, &bannedUser.Username, &bannedUser.SignUpDate, &bannedUser.LastSeenDate, &bannedUser.Bio, &bannedUser.ProfileImage, &bannedUser.Followers, &bannedUser.Following)
		if err != nil {
			return bannedUsers, fmt.Errorf("scanning banned user: %w", err)
		}
		bannedUsers = append(bannedUsers, bannedUser)
	}
	if err = rows.Err(); err != nil {
		return bannedUsers, fmt.Errorf("iterating banned users: %w", err)
	}
	return bannedUsers, nil
}

// BanUser bans the user with the given bannedID from the user with the given userID
func (db *appdbimpl) BanUser(userID string, bannedID string) error {

	// Check if the user is already banned
	banned, err := db.IsBanned(userID, bannedID)
	if err != nil {
		return fmt.Errorf("checking if user is already banned: %w", err)
	}
	if banned {
		return nil
	}

	// Ban user
	_, err = db.c.Exec(`
		INSERT INTO
			Ban (user_id, banned_user_id)
		VALUES
			(?, ?)`, userID, bannedID)
	if err != nil {
		return fmt.Errorf("inserting ban: %w", err)
	}
	return nil
}

// UnbanUser unbans the user with the given bannedID from the user with the given userID
func (db *appdbimpl) UnbanUser(userID string, bannedID string) error {

	// Check if the user is already banned
	banned, err := db.IsBanned(userID, bannedID)
	if err != nil {
		return fmt.Errorf("checking if user is already banned: %w", err)
	}
	if !banned {
		return nil
	}

	// Unban user
	_, err = db.c.Exec(`
		DELETE FROM
			Ban
		WHERE
			user_id = ? AND banned_user_id = ?`, userID, bannedID)
	if err != nil {
		return fmt.Errorf("deleting ban: %w", err)
	}
	return nil
}
