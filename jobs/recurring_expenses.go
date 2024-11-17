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
	scheduledExpenses, err := expenseRepo.GetDueScheduledExpenses()

	for _, scheduledExpense := range scheduledExpenses {
		templateExpense := scheduledExpense.TemplateExpense

		logger.Info.Println("Processing scheduled expense:", templateExpense.Title)

		newExpense := models.Expense{
			Title:              templateExpense.Title,
			Description:        templateExpense.Description,
			Amount:             templateExpense.Amount,
			Notes:              templateExpense.Notes,
			CurrencyCode:       templateExpense.CurrencyCode,
			CategoryID:         templateExpense.CategoryID,
			PaidByID:           templateExpense.PaidByID,
			PaidDate:           time.Now(),
			ScheduledExpenseID: &scheduledExpense.ID,
			CreatedByID:        templateExpense.CreatedByID,
		}

		for _, split := range templateExpense.ExpenseSplits {
			newSplit := models.ExpenseSplit{
				UserID:       split.UserID,
				SplitType:    models.SplitType(split.SplitType),
				SplitValue:   split.SplitValue,
				CurrencyCode: split.CurrencyCode,
			}
			newExpense.ExpenseSplits = append(newExpense.ExpenseSplits, newSplit)
		}

		err = expenseRepo.CreateExpense(&newExpense)
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
			newExpense.Title,
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
		scheduledExpense.NextDueDate = &scheduledExpense.StartDate
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
