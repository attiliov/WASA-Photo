package database

import (
	"errors"
	"fmt"
	"github.com/attiliov/WASA-Photo/service/structs"
)

/* This file contains the implementation of every function used to interact with the like tables
   i.e. the follwoing functions
   	GetPostLikes(postID string) ([]structs.Like, error)
	LikePost(postID string, likerID string) error
	UnlikePost(postID string, likerID string) error

	GetCommentLikes(commentID string) ([]structs.Like, error)
	LikeComment(commentID string, likerID string) error
	UnlikeComment(commentID string, likerID string) error
*/

// GetPostLikes returns all the likes of the post with the given postID
func (db *appdbimpl) GetPostLikes(postID string) ([]structs.Like, error) {
	var likes []structs.Like
	rows, err := db.c.Query(`
	SELECT 
		*
	FROM 
		PostLike 
	WHERE 
		post_id = ?`, 
	postID)
	if err != nil {
		return likes, fmt.Errorf("error getting likes: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var like structs.Like
		err := rows.Scan(&like.UserID, &like.Resource, &like.Username)
		if err != nil {
			return likes, fmt.Errorf("error getting like: %w", err)
		}
		likes = append(likes, like)
	}
	return likes, nil
}

// LikePost creates a new like in the database
func (db *appdbimpl) LikePost(postID string, likerID string) error {
	// Check if the post exists
	var postExists bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM Post WHERE id = ?)", postID).Scan(&postExists)
	if err != nil {
		return fmt.Errorf("error checking if post exists: %w", err)
	}
	if !postExists {
		return errors.New("post does not exist")
	}

	// Check if the user exists
	var userExists bool
	err = db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM User WHERE id = ?)", likerID).Scan(&userExists)
	if err != nil {
		return fmt.Errorf("error checking if user exists: %w", err)
	}
	if !userExists {
		return errors.New("user does not exist")
	}

	// Check if the like already exists
	var likeExists bool
	err = db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM PostLike WHERE post_id = ? AND user_id = ?)", postID, likerID).Scan(&likeExists)
	if err != nil {
		return fmt.Errorf("error checking if like exists: %w", err)
	}
	if likeExists {
		return nil
	}

	// Get user username
	var username string
	err = db.c.QueryRow("SELECT username FROM User WHERE id = ?", likerID).Scan(&username)
	if err != nil {
		return fmt.Errorf("error getting username: %w", err)
	}

	// Insert the like
	_, err = db.c.Exec("INSERT INTO PostLike (post_id, user_id, username) VALUES (?, ?, ?)", postID, likerID, username)
	if err != nil {
		return fmt.Errorf("error inserting like: %w", err)
	}

	// Update the post's like count
	_, err = db.c.Exec("UPDATE Post SET likes_count = likes_count + 1 WHERE id = ?", postID)
	if err != nil {
		return fmt.Errorf("error updating post like count: %w", err)
	}
	return nil
}

// UnlikePost deletes the like with the given postID and likerID
func (db *appdbimpl) UnlikePost(postID string, likerID string) error {
	// Check if the post exists
	var postExists bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM Post WHERE id = ?)", postID).Scan(&postExists)
	if err != nil {
		return fmt.Errorf("error checking if post exists: %w", err)
	}
	if !postExists {
		return errors.New("post does not exist")
	}

	// Check if the like exists
	var likeExists bool
	err = db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM PostLike WHERE post_id = ? AND user_id = ?)", postID, likerID).Scan(&likeExists)
	if err != nil {
		return fmt.Errorf("error checking if like exists: %w", err)
	}
	if !likeExists {
		return nil
	}

	// Delete the like
	_, err = db.c.Exec("DELETE FROM PostLike WHERE post_id = ? AND user_id = ?", postID, likerID)
	if err != nil {
		return fmt.Errorf("error deleting like: %w", err)
	}

	// Update the post's like count
	_, err = db.c.Exec("UPDATE Post SET likes_count = likes_count - 1 WHERE id = ?", postID)
	if err != nil {
		return fmt.Errorf("error updating post like count: %w", err)
	}
	return nil
}

// GetCommentLikes returns all the likes of the comment with the given commentID
func (db *appdbimpl) GetCommentLikes(commentID string) ([]structs.Like, error) {
	var likes []structs.Like
	rows, err := db.c.Query(`
	SELECT 
		*
	FROM 
		CommentLike 
	WHERE 
		comment_id = ?`, 
	commentID)
	if err != nil {
		return likes, fmt.Errorf("error getting likes: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var like structs.Like
		err := rows.Scan(&like.UserID, &like.Resource, &like.Username)
		if err != nil {
			return likes, fmt.Errorf("error getting like: %w", err)
		}
		likes = append(likes, like)
	}
	return likes, nil
}

// LikeComment creates a new like in the database
func (db *appdbimpl) LikeComment(commentID string, likerID string) error {
	// Check if the comment exists
	var commentExists bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM Comment WHERE id = ?)", commentID).Scan(&commentExists)
	if err != nil {
		return fmt.Errorf("error checking if comment exists: %w", err)
	}
	if !commentExists {
		return errors.New("comment does not exist")
	}

	// Check if the user exists
	var userExists bool
	err = db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM User WHERE id = ?)", likerID).Scan(&userExists)
	if err != nil {
		return fmt.Errorf("error checking if user exists: %w", err)
	}
	if !userExists {
		return errors.New("user does not exist")
	}

	// Check if the like already exists
	var likeExists bool
	err = db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM CommentLike WHERE comment_id = ? AND user_id = ?)", commentID, likerID).Scan(&likeExists)
	if err != nil {
		return fmt.Errorf("error checking if like exists: %w", err)
	}
	if likeExists {
		return nil
	}

	// Get user username
	var username string
	err = db.c.QueryRow("SELECT username FROM User WHERE id = ?", likerID).Scan(&username)
	if err != nil {
		return fmt.Errorf("error getting username: %w", err)
	}

	// Insert the like
	_, err = db.c.Exec("INSERT INTO CommentLike (comment_id, user_id, username) VALUES (?, ?, ?)", commentID, likerID, username)
	if err != nil {
		return fmt.Errorf("error inserting like: %w", err)
	}

	// Update the comment's like count
	_, err = db.c.Exec("UPDATE Comment SET likes_count = likes_count + 1 WHERE id = ?", commentID)
	if err != nil {
		return fmt.Errorf("error updating comment like count: %w", err)
	}
	return nil
}

// UnlikeComment deletes the like with the given commentID and likerID
func (db *appdbimpl) UnlikeComment(commentID string, likerID string) error {
	// Check if the comment exists
	var commentExists bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM Comment WHERE id = ?)", commentID).Scan(&commentExists)
	if err != nil {
		return fmt.Errorf("error checking if comment exists: %w", err)
	}
	if !commentExists {
		return errors.New("comment does not exist")
	}

	// Check if the like exists
	var likeExists bool
	err = db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM CommentLike WHERE comment_id = ? AND user_id = ?)", commentID, likerID).Scan(&likeExists)
	if err != nil {
		return fmt.Errorf("error checking if like exists: %w", err)
	}
	if !likeExists {
		return nil
	}

	// Delete the like
	_, err = db.c.Exec("DELETE FROM CommentLike WHERE comment_id = ? AND user_id = ?", commentID, likerID)
	if err != nil {
		return fmt.Errorf("error deleting like: %w", err)
	}

	// Update the comment's like count
	_, err = db.c.Exec("UPDATE Comment SET likes_count = likes_count - 1 WHERE id = ?", commentID)
	if err != nil {
		return fmt.Errorf("error updating comment like count: %w", err)
	}
	return nil
}