package files

import (
	"fmt"
	"go-simple-cloud/internal/helpers"
	"go-simple-cloud/internal/model"
	"go-simple-cloud/internal/services"
	"go-simple-cloud/internal/utils"
	"net/http"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewFileController(db *gorm.DB) *fileController {
	return &fileController{db: db}
}

func (u *fileController) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expiresIn := c.PostForm("expires_in")
	if expiresIn == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "expiration time is required"})
		return
	}

	expiresAt, err := helpers.CreateExpirationTime(helpers.Expiration(expiresIn))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := services.ValidateAndSaveFile(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url := utils.CreateRandomString()
	record := model.Files{
		Filename:  file.Filename,
		Url:       url,
		Filesize:  file.Size,
		Path:      response["file"].(map[string]interface{})["path"].(string),
		ExpiresAt: time.Unix(expiresAt, 0),
	}
	fileExt := path.Ext(file.Filename)

	var previewFile model.PreviewFiles

	// SELECT `preview_files`.`id` FROM `preview_files` JOIN `file_extensions` ON `preview_files`.`ext_id` = `file_extensions`.`id` WHERE ext = ?
	err = u.db.Model(&model.PreviewFiles{}).
		Joins("JOIN file_extensions ON preview_files.ext_id = file_extensions.id").
		Where("file_extensions.ext = ?", fileExt).
		First(&previewFile).Error

	if err == gorm.ErrRecordNotFound {
		if err := u.db.Find(&model.PreviewFiles{}).Where("id = 0").First(&previewFile).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	record.PreviewFileID = previewFile.ID
	if err := u.db.Create(&record).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fileData := u.filesAsMap(&record)
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
