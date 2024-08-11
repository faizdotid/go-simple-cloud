package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"os"
	"time"
)

// FUCKING IDIOTS, I DONT KNOW GIN HAVE FUNCTION SAVE UPLOADED FILE
func ValidateAndSaveFile(file *multipart.FileHeader) (gin.H, error) {

	// checking file size
	if file.Size > MAX_FILE_SIZE {
		return nil, ErrMaxFileSize
	}

	// open file
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	// create temp name and destination file
	temp := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	filePath := fmt.Sprintf("%s/%s", UploadsPath, temp)
	dst, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	// copy file content
	if _, err := io.Copy(dst, src); err != nil {
		return nil, err
	}

	// return
	return gin.H{
		"message": "File uploaded successfully",
		"file": map[string]interface{}{
			"filename": file.Filename,
			"filesize": file.Size,
			"path":     filePath,
		},
	}, nil
}
