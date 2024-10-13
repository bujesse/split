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
		&models.ScheduledExpense{},
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
			{Name: "Bars", Type: "Recreation", Icon: "martini-glass-citrus"},
			{Name: "Liquor Store", Type: "Recreation", Icon: "martini-glass-citrus"},
			{Name: "Breweries", Type: "Recreation", Icon: "beer-mug-empty"},
			{Name: "Dining Out", Type: "Recreation", Icon: "beer-mug-empty"},
			{Name: "Concerts", Type: "Recreation", Icon: "ticket"},
			{Name: "Events", Type: "Recreation", Icon: "ticket"},
			{Name: "Movies", Type: "Recreation", Icon: "film"},
			{Name: "Sports", Type: "Recreation", Icon: "table-tennis-paddle-ball"},
			{Name: "Other", Type: "Recreation", Icon: "beer-mug-empty"},

			// Groceries
			{Name: "Costco", Type: "Food", Icon: "cart-shopping"},
			{Name: "Grocery Store", Type: "Food", Icon: "basket-shopping"},
			{Name: "Ordering Food", Type: "Food", Icon: "bag-shopping"},

			// Bills
			{Name: "Electric", Type: "Bills", Icon: "lightbulb"},
			{Name: "Gas", Type: "Bills", Icon: "fire-flame-simple"},
			{Name: "Trash", Type: "Bills", Icon: "trash"},
			{Name: "Rent", Type: "Bills", Icon: "house"},
			{Name: "Internet", Type: "Bills", Icon: "wifi"},
			{Name: "Other", Type: "Bills", Icon: "money-bills"},

			// Uncategorized
			{Name: "General", Type: "Uncategorized", Icon: "dollar"},

			// Travel
			{Name: "Flights", Type: "Travel", Icon: "plane"},
			{Name: "Hotels", Type: "Travel", Icon: "hotel"},
			{Name: "Parking", Type: "Travel", Icon: "square-parking"},
			{Name: "Gas", Type: "Travel", Icon: "gas-pump"},
			{Name: "Taxi", Type: "Travel", Icon: "taxi"},
			{Name: "Buses/Trains", Type: "Travel", Icon: "train-subway"},
			{Name: "Other", Type: "Travel", Icon: "suitcase"},

			// Transportation
			{Name: "Car", Type: "Transportation", Icon: "car-side"},
			{Name: "Gas", Type: "Transportation", Icon: "gas-pump"},
			{Name: "Taxi", Type: "Transportation", Icon: "taxi"},
			{Name: "Other", Type: "Transportation", Icon: "car"},

			// Shopping
			{Name: "Clothes", Type: "Shopping", Icon: "shirt"},
			{Name: "Gifts", Type: "Shopping", Icon: "gift"},
			{Name: "Online", Type: "Shopping", Icon: "globe"},
			{Name: "Other", Type: "Shopping", Icon: "store"},
		}

		for _, category := range categories {
			db.Create(&category)
		}

		logger.Debug.Println("🌱 Seeded Categories")
	}

}
