package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"split/config"
	"split/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ApiResponse struct {
	Success   bool               `json:"success"`
	Terms     string             `json:"terms"`
	Privacy   string             `json:"privacy"`
	Timestamp int64              `json:"timestamp"`
	Date      string             `json:"date"`
	Base      string             `json:"base"`
	Rates     map[string]float64 `json:"rates"`
}

func FetchAndStoreRates(db *gorm.DB) error {
	apiKey := config.GetFxRatesApiToken()
	url := fmt.Sprintf("https://api.fxratesapi.com/latest?api_key=%s", apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var apiResponse ApiResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		return err
	}

	baseCurrency := models.Currency{
		Code:           apiResponse.Base,
		IsBaseCurrency: true,
	}

	// Ensure the base currency exists or create it
	db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&baseCurrency)

	for currencyCode, rate := range apiResponse.Rates {
		currency := models.Currency{
			Code: currencyCode,
		}

		// Ensure the target currency exists or create it
		db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&currency)

		dbRate := models.FxRate{
			FromCurrencyCode: apiResponse.Base,
			ToCurrencyCode:   currencyCode,
			Rate:             rate,
			Date:             apiResponse.Date,
		}

		// Store the rate
		db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&dbRate)
	}

	return nil
}
