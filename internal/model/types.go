package model

import (
	"gorm.io/gorm"
	"time"
)

type Files struct {
	ID            uint         `gorm:"column:id" json:"id"`
	Filename      string       `gorm:"column:filename" json:"filename"`
	Url           string       `gorm:"column:url" json:"url"`
	Filesize      int64        `gorm:"column:filesize" json:"filesize"`
	Path          string       `gorm:"column:path" json:"path"`
	PreviewFileID uint         `gorm:"column:preview_file_id" json:"preview_file_id"`
	PreviewFile   PreviewFiles `gorm:"foreignKey:PreviewFileID" json:"preview_file"`
	ExpiresAt     time.Time    `gorm:"column:expires_at" json:"expires_at"`
	gorm.DeletedAt
}

type PreviewFiles struct {
	ID   uint   `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	Url  string `gorm:"column:url" json:"url"`
}
