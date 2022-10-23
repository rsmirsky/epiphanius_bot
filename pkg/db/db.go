package db

import (
	"epiphanius_bot/pkg/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect(dsn string) *gorm.DB {
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("connect to db error", err)
	}
	db = d
	return db
}

func RunMigrates() {
	if err := db.AutoMigrate(&models.Holiday{}, &models.Users{}); err != nil {
		fmt.Println("migration error", err)
	}
}
