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

func (h *CurrencyHandler) ToggleCurrency(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	currency, err := h.repo.GetByCode(code)
	if err != nil {
		http.Error(w, "Failed to find currency with code: "+code, http.StatusNotFound)
		return
	}

	currency.IsActive = !currency.IsActive

	if err := h.repo.Update(currency); err != nil {
		http.Error(w, "Failed to update currency", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
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
		logger.Error.Println("Failed to save currency: ", err)
		http.Error(w, "Failed to save currency", http.StatusInternalServerError)
		return
	}

	// If the new currency is not able to be fetched, delete it
	_, err := services.FetchAndStoreFxRates(h.repo, h.fxRateRepo)
	if err != nil {
		logger.Error.Println("Error when fetching and storing fx rates: ", err)
		h.repo.Delete(&newCurrency)
		http.Error(
			w,
			"Failed to fetch and store fx rates: "+err.Error(),
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

	currency, err := h.repo.GetByCode(code)
	if err != nil {
		http.Error(w, "Failed to find currency", http.StatusNotFound)
		return
	}

	if err := h.repo.Delete(currency); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("HX-Trigger", "reloadCurrencies")
	w.WriteHeader(http.StatusOK)
}
