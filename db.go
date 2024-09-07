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
	seedCategories(db)

	return nil
}

func seedCurrencies(db *gorm.DB) {
	var count int64
	db.Model(&models.Currency{}).Count(&count)

	if count == 0 {
		currencies := []models.Currency{
			{
				Code:               "CAD",
				Name:               "Canadian Dollar",
				LatestFxRateUSD:    1.3466401488,
				IsBaseCurrency:     false,
				TwoCharCountryCode: "CA",
			},
			{
				Code:               "EUR",
				Name:               "Euro",
				LatestFxRateUSD:    0.8986501054,
				IsBaseCurrency:     false,
				TwoCharCountryCode: "EU", // No specific ISO 2-char, EU is used informally
			},
			{
				Code:               "GBP",
				Name:               "British Pound Sterling",
				LatestFxRateUSD:    0.7567701029,
				IsBaseCurrency:     false,
				TwoCharCountryCode: "GB",
			},
			{
				Code:               "HKD",
				Name:               "Hong Kong Dollar",
				LatestFxRateUSD:    7.7957609828,
				IsBaseCurrency:     false,
				TwoCharCountryCode: "HK",
			},
			{
				Code:               "JPY",
				Name:               "Japanese Yen",
				LatestFxRateUSD:    144.2905964449,
				IsBaseCurrency:     false,
				TwoCharCountryCode: "JP",
			},
			{
				Code:               "MXN",
				Name:               "Mexican Peso",
				LatestFxRateUSD:    19.6518136887,
				IsBaseCurrency:     false,
				TwoCharCountryCode: "MX",
			},
			{
				Code:               "NZD",
				Name:               "New Zealand Dollar",
				LatestFxRateUSD:    1.602150244,
				IsBaseCurrency:     false,
				TwoCharCountryCode: "NZ",
			},
			{
				Code:               "USD",
				Name:               "United States Dollar",
				LatestFxRateUSD:    1,
				IsBaseCurrency:     true,
				TwoCharCountryCode: "US",
			},
		}

		for _, currency := range currencies {
			db.Create(&currency)
		}

		logger.Debug.Println("🌱 Seeded Currencies")
	}
}

func seedCategories(db *gorm.DB) {
	var count int64
	db.Model(&models.Category{}).Count(&count)

	if count == 0 {
		categories := []models.Category{
			// Recreation
			{Name: "Bars", Type: "Recreation"},
			{Name: "Breweries", Type: "Recreation"},
			{Name: "Dining Out", Type: "Recreation"},
			{Name: "Concerts", Type: "Recreation"},
			{Name: "Events", Type: "Recreation"},
			{Name: "Movies", Type: "Recreation"},
			{Name: "Sports", Type: "Recreation"},
			{Name: "Other", Type: "Recreation"},

			// Groceries
			{Name: "Costco", Type: "Groceries"},
			{Name: "Grocery Store", Type: "Groceries"},
			{Name: "Liquor Store", Type: "Groceries"},
			{Name: "Specialty Store", Type: "Groceries"},

			// Bills
			{Name: "Electric", Type: "Bills"},
			{Name: "Gas", Type: "Bills"},
			{Name: "Trash", Type: "Bills"},
			{Name: "Rent", Type: "Bills"},
			{Name: "Internet", Type: "Bills"},
			{Name: "Other", Type: "Bills"},

			// Uncategorized
			{Name: "General", Type: "Uncategorized"},

			// Travel
			{Name: "Flights", Type: "Travel"},
			{Name: "Hotels", Type: "Travel"},
			{Name: "Parking", Type: "Travel"},
			{Name: "Gas", Type: "Travel"},
			{Name: "Taxi", Type: "Travel"},
			{Name: "Buses/Trains", Type: "Travel"},
			{Name: "Other", Type: "Travel"},

			// Transportation
			{Name: "Car", Type: "Transportation"},
			{Name: "Gas", Type: "Transportation"},
			{Name: "Taxi", Type: "Transportation"},
			{Name: "Other", Type: "Transportation"},

			// Shopping
			{Name: "Clothes", Type: "Shopping"},
			{Name: "Gifts", Type: "Shopping"},
			{Name: "Other", Type: "Shopping"},
		}

		for _, category := range categories {
			db.Create(&category)
		}

		logger.Debug.Println("🌱 Seeded Categories")
	}

}
