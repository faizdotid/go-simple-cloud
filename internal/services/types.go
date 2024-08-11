package services

import (
	"errors"
	"os"
)

var (
	UploadsPath    = "assets/uploads"
	ErrMaxFileSize = errors.New("maximum file size is 10MB")

	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")
)
