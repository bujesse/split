package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"split/helpers"
	"split/models"
	"split/repositories"
	"strings"
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

func FetchAndStoreFxRates(
	currencyRepo repositories.CurrencyRepository,
	fxRateRepo repositories.FxRateRepository,
) ([]models.FxRate, error) {
	apiKey := os.Getenv("SPLIT_FX_RATES_API_TOKEN")
	currencies, err := currencyRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get currencies: %w", err)
	}

	currencyCodes := make([]string, len(currencies))
	for i, currency := range currencies {
		currencyCodes[i] = currency.Code
	}

	joinedCodes := strings.Join(currencyCodes, ",")

	url := fmt.Sprintf(
		"https://api.fxratesapi.com/latest?api_key=%s&currencies=%s",
		apiKey,
		joinedCodes,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch fx rates: %w", err)
	}
	defer resp.Body.Close()

	var apiResponse ApiResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to decode fx rate response: %w", err)
	}

	currencyMap := make(map[string]*models.Currency)
	for i := range currencies {
		currencyMap[currencies[i].Code] = &currencies[i]
	}

	createdFxRates := []models.FxRate{}
	for currencyCode, rate := range apiResponse.Rates {
		currency, found := currencyMap[currencyCode]
		if currency.Code == "USD" || !found {
			continue
		}

		currency.LatestFxRateUSD = rate
		if err := currencyRepo.Update(currency); err != nil {
			return nil, fmt.Errorf("failed to update currency rate: %w", err)
		}

		date, err := helpers.ParseDate(apiResponse.Date)
		if err != nil {
			return nil, fmt.Errorf("failed to parse date: %w", err)
		}

		fxRate := models.FxRate{
			FromCurrencyCode: apiResponse.Base,
			ToCurrencyCode:   currencyCode,
			Rate:             rate,
			Date:             *date,
		}

		if err := fxRateRepo.Create(&fxRate); err != nil {
			return nil, fmt.Errorf("failed to save fx rate: %w", err)
		}

		createdFxRates = append(createdFxRates, fxRate)
	}

	return createdFxRates, nil
}
