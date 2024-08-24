package main

import (
	"log"
	"split/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetConnection() *gorm.DB {
	if db != nil {
		return db
	}

	db, err := gorm.Open(sqlite.Open("split.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("🔥 failed to connect to the database: %s", err.Error())
	}

	log.Println("🚀 Connected Successfully to the Database")

	return db
}

func MakeMigrations() error {
	db := GetConnection()

	err := db.AutoMigrate(
		&models.Category{},
		&models.Currency{},
		&models.Settlement{},
		&models.Expense{},
		&models.ExpenseOwed{},
		&models.User{},
	)
	if err != nil {
		log.Fatalf("failed to migrate database schema: %v", err)
	}

	return nil
}
