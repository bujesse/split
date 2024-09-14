package handlers

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

type FxRateHandler struct {
	repo         repositories.FxRateRepository
	currencyRepo repositories.CurrencyRepository
}

func NewFxRateHandler(
	repo repositories.FxRateRepository,
	currencyRepo repositories.CurrencyRepository,
) *FxRateHandler {
	return &FxRateHandler{
		repo,
		currencyRepo,
	}
}

type ApiResponse struct {
	Success   bool               `json:"success"`
	Terms     string             `json:"terms"`
	Privacy   string             `json:"privacy"`
	Timestamp int64              `json:"timestamp"`
	Date      string             `json:"date"`
	Base      string             `json:"base"`
	Rates     map[string]float64 `json:"rates"`
}

func (h *FxRateHandler) FetchAndStoreRates(w http.ResponseWriter, r *http.Request) {
	apiKey := os.Getenv("SPLIT_FX_RATES_API_TOKEN")
	currencies, _ := h.currencyRepo.GetAll()
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
		http.Error(w, "Failed to fetch fx rates", http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	var apiResponse ApiResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		http.Error(w, "Failed to decode fx rate response", http.StatusInternalServerError)
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
		err := h.currencyRepo.Update(currency)
		if err != nil {
			http.Error(
				w,
				"Failed to update currency rate: "+err.Error(),
				http.StatusInternalServerError,
			)
			return
		}

		date, _ := helpers.ConvertToServerTime(apiResponse.Date)

		fxRate := models.FxRate{
			FromCurrencyCode: apiResponse.Base,
			ToCurrencyCode:   currencyCode,
			Rate:             rate,
			Date:             date,
		}

		if err := h.repo.Create(&fxRate); err != nil {
			http.Error(w, "Failed to save fx rate: "+err.Error(), http.StatusInternalServerError)
			return
		}

		createdFxRates = append(createdFxRates, fxRate)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdFxRates)
}
