package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"sort"
	"split/config/logger"
	"split/helpers"
	"split/jobs"
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

func (h *ExpenseHandler) doCreateExpense(
	w http.ResponseWriter,
	r *http.Request,
	isTemplate bool,
) *models.Expense {
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
		IsTemplate: isTemplate,
	}

	if expense.CategoryID == nil {
		defaultCategory, _ := h.categoryRepo.GetByName("General")
		expense.CategoryID = &defaultCategory.ID
	}

	if err := h.expenseRepo.CreateExpense(&expense); err != nil {
		http.Error(w, "Failed to save expense", http.StatusInternalServerError)
		return nil
	}

	logger.Debug.Println("Created Expense with ID: ", expense.ID)

	return &expense
}

func (h *ExpenseHandler) CreateExpense(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("IsScheduled") == "true" {
		scheduledExpense, _ := h.doCreateScheduledExpense(w, r)
		if scheduledExpense.StartDate.After(time.Now()) {
			logger.Debug.Println("Scheduled expense is in the future, not creating expense")
			w.WriteHeader(http.StatusCreated)
			return
		}
	}

	expense := h.doCreateExpense(w, r, false)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("HX-Trigger", "reloadExpenses")
	json.NewEncoder(w).Encode(expense)
	w.WriteHeader(http.StatusCreated)
}

// GetExpenses returns all expenses and settlements together, sorted by date descending
func (h *ExpenseHandler) GetExpenses(response http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	offsetParam := query.Get("offset")
	offset, err := strconv.Atoi(offsetParam)
	if err != nil {
		offset = 0
	}

	logger.Debug.Println("Getting expenses, offset:", offset)
	var expenses []repositories.ExpenseWithFxRate
	if offset == 0 {
		expenses, err = h.expenseRepo.GetExpensesSinceLastSettlement()
	} else {
		expenses, err = h.expenseRepo.GetExpensesBetweenZeros(offset)
	}
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "text/html")

	var entries []interface{}
	for _, expense := range expenses {
		entries = append(entries, expense)
	}

	var settlements []models.Settlement
	if offset == 0 {
		settlements, _ = h.settlementRepo.GetAllSinceLastZeroSettlement()
	} else {
		settlements, _ = h.settlementRepo.GetSettlementsBetweenZeros(offset)
	}
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

	numZeroSettlements, _ := h.settlementRepo.GetNumZeroSettlements()
	isLastOffset := offset >= int(numZeroSettlements)
	categories, _ := h.categoryRepo.GetAll()
	partials.ExpensesTable(entries, categories, isLastOffset).Render(context.Background(), response)
}

func (h *ExpenseHandler) GetStats(response http.ResponseWriter, request *http.Request) {
	expenses, err := h.expenseRepo.GetExpensesSinceLastSettlement()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "text/html")
	settlements, _ := h.settlementRepo.GetAllSinceLastZeroSettlement()
	components.Stats(expenses, settlements).Render(context.Background(), response)
}

func (h *ExpenseHandler) CreateNewExpensePartial(w http.ResponseWriter, request *http.Request) {
	categories, _ := h.categoryRepo.GetAll()
	currencies, _ := h.currencyRepo.GetAllActive()
	users, _ := h.userRepo.GetAll()

	components.ExpenseForm(
		nil,
		categories,
		currencies,
		users,
		nil,
	).Render(request.Context(), w)
}

func (h *ExpenseHandler) EditExpenseByIDPartial(w http.ResponseWriter, r *http.Request) {
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
	currencies, _ := h.currencyRepo.GetAllActive()
	users, _ := h.userRepo.GetAll()

	components.ExpenseForm(
		expense,
		categories,
		currencies,
		users,
		nil,
	).Render(r.Context(), w)
}

func (h *ExpenseHandler) EditScheduledExpenseByIDPartial(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	scheduledExpense, err := h.expenseRepo.GetScheduledExpenseByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	categories, _ := h.categoryRepo.GetAll()
	currencies, _ := h.currencyRepo.GetAllActive()
	users, _ := h.userRepo.GetAll()

	components.ExpenseForm(
		&scheduledExpense.TemplateExpense,
		categories,
		currencies,
		users,
		scheduledExpense,
	).Render(r.Context(), w)
}

