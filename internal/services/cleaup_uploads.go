package services

import (
	"go-simple-cloud/internal/model"
	"os"

	"gorm.io/gorm"
)

func NewCleanupUploadsService(db *gorm.DB) *CleanupUploadsService {
	return &CleanupUploadsService{
		db: db,
	}
}

func (s *CleanupUploadsService) deleteFile(filepath string) error {
	if err := os.Remove(filepath); err != nil {
		return err
	}
	return nil
}

func (s *CleanupUploadsService) do() error {
	var files []model.Files
	if err := s.db.Preload("PreviewFile").Where("expires_at < NOW()").Where("deleted_at IS NULL").Find(&files).Error; err != nil {
		return err
	}
	for _, file := range files {
		if err := s.deleteFile(file.Path); err != nil {
			return err
		}
		if err := s.db.Find(&file).Delete(&file).Error; err != nil {
			return err
		}
	}
	return nil

}

func (s *CleanupUploadsService) Do() error {
	return s.do()
}
