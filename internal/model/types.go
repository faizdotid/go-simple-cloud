package model

import "time"

type Files struct {
	Filename  string    `gorm:"column:filename" json:"filename"`
	Url       string    `gorm:"column:url" json:"url"`
	Filesize  int64     `gorm:"column:filesize" json:"filesize"`
	Path      string    `gorm:"column:path" json:"path"`
	ExpiresAt time.Time `gorm:"column:expires_at" json:"expires_at"`
}
