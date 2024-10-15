package handlers

import (
	"encoding/json"
	"net/http"
	"split/repositories"
	"split/services"
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

func (h *FxRateHandler) FetchAndStoreRates(w http.ResponseWriter, r *http.Request) {
	createdFxRates, err := services.FetchAndStoreFxRates(h.currencyRepo, h.repo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("HX-Trigger", "reloadCurrencies")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdFxRates)
}
