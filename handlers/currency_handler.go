package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"split/config/logger"
	"split/models"
	"split/repositories"
	"split/services"
	"split/views/partials"
)

type CurrencyHandler struct {
	repo       repositories.CurrencyRepository
	fxRateRepo repositories.FxRateRepository
}

func NewCurrencyHandler(
	repo repositories.CurrencyRepository,
	fxRateRepo repositories.FxRateRepository,
) *CurrencyHandler {
	return &CurrencyHandler{
		repo,
		fxRateRepo,
	}
}

func (h *CurrencyHandler) GetAllCurrencies(w http.ResponseWriter, r *http.Request) {
	currencies, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	partials.CurrenciesTable(currencies).Render(context.Background(), w)
}

func (h *CurrencyHandler) CreateCurrency(w http.ResponseWriter, r *http.Request) {
	logger.Debug.Println("Creating currency...")

	code := r.FormValue("Code")
	name := r.FormValue("Name")
	twoCharCountryCode := r.FormValue("TwoCharCountryCode")
	isBaseCurrency := r.FormValue("IsBaseCurrency") == "true"

	newCurrency := models.Currency{
		Code:               code,
		Name:               name,
		IsBaseCurrency:     isBaseCurrency,
		TwoCharCountryCode: twoCharCountryCode,
	}

	if err := h.repo.Create(&newCurrency); err != nil {
		http.Error(w, "Failed to save currency", http.StatusInternalServerError)
		return
	}

	_, err := services.FetchAndStoreFxRates(h.repo, h.fxRateRepo)
	if err != nil {
		http.Error(
			w,
			"Failed to fetch and store fx rates for: "+newCurrency.Code,
			http.StatusInternalServerError,
		)
		return
	}

	logger.Debug.Println("Created Currncy with code: ", newCurrency.Code)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("HX-Trigger", "reloadCurrencies")
	json.NewEncoder(w).Encode(newCurrency)
	w.WriteHeader(http.StatusCreated)
}

func (h *CurrencyHandler) DeleteCurrency(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	if err := h.repo.Delete(code); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("HX-Trigger", "reloadCurrencies")
	w.WriteHeader(http.StatusOK)
}
