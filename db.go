package main

import (
	"os"
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

	dbURL := os.Getenv("SPLIT_DATABASE_URL")
	db, err := gorm.Open(sqlite.Open(dbURL), &gorm.Config{})
	if err != nil {
		logger.Fatal("🔥 failed to connect to the database: %s", err.Error())
	} else {
		logger.Debug.Println("🚀 Connected Successfully to the Database:", dbURL)
	}

	return db
}

func MakeMigrations() error {
	db := GetConnection()
	// db := GetConnection().Debug()

	err := db.AutoMigrate(
		&models.Category{},
		&models.Currency{},
		&models.Settlement{},
		&models.Expense{},
		&models.ExpenseSplit{},
		&models.User{},
		&models.FxRate{},
	)
	if err != nil {
		logger.Fatal("failed to migrate database schema: %v", err)
	}

	seedCurrencies(db)

	return nil
}

func seedCurrencies(db *gorm.DB) {
	var count int64
	db.Model(&models.Currency{}).Count(&count)

	if count == 0 {
		currencies := []models.Currency{
			{
				Code:            "CAD",
				Name:            "Canadian Dollar",
				LatestFxRateUSD: 1.3466401488,
				IsBaseCurrency:  false,
			},
			{Code: "EUR", Name: "Euro", LatestFxRateUSD: 0.8986501054, IsBaseCurrency: false},
			{
				Code:            "GBP",
				Name:            "British Pound Sterling",
				LatestFxRateUSD: 0.7567701029,
				IsBaseCurrency:  false,
			},
			{
				Code:            "HKD",
				Name:            "Hong Kong Dollar",
				LatestFxRateUSD: 7.7957609828,
				IsBaseCurrency:  false,
			},
			{
				Code:            "JPY",
				Name:            "Japanese Yen",
				LatestFxRateUSD: 144.2905964449,
				IsBaseCurrency:  false,
			},
			{
				Code:            "MXN",
				Name:            "Mexican Peso",
				LatestFxRateUSD: 19.6518136887,
				IsBaseCurrency:  false,
			},
			{
				Code:            "NZD",
				Name:            "New Zealand Dollar",
				LatestFxRateUSD: 1.602150244,
				IsBaseCurrency:  false,
			},
			{Code: "USD", Name: "United States Dollar", LatestFxRateUSD: 1, IsBaseCurrency: true},
		}

		for _, currency := range currencies {
			db.Create(&currency)
		}
	}
}
