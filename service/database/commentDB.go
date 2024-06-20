package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/attiliov/WASA-Photo/service/structs"
	"github.com/gofrs/uuid"
)

/* This file contains the implementation of every function used to interact with the comment table
   i.e. the follwoing functions
   	GetPostComments(postID string) ([]structs.Comment, error)
	CreateComment(postID string, comment structs.Comment) error
	GetComment(commentID string) (structs.Comment, error)
	EditComment(commentID string, comment structs.Comment) error
	DeleteComment(commentID string) error
*/

// GetPostComments returns all the comments of the post with the given postID
func (db *appdbimpl) GetPostComments(postID string) ([]structs.Comment, error) {
	var comments []structs.Comment
	rows, err := db.c.Query(`
	SELECT 
		id, 
		author_id, 
		username,
		creation_date, 
		caption, 
		like_count 
	FROM 
		Comment 
	WHERE 
		post_id = ?`,
		postID)
	if err != nil {
		return comments, fmt.Errorf("error getting comments: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var comment structs.Comment
		err := rows.Scan(&comment.CommentID, &comment.AuthorID, &comment.AuthorUsername, &comment.CreationDate, &comment.Caption, &comment.LikeCount)
		if err != nil {
			return comments, fmt.Errorf("error getting comment: %w", err)
		}
		comments = append(comments, comment)
	}
	if err := rows.Err(); err != nil {
		return comments, fmt.Errorf("error iterating over comments: %w", err)
	}
	return comments, nil
}

// CreateComment creates a new comment in the database
func (db *appdbimpl) CreateComment(postID string, comment structs.Comment) error {
	// Check if the post exists
	var postExists bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM Post WHERE id = ?)", postID).Scan(&postExists)
	if err != nil {
		return fmt.Errorf("error checking if post exists: %w", err)
	}
	if !postExists {
		return errors.New("post does not exist")
	}

	// Check if the author exists
	var authorExists bool
	err = db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM User WHERE id = ?)", comment.AuthorID).Scan(&authorExists)
	if err != nil {
		return fmt.Errorf("error checking if author exists: %w", err)
	}
	if !authorExists {
		return errors.New("author does not exist")
	}

	// Generate a new UUID v4
	id, err := uuid.NewV4()
	if err != nil {
		return fmt.Errorf("error generating UUID: %w", err)
	}
	comment.CommentID = id.String()

	_, err = db.c.Exec(`
	INSERT INTO 
		Comment(id, username, post_id, author_id, creation_date, caption, like_count) 
	VALUES 
		(?, ?, ?, ?, ?, ?, ?)`,
		comment.CommentID, comment.AuthorUsername, postID, comment.AuthorID, comment.CreationDate, comment.Caption, 0)
	if err != nil {
		return fmt.Errorf("error creating comment: %w", err)
	}

	// Update the post's comments count
	_, err = db.c.Exec(`
	UPDATE
		Post
	SET
		comment_count = comment_count + 1
	WHERE
		id = ?`,
		postID)
	if err != nil {
		return fmt.Errorf("error updating post's comment count: %w", err)
	}

	return nil
}

// GetComment returns the comment with the given commentID
func (db *appdbimpl) GetComment(commentID string) (structs.Comment, error) {
	var comment structs.Comment
	err := db.c.QueryRow(`
	SELECT 
		id, 
		author_id, 
		username,
		creation_date, 
		caption, 
		like_count 
	FROM 
		Comment 
	WHERE 
		id = ?`,
		commentID).Scan(&comment.CommentID, &comment.AuthorID, &comment.AuthorUsername, &comment.CreationDate, &comment.Caption, &comment.LikeCount)
	if err != nil {
		if err == sql.ErrNoRows {
			return comment, errors.New("comment does not exist")
		}
		return comment, fmt.Errorf("error getting comment: %w", err)
	}
	return comment, nil
}

// EditComment edits the comment with the given commentID
func (db *appdbimpl) EditComment(commentID string, comment structs.Comment) error {
	// Check if the comment exists
	var commentExists bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM Comment WHERE id = ?)", commentID).Scan(&commentExists)
	if err != nil {
		return fmt.Errorf("error checking if comment exists: %w", err)
	}
	if !commentExists {
		return errors.New("comment does not exist")
	}

	_, err = db.c.Exec(`
	UPDATE 
		Comment 
	SET 
		caption = ? 
	WHERE 
		id = ?`,
		comment.Caption, commentID)
	if err != nil {
		return fmt.Errorf("error editing comment: %w", err)
	}
	return nil
}

// DeleteComment deletes the comment with the given commentID
func (db *appdbimpl) DeleteComment(commentID string) error {

	// Check if the comment exists
	var commentExists bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM Comment WHERE id = ?)", commentID).Scan(&commentExists)
	if err != nil {
		return fmt.Errorf("error checking if comment exists: %w", err)
	}
	if !commentExists {
		return errors.New("comment does not exist")
	}

	// Get the postID of the comment
	var postID string
	err = db.c.QueryRow("SELECT post_id FROM Comment WHERE id = ?", commentID).Scan(&postID)
	if err != nil {
		return fmt.Errorf("error getting postID of comment: %w", err)
	}

	// Delete the comment
	_, err = db.c.Exec("DELETE FROM Comment WHERE id = ?", commentID)
	if err != nil {
		return fmt.Errorf("error deleting comment: %w", err)
	}

	// Update the post's comments count
	_, err = db.c.Exec(`
	UPDATE
		Post
	SET
		comment_count = comment_count - 1
	WHERE
		id = ?`,
		postID)
	if err != nil {
		return fmt.Errorf("error updating post's comment count: %w", err)
	}

	return nil
}
