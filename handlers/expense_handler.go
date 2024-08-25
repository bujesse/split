package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"split/config/logger"
	"split/models"
	"split/services"
	"split/views/components"
	"strconv"
)

type ExpenseHandler struct {
	Service services.ExpenseService
}

func NewExpenseHandler(service services.ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{Service: service}
}

func (h *ExpenseHandler) CreateExpense(response http.ResponseWriter, request *http.Request) {
	logger.Debug.Println("Creating expense")

	title := request.FormValue("title")
	amountStr := request.FormValue("amount")
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		http.Error(response, "Invalid amount", http.StatusBadRequest)
		return
	}

	expense := models.Expense{
		Title:  title,
		Amount: amount,
	}

	if err := h.Service.CreateExpense(&expense); err != nil {
		http.Error(response, "Failed to save expense", http.StatusInternalServerError)
		return
	}

	logger.Debug.Println("Created Expense with ID: ", expense.ID)

	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("HX-Trigger", "newExpense")
	json.NewEncoder(response).Encode(expense)
	response.WriteHeader(http.StatusCreated)
	// response.Header().Set("Content-Type", "text/html")
	// components.ExpensesTable(expenses).Render(context.Background(), response)
}

func (h *ExpenseHandler) GetAllExpenses(response http.ResponseWriter, request *http.Request) {
	expenses, err := h.Service.GetAllExpenses()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "text/html")
	components.ExpensesTable(expenses).Render(context.Background(), response)
}

func (h *ExpenseHandler) GetExpenseByID(w http.ResponseWriter, request *http.Request) {
	idStr := request.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	expense, err := h.Service.GetExpenseByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(expense)
}

func (h *ExpenseHandler) UpdateExpense(w http.ResponseWriter, request *http.Request) {
	var expense models.Expense
	if err := json.NewDecoder(request.Body).Decode(&expense); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Service.UpdateExpense(&expense); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(expense)
}
