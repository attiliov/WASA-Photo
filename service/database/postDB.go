package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/attiliov/WASA-Photo/service/structs"
)

/*
	This file contains the implementation of every function used to interact with the post table
	i.e. the follwoing functions
	GetUserPosts(userID string) ([]structs.ResourceID, error)
	AddPost(userID string, post structs.UserPost) (structs.ResourceID, error)
	GetPost(postID string) (structs.UserPost, error)
	UpdatePost(postID string, post structs.UserPost) error
	DeletePost(postID string) error

	GetUserFeed(userID string) ([]structs.ResourceID, error)

*/

// GetUserPosts returns the posts of the user with the given userID
func (db *appdbimpl) GetUserPosts(userID string) ([]structs.ResourceID, error) {
	var posts []structs.ResourceID
	rows, err := db.c.Query(`
	SELECT 
		id 
	FROM 
		Post 
	WHERE 
		author_id = ?`, 
	userID)
	if err != nil {
		return posts, fmt.Errorf("error getting user posts: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var post structs.ResourceID
		err := rows.Scan(&post.Value)
		if err != nil {
			return posts, fmt.Errorf("error scanning user posts: %w", err)
		}
		posts = append(posts, post)
	}
	return posts, nil
}

// AddPost adds a new post to the database
func (db *appdbimpl) AddPost(post structs.UserPost) (structs.ResourceID, error) {
    // Generate a new UUID v4
    id, err := uuid.NewV4()
    if err != nil {
        return post.PostID, fmt.Errorf("error generating UUID: %w", err)
    }
    post.PostID.Value = id.String()

    // Insert the post in the DB
    _, err = db.c.Exec(`
    INSERT INTO 
        Post (id, author_id, author_username, creation_date, caption, image_id, like_count, comment_count) 
    VALUES 
        (?, ?, ?, ?, ?, ?, ?, ?)`, 
    post.PostID.Value, post.AuthorUsername.Value, post.CreationDate.Value, post.Caption.Value, post.Image.URI, post.LikeCount.Value, post.CommentCount.Value)
    if err != nil {
        return post.PostID, fmt.Errorf("error inserting post: %w", err)
    }
    return post.PostID, nil
}

// GetPost returns the post with the given postID
func (db *appdbimpl) GetPost(postID string) (structs.UserPost, error) {
	var post structs.UserPost
	err := db.c.QueryRow(`
	SELECT 
		id, 
		author_id, 
		author_username, 
		creation_date, 
		caption, 
		image_id, 
		like_count, 
		comment_count 
	FROM 
		Post 
	WHERE 
		id = ?`, 
	postID).Scan(&post.PostID, &post.AuthorID, &post.AuthorUsername, &post.CreationDate, &post.Caption, &post.Image, &post.LikeCount, &post.CommentCount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return post, fmt.Errorf("post not found: %w", err)
		}
		return post, fmt.Errorf("error getting post: %w", err)
	}
	return post, nil
}

// UpdatePost updates the post with the given postID
func (db *appdbimpl) UpdatePost(postID string, post structs.UserPost) error {
	_, err := db.c.Exec(`
	UPDATE 
		Post 
	SET 
		author_id = ?, 
		author_username = ?, 
		caption = ?, 
		image_id = ?, 
		like_count = ?, 
		comment_count = ? 
	WHERE 
		id = ?`, 
	post.AuthorID.Value, post.AuthorUsername.Value, post.Caption.Value, post.Image.URI, post.LikeCount.Value, post.CommentCount.Value, postID)
	if err != nil {
		return fmt.Errorf("error updating post: %w", err)
	}
	return nil
}

// DeletePost deletes the post with the given postID
func (db *appdbimpl) DeletePost(postID string) error {
	_, err := db.c.Exec(`
	DELETE FROM 
		Post 
	WHERE 
		id = ?`, 
	postID)
	if err != nil {
		return fmt.Errorf("error deleting post: %w", err)
	}
	return nil
}


// GetUserFeed returns the posts of the users followed by the user with the given userID
func (db *appdbimpl) GetUserFeed(userID string) ([]structs.ResourceID, error) {
	var posts []structs.ResourceID
	rows, err := db.c.Query(`
	SELECT 
		Post.id 
	FROM 
		Post 
	INNER JOIN 
		Follow ON Post.author_id = Follow.following_id 
	WHERE 
		Follow.user_id = ?`, 
	userID)
	if err != nil {
		return posts, fmt.Errorf("error getting user feed: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var post structs.ResourceID
		err := rows.Scan(&post.Value)
		if err != nil {
			return posts, fmt.Errorf("error scanning user feed: %w", err)
		}
		posts = append(posts, post)
	}
	return posts, nil
}