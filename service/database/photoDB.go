package database

import (
	"fmt"
	"github.com/gofrs/uuid"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

/* This file contains the implementation of every function used to interact with the photo table
	and saving photos
   i.e. the follwoing functions
   	SavePhoto(userID string, photo multipart.File) error
	GetPhoto(userID string, photoID string) ([]byte, error)
	DeletePhoto(userID string, photoID string) error
*/

// SavePhoto saves a photo in the database
func (db *appdbimpl) SavePhoto(userID string, photo multipart.File) (string, error) {
	// Generate a new UUID v4
	id, err := uuid.NewV4()
	if err != nil {
		return "", fmt.Errorf("error generating UUID: %w", err)
	}
	newId := id.String()

	// Save the photo in the filesystem on the ./photos directory with name <photoID>.jpg
	filePath := filepath.Join("/tmp/", newId+".jpg")
	file, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	_, err = io.Copy(file, photo)
	if err != nil {
		return "", fmt.Errorf("error copying photo: %w", err)
	}

	return newId, nil
}

// GetPhoto returns the photo with the given photoID
func (db *appdbimpl) GetPhoto(userID string, photoID string) ([]byte, error) {
	// Get the photo from the filesystem
	filePath := filepath.Join("./photos", photoID+".jpg")
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Read the photo
	photo, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return photo, nil
}

// DeletePhoto deletes the photo with the given photoID
func (db *appdbimpl) DeletePhoto(userID string, photoID string) error {
	// Delete the photo from the filesystem
	filePath := filepath.Join("./photos", photoID+".jpg")
	err := os.Remove(filePath)
	if err != nil {
		return fmt.Errorf("error deleting file: %w", err)
	}

	return nil
}
