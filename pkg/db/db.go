package db

import (
	"TaskWeave/pkg/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err) 
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	db.AutoMigrate(&models.Task{})
	return db, nil
}