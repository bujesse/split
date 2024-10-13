package jobs

import (
	"fmt"
	"split/config/logger"
	"split/models"
	"split/repositories"
	"time"

	"gorm.io/gorm"
)

func ProcessRecurringExpenses(expenseRepo repositories.ExpenseRepository) {
	logger.Info.Println("Checking for recurring expenses...")
	scheduledExpenses, err := expenseRepo.GetScheduledExpenses()

	for _, scheduledExpense := range scheduledExpenses {
		logger.Info.Println("Processing scheduled expense:", scheduledExpense.Title)
		expense := models.Expense{
			Title:              scheduledExpense.Title,
			Description:        scheduledExpense.Description,
			Amount:             scheduledExpense.Amount,
			Notes:              scheduledExpense.Notes,
			CurrencyCode:       scheduledExpense.CurrencyCode,
			CategoryID:         scheduledExpense.CategoryID,
			PaidByID:           scheduledExpense.PaidByID,
			PaidDate:           time.Now(),
			ScheduledExpenseID: &scheduledExpense.ID,
			CreatedByID:        scheduledExpense.CreatedByID,
		}

		err = expenseRepo.CreateExpense(&expense)
		if err != nil {
			fmt.Println("Error creating expense:", err)
			continue
		}

		nextDueDate := CalculateNextDueDate(&scheduledExpense)

		err = expenseRepo.UpdateScheduledExpense(&scheduledExpense)
		if err != nil {
			logger.Info.Println("Error updating scheduled expense:", err)
		}

		logger.Info.Printf(
			"Created recurring expense: %s, next due: %s\n",
			expense.Title,
			nextDueDate,
		)
	}
}

func fetchDueExpenses(db *gorm.DB) ([]models.Expense, error) {
	var expenses []models.Expense
	today := time.Now().Truncate(24 * time.Hour) // Truncate to remove time component
	err := db.Where("next_due = ?", today).Find(&expenses).Error
	return expenses, err
}

// Calulates AND sets the next due date for a scheduled expense
func CalculateNextDueDate(scheduledExpense *models.ScheduledExpense) *time.Time {
	if scheduledExpense.NextDueDate == nil {
		scheduledExpense.NextDueDate = &scheduledExpense.StartDate
	}

	if scheduledExpense.EndDate != nil && time.Now().After(*scheduledExpense.EndDate) {
		scheduledExpense.NextDueDate = nil
		return nil
	}

	if scheduledExpense.StartDate.After(time.Now()) {
		return scheduledExpense.NextDueDate
	}

	var nextDueDate time.Time
	switch scheduledExpense.RecurrenceType {
	case models.Daily:
		nextDueDate = scheduledExpense.NextDueDate.AddDate(
			0,
			0,
			scheduledExpense.RecurrenceInterval,
		)
	case models.Weekly:
		nextDueDate = scheduledExpense.NextDueDate.AddDate(
			0,
			0,
			7*scheduledExpense.RecurrenceInterval,
		)
	case models.Monthly:
		nextDueDate = scheduledExpense.NextDueDate.AddDate(
			0,
			scheduledExpense.RecurrenceInterval,
			0,
		)
	case models.Yearly:
		nextDueDate = scheduledExpense.NextDueDate.AddDate(
			scheduledExpense.RecurrenceInterval,
			0,
			0,
		)
	default:
		nextDueDate = *scheduledExpense.NextDueDate
	}

	scheduledExpense.NextDueDate = &nextDueDate
	return &nextDueDate
}
