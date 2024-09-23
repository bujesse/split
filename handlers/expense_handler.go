package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"sort"
	"split/config/logger"
	"split/helpers"
	"split/models"
	"split/repositories"
	"split/views/components"
	"split/views/partials"
	"strconv"
	"time"
)

type ExpenseHandler struct {
	expenseRepo    repositories.ExpenseRepository
	categoryRepo   repositories.CategoryRepository
	currencyRepo   repositories.CurrencyRepository
	userRepo       repositories.UserRepository
	settlementRepo repositories.SettlementRepository
}

func NewExpenseHandler(
	expenseRepo repositories.ExpenseRepository,
	categoryRepo repositories.CategoryRepository,
	currencyRepo repositories.CurrencyRepository,
	userRepo repositories.UserRepository,
	settlementRepo repositories.SettlementRepository,
) *ExpenseHandler {
	return &ExpenseHandler{
		expenseRepo,
		categoryRepo,
		currencyRepo,
		userRepo,
		settlementRepo,
	}
}

func (h *ExpenseHandler) CreateExpense(response http.ResponseWriter, r *http.Request) {
	logger.Debug.Println("Creating expense")

	title := r.FormValue("title")
	amountStr := r.FormValue("amount")
	amount, _ := strconv.ParseFloat(amountStr, 64)
	SplitType := r.FormValue("SplitType")
	SplitValueStr := r.FormValue("SplitValue")
	SplitValue, _ := strconv.ParseFloat(SplitValueStr, 64)
	notes := r.FormValue("notes")
	currencyCode := r.FormValue("currencyCode")
	paidByID := r.FormValue("paidByID")
	parsedPaidByID, _ := helpers.StringToUint(paidByID)
	splitBy := r.FormValue("splitByID")
	parsedSplitByID, _ := helpers.StringToUint(splitBy)
	categoryID := r.FormValue("categoryID")
	parsedCatID, _ := helpers.StringToUintPointer(categoryID)

	claims, _ := GetCurrentUserClaims(r)
	currentUserID := uint(claims.UserID)

	expense := models.Expense{
		Title:        title,
		Amount:       amount,
		CreatedByID:  currentUserID,
		Notes:        notes,
		CurrencyCode: currencyCode,
		CategoryID:   parsedCatID,
		PaidByID:     parsedPaidByID,
		ExpenseSplits: []models.ExpenseSplit{
			{
				UserID:       parsedSplitByID,
				SplitType:    models.SplitType(SplitType),
				SplitValue:   SplitValue,
				CurrencyCode: currencyCode,
			},
		},
	}

	if expense.CategoryID == nil {
		defaultCategory, _ := h.categoryRepo.GetByName("General")
		expense.CategoryID = &defaultCategory.ID
	}

	if err := h.expenseRepo.CreateExpense(&expense); err != nil {
		http.Error(response, "Failed to save expense", http.StatusInternalServerError)
		return
	}

	logger.Debug.Println("Created Expense with ID: ", expense.ID)

	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("HX-Trigger", "reloadExpenses")
	json.NewEncoder(response).Encode(expense)
	response.WriteHeader(http.StatusCreated)
}

// GetAllExpenses returns all expenses and settlements together, sorted by date descending
func (h *ExpenseHandler) GetAllExpenses(response http.ResponseWriter, request *http.Request) {
	expenses, err := h.expenseRepo.GetExpensesSinceLastSettlement()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "text/html")

	var entries []interface{}
	for _, expense := range expenses {
		entries = append(entries, expense)
	}

	settlements, _ := h.settlementRepo.GetAll()
	for _, settlement := range settlements {
		entries = append(entries, settlement)
	}

	sort.Slice(entries, func(i, j int) bool {
		var dateI, dateJ time.Time

		switch v := entries[i].(type) {
		case repositories.ExpenseWithFxRate:
			dateI = v.PaidDate
		case models.Settlement:
			dateI = v.SettlementDate
		}

		switch v := entries[j].(type) {
		case repositories.ExpenseWithFxRate:
			dateJ = v.PaidDate
		case models.Settlement:
			dateJ = v.SettlementDate
		}

		return dateI.After(dateJ)
	})

	categories, _ := h.categoryRepo.GetAll()
	partials.ExpensesTable(entries, categories).Render(context.Background(), response)
}

func (h *ExpenseHandler) GetStats(response http.ResponseWriter, request *http.Request) {
	expenses, err := h.expenseRepo.GetExpensesWithFxRate()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "text/html")
	settlements, _ := h.settlementRepo.GetAll()
	components.Stats(expenses, settlements).Render(context.Background(), response)
}

func (h *ExpenseHandler) CreateNewExpensePartial(w http.ResponseWriter, request *http.Request) {
	categories, _ := h.categoryRepo.GetAll()
	currencies, _ := h.currencyRepo.GetAll()
	users, _ := h.userRepo.GetAll()

	components.ExpenseForm(
		nil,
		categories,
		currencies,
		users,
	).Render(request.Context(), w)
}

func (h *ExpenseHandler) EditExpenseByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	expense, err := h.expenseRepo.GetByID(uint(id), "ExpenseSplits", "Currency", "Category")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	categories, _ := h.categoryRepo.GetAll()
	currencies, _ := h.currencyRepo.GetAll()
	users, _ := h.userRepo.GetAll()

	components.ExpenseForm(
		expense,
		categories,
		currencies,
		users,
	).Render(r.Context(), w)
}

func (h *ExpenseHandler) UpdateExpense(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	amount := r.FormValue("amount")
	notes := r.FormValue("notes")
	currencyCode := r.FormValue("currencyCode")
	categoryID := r.FormValue("categoryID")
	SplitType := r.FormValue("SplitType")
	SplitValue := r.FormValue("SplitValue")
	paidByID := r.FormValue("paidByID")
	parsedPaidByID, _ := helpers.StringToUint(paidByID)
	splitBy := r.FormValue("splitByID")
	parsedSplitByID, _ := helpers.StringToUint(splitBy)

	expense, err := h.expenseRepo.GetByID(uint(id), "ExpenseSplits")

	expense.Title = title
	expense.Amount, _ = strconv.ParseFloat(amount, 64)
	expense.Notes = notes
	expense.CurrencyCode = currencyCode
	expense.PaidByID = parsedPaidByID
	if catID, err := helpers.StringToUintPointer(categoryID); err == nil && categoryID != "" {
		expense.CategoryID = catID
	} else {
		defaultCategory, _ := h.categoryRepo.GetByName("General")
		expense.CategoryID = &defaultCategory.ID
	}

	// TODO: Make this work for multiple splits (currently only works for one split)
	for i := range expense.ExpenseSplits {
		expense.ExpenseSplits[i].UserID = parsedSplitByID
		expense.ExpenseSplits[i].SplitType = models.SplitType(SplitType)
		expense.ExpenseSplits[i].SplitValue, _ = strconv.ParseFloat(SplitValue, 64)
		expense.ExpenseSplits[i].CurrencyCode = currencyCode
	}

	if err := h.expenseRepo.UpdateExpense(expense); err != nil {
		http.Error(w, "Failed to update expense: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Trigger", "reloadExpenses")
	w.WriteHeader(http.StatusOK)
}

func (h *ExpenseHandler) DeleteExpense(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)

	expense, err := h.expenseRepo.GetByID(uint(id))
	if err != nil {
		http.Error(w, "Failed to find expense", http.StatusNotFound)
		return
	}

	if err := h.expenseRepo.DeleteExpense(expense); err != nil {
		http.Error(w, "Failed to delete expense", http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Trigger", "reloadExpenses")
	w.WriteHeader(http.StatusOK)
}