func (h *ExpenseHandler) UpdateScheduledExpense(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	scheduledExpense, err := h.expenseRepo.GetScheduledExpenseByID(uint(id))
	expense, err := h.expenseRepo.GetByID(scheduledExpense.TemplateExpenseID, "ExpenseSplits")

	expense = h.doUpdateExpense(w, r, expense)

	recurrenceType := r.FormValue("RecurrenceType")
	recurrenceInterval, _ := strconv.Atoi(r.FormValue("RecurrenceInterval"))
	startDate, err := helpers.ParseDate(r.FormValue("StartDate"))
	if err != nil {
		logger.Error.Println("Failed to parse start date:", err)
	}

	endDateStr := r.FormValue("EndDate")
	var endDate *time.Time
	if endDateStr == "" {
		endDate = nil
	} else {
		endDate, err = helpers.ParseDate(endDateStr)
		if err != nil {
			logger.Error.Println("Failed to parse end date:", err)
			return
		}
	}

	if err != nil {
		http.Error(w, "Failed to find scheduled expense", http.StatusNotFound)
		return
	}

	scheduledExpense.RecurrenceType = models.RecurrenceTypes(recurrenceType)
	scheduledExpense.RecurrenceInterval = recurrenceInterval
	scheduledExpense.StartDate = *startDate
	scheduledExpense.EndDate = endDate
	now := time.Now()
	scheduledExpense.NextDueDate = &now

	nextDueDate := jobs.CalculateNextDueDate(scheduledExpense)

	if err := h.expenseRepo.UpdateScheduledExpense(scheduledExpense); err != nil {
		http.Error(
			w,
			"Failed to update scheduled expense: "+err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	logger.Info.Println("Created scheduled expense: ", scheduledExpense.TemplateExpense.Title)
	logger.Info.Println("Next due date: ", nextDueDate)

	w.Header().Set("HX-Trigger", "reloadScheduledExpenses")
	w.WriteHeader(http.StatusOK)
}

func (h *ExpenseHandler) DeleteScheduledExpense(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)

	scheduledExpense, err := h.expenseRepo.GetScheduledExpenseByID(uint(id))
	if err != nil {
		http.Error(w, "Failed to find scheduled expense", http.StatusNotFound)
		return
	}

	if err := h.expenseRepo.DeleteScheduledExpense(scheduledExpense); err != nil {
		http.Error(w, "Failed to delete scheduled expense", http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Trigger", "reloadScheduledExpenses")
	w.WriteHeader(http.StatusOK)
}

func (h *ExpenseHandler) doUpdateExpense(
	w http.ResponseWriter,
	r *http.Request,
	expense *models.Expense,
) *models.Expense {
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
		return nil
	}

	return expense
}

func (h *ExpenseHandler) UpdateExpense(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	expense, err := h.expenseRepo.GetByID(uint(id), "ExpenseSplits")

	h.doUpdateExpense(w, r, expense)

	if r.FormValue("IsScheduled") == "true" {
		h.doCreateScheduledExpense(w, r)
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

func (h *ExpenseHandler) doCreateScheduledExpense(
	w http.ResponseWriter,
	r *http.Request,
) (*models.ScheduledExpense, error) {
	logger.Debug.Println("Creating scheduled expense...")

	templateExpense := h.doCreateExpense(w, r, true)

	recurrenceType := r.FormValue("RecurrenceType")
	recurrenceInterval, _ := strconv.Atoi(r.FormValue("RecurrenceInterval"))

	startDate, err := helpers.ParseDate(r.FormValue("StartDate"))
	if err != nil {
		logger.Error.Println("Failed to parse start date:", err)
	}

	endDateStr := r.FormValue("EndDate")
	var endDate *time.Time
	if endDateStr == "" {
		endDate = nil
	} else {
		endDate, err = helpers.ParseDate(endDateStr)
		if err != nil {
			logger.Error.Println("Failed to parse end date:", err)
			return nil, err
		}
	}

	if err != nil {
		logger.Error.Println("Failed to parse date:", err)
		return nil, err
	}

	scheduledExpense := models.ScheduledExpense{
		TemplateExpenseID:  templateExpense.ID,
		RecurrenceType:     models.RecurrenceTypes(recurrenceType),
		RecurrenceInterval: recurrenceInterval,
		StartDate:          *startDate,
		EndDate:            endDate,
	}

	nextDueDate := jobs.CalculateNextDueDate(&scheduledExpense)

	if err := h.expenseRepo.CreateScheduledExpense(&scheduledExpense); err != nil {
		http.Error(w, "Failed to save scheduled expense", http.StatusInternalServerError)
		return nil, err
	}

	logger.Info.Println("Created scheduled expense: ", scheduledExpense.TemplateExpense.Title)
	logger.Info.Println("Next due date: ", nextDueDate)

	return &scheduledExpense, nil
}

func (h *ExpenseHandler) GetScheduledExpenses(w http.ResponseWriter, r *http.Request) {
	scheduledExpenses, err := h.expenseRepo.GetAllScheduledExpenses()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	partials.ScheduledExpensesTable(scheduledExpenses).Render(context.Background(), w)
}
