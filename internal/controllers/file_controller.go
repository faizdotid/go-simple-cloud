package controllers

import (
	"fmt"
	"go-simple-cloud/internal/helpers"
	"go-simple-cloud/internal/model"
	"go-simple-cloud/internal/services"
	"go-simple-cloud/internal/utils"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// NewFileController initializes a new file controller with a given database
func NewFileController(db *gorm.DB) *fileController {
	return &fileController{db: db}
}

// Upload handles file uploads
func (u *fileController) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expirationStr := c.PostForm("expires")
	if expirationStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Expiration time is required"})
		return
	}

	expiresAt := helpers.CreateExpirationTime(helpers.Expiration(expirationStr))
	if expiresAt == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid expiration time"})
		return
	}

	response, err := services.ValidateAndSaveFile(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url := utils.CreateRandomString()
	fileRecord := model.Files{
		Filename:  file.Filename,
		Url:       url,
		Filesize:  file.Size,
		Path:      response["file"].(map[string]interface{})["path"].(string),
		ExpiresAt: time.Unix(expiresAt, 0),
	}
	fileRecordExt := strings.TrimPrefix(path.Ext(fileRecord.Filename), ".")

	var previewFile model.PreviewFiles
	if err := u.db.Model(&model.PreviewFiles{}).Where("name = ?", fileRecordExt).First(&previewFile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fileRecord.PreviewFileID = previewFile.ID
	if err := u.db.Create(&fileRecord).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fileData := u.filesAsMap(&fileRecord)
	fileData["url"] = url
	fileData["preview"] = map[string]interface{}{
		"name": previewFile.Name,
		"url":  previewFile.Url,
	}

	c.JSON(http.StatusOK, gin.H{"file": fileData})
}

// Index lists recent files that have not expired
func (u *fileController) Index(c *gin.Context) {
	var files []model.Files
	if err := u.db.Preload("PreviewFile").Where("expires_at > ?", time.Now()).Order("created_at DESC").Limit(10).Find(&files).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responseFiles := make([]map[string]interface{}, len(files))
	for i, file := range files {
		responseFiles[i] = u.filesAsMap(&file)
		responseFiles[i]["url"] = file.Url
	}

	c.JSON(http.StatusOK, gin.H{"files": responseFiles})
}

// Show displays details of a specific file by URL
func (u *fileController) Show(c *gin.Context) {
	url := c.Param("url")
	var file model.Files
	if err := u.db.Preload("PreviewFile").Where("url = ? AND expires_at > ?", url, time.Now()).First(&file).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"file": u.filesAsMap(&file)})
}

// Download allows users to download a file by URL
func (u *fileController) Download(c *gin.Context) {
	url := c.Param("url")
	var file model.Files
	if err := u.db.Where("url = ? AND expires_at > ?", url, time.Now()).First(&file).Error; err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file.Filename))
	c.Header("Content-Length", fmt.Sprintf("%d", file.Filesize))
	c.Header("Content-Type", "application/octet-stream")
	c.File(file.Path)
}

// filesAsMap formats a file record as a map for JSON response
func (u *fileController) filesAsMap(file *model.Files) map[string]interface{} {
	return map[string]interface{}{
		"filename":   file.Filename,
		"filesize":   file.Filesize,
		"expires_at": file.ExpiresAt.Format(time.RFC3339),
		"preview": map[string]interface{}{
			"name": file.PreviewFile.Name,
			"url":  file.PreviewFile.Url,
		},
	}
}
