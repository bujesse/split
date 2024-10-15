package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"split/models"
	"split/repositories"
	"split/views/partials"
	"strconv"
)

type CurrencyHandler struct {
	repo repositories.CurrencyRepository
}

func NewCurrencyHandler(
	repo repositories.CurrencyRepository,
) *CurrencyHandler {
	return &CurrencyHandler{
		repo,
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
	var currency models.Currency
	if err := json.NewDecoder(r.Body).Decode(&currency); err != nil {
		http.Error(w, "Failed to decode currency", http.StatusBadRequest)
		return
	}

	if err := h.repo.Create(&currency); err != nil {
		http.Error(w, "Failed to save currency", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("HX-Trigger", "reloadCurrencies")
	json.NewEncoder(w).Encode(currency)
	w.WriteHeader(http.StatusCreated)
}

func (h *CurrencyHandler) DeleteCurrency(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := h.repo.Delete(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("HX-Trigger", "reloadCurrencies")
}
