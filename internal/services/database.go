package services

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "gorm.io/gorm/logger"
)

func NewDatabaseService() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)
	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{
			// Logger: logger.Default.LogMode(logger.Info), // Activate this when debugging
		},
	)
	if err != nil {
		return nil, err
	}
	log.Println("[DatabaseService] Database connection established")
	return db, nil
}
