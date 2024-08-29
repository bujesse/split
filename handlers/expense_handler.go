package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"split/config/logger"
	"split/models"
	"split/repositories"
	"split/views/components"
	"strconv"
)

type ExpenseHandler struct {
	expenseRepo  repositories.ExpenseRepository
	categoryRepo repositories.CategoryRepository
	currencyRepo repositories.CurrencyRepository
	userRepo     repositories.UserRepository
}

func NewExpenseHandler(
	expenseRepo repositories.ExpenseRepository,
	categoryRepo repositories.CategoryRepository,
	currencyRepo repositories.CurrencyRepository,
	userRepo repositories.UserRepository,
) *ExpenseHandler {
	return &ExpenseHandler{
		expenseRepo,
		categoryRepo,
		currencyRepo,
		userRepo,
	}
}

func (h *ExpenseHandler) CreateExpense(response http.ResponseWriter, r *http.Request) {
	logger.Debug.Println("Creating expense")

	title := r.FormValue("title")
	amountStr := r.FormValue("amount")
	amount, _ := strconv.ParseFloat(amountStr, 64)
	notes := r.FormValue("notes")
	currencyCode := r.FormValue("currencyCode")
	categoryID := r.FormValue("categoryID")
	var parsedCatID *uint
	if catID, err := strconv.ParseUint(categoryID, 10, 64); err == nil && categoryID != "" {
		parsedID := uint(catID)
		parsedCatID = &parsedID
	}

	claims, _ := GetCurrentUserClaims(r)
	userID := uint(claims.UserID)

	expense := models.Expense{
		Title:        title,
		Amount:       amount,
		CreatedByID:  userID,
		Notes:        notes,
		CurrencyCode: currencyCode,
		CategoryID:   parsedCatID,
	}

	if err := h.expenseRepo.Create(&expense); err != nil {
		http.Error(response, "Failed to save expense", http.StatusInternalServerError)
		return
	}

	logger.Debug.Println("Created Expense with ID: ", expense.ID)

	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("HX-Trigger", "reloadExpenses")
	json.NewEncoder(response).Encode(expense)
	response.WriteHeader(http.StatusCreated)
	// response.Header().Set("Content-Type", "text/html")
	// components.ExpensesTable(expenses).Render(context.Background(), response)
}

func (h *ExpenseHandler) GetAllExpenses(response http.ResponseWriter, request *http.Request) {
	expenses, err := h.expenseRepo.GetAll()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "text/html")
	components.ExpensesTable(expenses).Render(context.Background(), response)
}

func (h *ExpenseHandler) CreateNewExpense(w http.ResponseWriter, request *http.Request) {
	categories, _ := h.categoryRepo.GetAll()
	currencies, _ := h.currencyRepo.GetAll()
	users, _ := h.userRepo.GetAll()

	components.Modal(components.ExpenseForm(
		nil,
		categories,
		currencies,
		users,
	)).Render(request.Context(), w)
}

func (h *ExpenseHandler) EditExpenseByID(w http.ResponseWriter, request *http.Request) {
	idStr := request.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	expense, err := h.expenseRepo.GetByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	categories, _ := h.categoryRepo.GetAll()
	currencies, _ := h.currencyRepo.GetAll()
	users, _ := h.userRepo.GetAll()

	components.Modal(components.ExpenseForm(
		expense,
		categories,
		currencies,
		users,
	)).Render(context.Background(), w)
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

	expense, err := h.expenseRepo.GetByID(uint(id))

	expense.Title = title
	expense.Amount, _ = strconv.ParseFloat(amount, 64)
	expense.Notes = notes
	expense.CurrencyCode = currencyCode
	if catID, err := strconv.ParseUint(categoryID, 10, 64); err == nil && categoryID != "" {
		parsedCatID := uint(catID)
		expense.CategoryID = &parsedCatID
	} else {
		expense.CategoryID = nil
	}

	if err := h.expenseRepo.Update(expense); err != nil {
		http.Error(w, "Failed to update expense: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Trigger", "reloadExpenses")
	w.WriteHeader(http.StatusOK)
}
