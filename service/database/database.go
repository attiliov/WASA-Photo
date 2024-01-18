/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/attiliov/WASA-Photo/service/structs"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {

	GetUser(username string) (structs.User, error)
	CreateUser(username string) (structs.User, error)
	SearchUsername(username string) ([]structs.User, error)
	UpdateUser(userID string, user structs.User) error
	DeleteUser(userID string) error

	GetUserPosts(userID string) ([]structs.ResourceID, error)
	AddPost(post structs.UserPost) (structs.ResourceID, error)
	GetPost(postID string) (structs.UserPost, error)
	UpdatePost(postID string, post structs.UserPost) error
	DeletePost(postID string) error

	GetPostComments(postID string) ([]structs.Comment, error)
	CreateComment(postID string, comment structs.Comment) error
	GetComment(commentID string) (structs.Comment, error)
	EditComment(commentID string, comment structs.Comment) error
	DeleteComment(commentID string) error

	GetPostLikes(postID string) ([]structs.Like, error)
	LikePost(postID string, likerID string) error
	UnlikePost(postID string, likerID string) error

	GetCommentLikes(commentID string) ([]structs.Like, error)
	LikeComment(commentID string, likerID string) error
	UnlikeComment(commentID string, likerID string) error

	GetFollowersList(userID string) ([]structs.User, error)
	GetFollowingsList(userID string) ([]structs.User, error)
	FollowUser(userID string, followingID string) error
	UnfollowUser(userID string, followingID string) error

	IsBanned(userID string, bannedID string) (bool, error)
	GetUserBanList(userID string) ([]structs.User, error)
	BanUser(userID string, bannedID string) error
	UnbanUser(userID string, bannedID string) error

	GetUserFeed(userID string) ([]structs.ResourceID, error)

	SavePhoto(userID string, photo multipart.File) error
	GetPhoto(userID string, photoID string) ([]byte, error)
	DeletePhoto(userID string, photoID string) error

	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// ---Initialize the database
    // Get the directory of the current file
    _, filename, _, ok := runtime.Caller(0)
    if !ok {
        return nil, errors.New("unable to get current directory")
    }
    dir := filepath.Dir(filename)
    // Read the init.sql file
    initSQL, err := os.ReadFile(filepath.Join(dir, "init.sql"))
    if err != nil {
        return nil, fmt.Errorf("error reading init.sql: %w", err)
    }
    // Split the SQL statements
    statements := strings.Split(string(initSQL), ";")
    // Execute each statement
    for _, statement := range statements {
        statement = strings.TrimSpace(statement) // remove leading and trailing spaces
        if statement != "" {
            _, err = db.Exec(statement)
            if err != nil {
                return nil, fmt.Errorf("error executing statement %q: %w", statement, err)
            }
        }
    }

	
	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
