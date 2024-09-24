package handlers

import (
	"encoding/json"
	"net/http"
	"split/config/logger"
	"split/helpers"
	"split/models"
	"split/repositories"
	"split/views/components"
	"strconv"
)

type SettlementHandler struct {
	repo         repositories.SettlementRepository
	currencyRepo repositories.CurrencyRepository
	userRepo     repositories.UserRepository
	expenseRepo  repositories.ExpenseRepository
}

func NewSettlementHandler(
	repo repositories.SettlementRepository,
	currencyRepo repositories.CurrencyRepository,
	userRepo repositories.UserRepository,
	expenseRepo repositories.ExpenseRepository,
) *SettlementHandler {
	return &SettlementHandler{
		repo,
		currencyRepo,
		userRepo,
		expenseRepo,
	}
}

func (h *SettlementHandler) CreateSettlement(w http.ResponseWriter, r *http.Request) {
	logger.Debug.Println("Creating settlement")

	amount, _ := strconv.ParseFloat(r.FormValue("Amount"), 64)
	settledByID := r.FormValue("SettledByID")
	parsedSettledByID, _ := helpers.StringToUint(settledByID)
	settledToZero := r.FormValue("SettledToZero") == "true"

	settlement := models.Settlement{
		SettledByID:   parsedSettledByID,
		Amount:        amount,
		CurrencyCode:  r.FormValue("CurrencyCode"),
		Notes:         r.FormValue("Notes"),
		SettledToZero: settledToZero,
	}

	if err := h.repo.Create(&settlement); err != nil {
		http.Error(w, "Failed to save settlement", http.StatusInternalServerError)
		return
	}

	logger.Debug.Println("Created Settlement with ID: ", settlement.ID)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("HX-Trigger", "reloadExpenses")
	json.NewEncoder(w).Encode(settlement)
	w.WriteHeader(http.StatusCreated)
}

func (h *SettlementHandler) CreateNewSettlementPartial(
	w http.ResponseWriter,
	r *http.Request,
) {
	currencies, _ := h.currencyRepo.GetAll()
	users, _ := h.userRepo.GetAll()
	expenses, _ := h.expenseRepo.GetExpensesSinceLastSettlement()
	settlements, _ := h.repo.GetAllSinceLastSettlement()
	owedDetails := helpers.CalculateOwedDetails(expenses, settlements)

	components.SettlementsForm(
		nil,
		owedDetails,
		currencies,
		users,
	).Render(r.Context(), w)
}

func (h *SettlementHandler) DeleteSettlement(response http.ResponseWriter, request *http.Request) {
	idStr := request.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(response, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := h.repo.Delete(uint(id)); err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("HX-Trigger", "reloadExpenses")
	response.WriteHeader(http.StatusNoContent)
}

func (h *SettlementHandler) EditSettlementByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	settlement, err := h.repo.GetByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	currencies, _ := h.currencyRepo.GetAll()
	users, _ := h.userRepo.GetAll()
	expenses, _ := h.expenseRepo.GetExpensesSinceLastSettlement()

	settlements, _ := h.repo.GetAllSinceLastSettlement()
	owedDetails := helpers.CalculateOwedDetails(expenses, settlements)
	components.SettlementsForm(settlement, owedDetails, currencies, users).
		Render(r.Context(), w)
}

func (h *SettlementHandler) UpdateSettlement(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	amount, _ := strconv.ParseFloat(r.FormValue("Amount"), 64)
	settledByID := r.FormValue("SettledByID")
	parsedSettledByID, _ := helpers.StringToUint(settledByID)
	settledToZero := r.FormValue("SettledToZero") == "true"

	settlement, err := h.repo.GetByID(uint(id))

	settlement.SettledByID = parsedSettledByID
	settlement.Amount = amount
	settlement.CurrencyCode = r.FormValue("CurrencyCode")
	settlement.Notes = r.FormValue("Notes")
	settlement.SettledToZero = settledToZero

	if err := h.repo.Update(settlement); err != nil {
		http.Error(w, "Failed to update settlement", http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Trigger", "reloadExpenses")
	w.WriteHeader(http.StatusOK)
}
