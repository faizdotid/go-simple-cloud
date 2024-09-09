package model

import (
	"gorm.io/gorm"
	"time"
)

// `file_extensions` table schema
type FileExtensions struct {
	ID  uint   `gorm:"column:id" json:"id"`
	Ext string `gorm:"column:ext" json:"ext"`
}

// `preview_files` table schema
type PreviewFiles struct {
	ID    uint           `gorm:"column:id" json:"id"`
	Name  string         `gorm:"column:name" json:"name"`
	ExtId uint           `gorm:"column:ext_id" json:"ext_id"`
	Ext   FileExtensions `gorm:"foreignKey:ExtId" json:"ext"`
	Url   string         `gorm:"column:url" json:"url"`
}

// `files` table schema
type Files struct {
	ID            uint         `gorm:"column:id" json:"id"`
	Url           string       `gorm:"column:url" json:"url"`
	Filename      string       `gorm:"column:filename" json:"filename"`
	Filesize      int64        `gorm:"column:filesize" json:"filesize"`
	Path          string       `gorm:"column:path" json:"path"`
	PreviewFileID uint         `gorm:"column:preview_file_id" json:"preview_file_id"`
	PreviewFile   PreviewFiles `gorm:"foreignKey:PreviewFileID" json:"preview_file"`
	CreatedAt     time.Time    `gorm:"column:created_at" json:"created_at"`
	ExpiresAt     time.Time    `gorm:"column:expires_at" json:"expires_at"`
	gorm.DeletedAt
}
