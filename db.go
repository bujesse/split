package main

import (
	"split/config/logger"
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
		logger.Fatal("🔥 failed to connect to the database: %s", err.Error())
	}

	logger.Debug.Println("🚀 Connected Successfully to the Database")

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
		&models.FxRate{},
	)
	if err != nil {
		logger.Fatal("failed to migrate database schema: %v", err)
	}

	return nil
}
