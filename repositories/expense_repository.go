package repositories

import (
	"split/models"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ExpenseRepository interface {
	CreateExpense(expense *models.Expense) error
	GetByID(id uint, preloads ...string) (*models.Expense, error)
	GetExpensesWithFxRate() ([]ExpenseWithFxRate, error)
	GetExpensesSinceLastSettlement() ([]ExpenseWithFxRate, error)
	GetExpensesBetweenZeros(offset int) ([]ExpenseWithFxRate, error)
	UpdateExpense(expense *models.Expense) error
	DeleteExpense(expense *models.Expense) error

	GetScheduledExpenseByID(id uint) (*models.ScheduledExpense, error)
	GetAllScheduledExpenses() ([]models.ScheduledExpense, error)
	GetDueScheduledExpenses() ([]models.ScheduledExpense, error)
	CreateScheduledExpense(scheduledExpense *models.ScheduledExpense) error
	UpdateScheduledExpense(scheduledExpense *models.ScheduledExpense) error
	DeleteScheduledExpense(scheduledExpense *models.ScheduledExpense) error
}

type expenseRepository struct {
	db *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) ExpenseRepository {
	return &expenseRepository{db}
}

func (r *expenseRepository) GetScheduledExpenseByID(
	id uint,
) (*models.ScheduledExpense, error) {
	var scheduledExpense models.ScheduledExpense
	result := r.db.
		Preload("TemplateExpense.PaidBy").
		Preload("TemplateExpense.ExpenseSplits").
		Preload("TemplateExpense.Category").
		Preload("TemplateExpense.ExpenseSplits.User").
		Preload(clause.Associations).
		First(&scheduledExpense, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &scheduledExpense, nil
}

func (r *expenseRepository) GetAllScheduledExpenses() ([]models.ScheduledExpense, error) {
	var scheduledExpenses []models.ScheduledExpense
	result := r.db.
		Preload("TemplateExpense.PaidBy").
		Preload("TemplateExpense.ExpenseSplits").
		Preload("TemplateExpense.ExpenseSplits.User").
		Preload(clause.Associations).
		Find(&scheduledExpenses)
	if result.Error != nil {
		return nil, result.Error
	}
	return scheduledExpenses, nil
}

// Fetch all scheduled expenses where the NextDueDate is today or in the past
func (r *expenseRepository) GetDueScheduledExpenses() ([]models.ScheduledExpense, error) {
	var scheduledExpenses []models.ScheduledExpense
	result := r.db.Where("next_due_date <= ?", time.Now()).
		Preload("TemplateExpense.PaidBy").
		Preload("TemplateExpense.ExpenseSplits").
		Preload("TemplateExpense.ExpenseSplits.User").
		Preload(clause.Associations).
		Find(&scheduledExpenses)
	if result.Error != nil {
		return nil, result.Error
	}
	return scheduledExpenses, nil
}

type ExpenseWithFxRate struct {
	models.Expense
	FxRate float64 `json:"fx_rate"`
}

func (r *expenseRepository) GetExpensesWithFxRate() ([]ExpenseWithFxRate, error) {
	var expenses []ExpenseWithFxRate

	subQuery := r.db.Table("fx_rates").
		Select("rate").
		Where("to_currency_code = expenses.currency_code AND from_currency_code = 'USD' AND DATE(date) = DATE(expenses.paid_date)").
		Order("date DESC").
		Limit(1)

	result := r.db.Table("expenses").
		Select("expenses.*, (?) AS fx_rate", subQuery).
		Where("is_template = 0").
		Preload("ExpenseSplits.User").
		Preload(clause.Associations).
		Order("expenses.paid_date DESC").
		Find(&expenses)

	if result.Error != nil {
		return nil, result.Error
	}

	return expenses, nil
}

func (r *expenseRepository) GetExpensesSinceLastSettlement() ([]ExpenseWithFxRate, error) {
	var expenses []ExpenseWithFxRate

	latestZeroSettlementSubQuery := r.db.Table("settlements").
		Select("settlement_date").
		Where("settled_to_zero = ?", true).
		Order("settlement_date DESC").
		Limit(1)

	var settlementCount int64
	r.db.Table("settlements").Where("settled_to_zero = ?", true).Count(&settlementCount)

	fxRateSubQuery := r.db.Table("fx_rates").
		Select("rate").
		Where("to_currency_code = expenses.currency_code AND from_currency_code = 'USD' AND DATE(date) = DATE(expenses.paid_date)").
		Order("date DESC").
		Limit(1)

	query := r.db.Table("expenses").
		Select("expenses.*, (?) AS fx_rate", fxRateSubQuery).
		Where("is_template = 0").
		Preload(clause.Associations).
		Preload("ExpenseSplits.User").
		Order("paid_date DESC")

	if settlementCount > 0 {
		query = query.Where("paid_date > (?)", latestZeroSettlementSubQuery)
	}

	result := query.Find(&expenses)

	if result.Error != nil {
		return nil, result.Error
	}
	return expenses, nil
}

// Either returns all expenses since the last zero-settled settlement,
// or all expenses between the nth and n+1 zero-settled settlements
func (r *expenseRepository) GetExpensesBetweenZeros(offset int) ([]ExpenseWithFxRate, error) {
	var expenses []ExpenseWithFxRate

	// Subquery to get the nth latest settlement date where SettledToZero is true
	subqueryNthZero := func(n int) *gorm.DB {
		return r.db.Table("settlements").
			Select("settlement_date").
			Where("settled_to_zero = ?", true).
			Order("settlement_date desc").
			Offset(n).
			Limit(1)
	}

	// Subquery for FX rate for each expense
	fxRateSubQuery := r.db.Table("fx_rates").
		Select("rate").
		Where("to_currency_code = expenses.currency_code AND from_currency_code = 'USD' AND DATE(date) = DATE(expenses.paid_date)").
		Order("date DESC").
		Limit(1)

	var totalZeroSettlements int64
	r.db.Model(&models.Settlement{}).Where("settled_to_zero = ?", true).Count(&totalZeroSettlements)
	if totalZeroSettlements > 0 && int64(offset) >= totalZeroSettlements {
		// Get the earliest zero-settled settlement date
		var earliestZeroDate time.Time
		if err := subqueryNthZero(int(totalZeroSettlements - 1)).Scan(&earliestZeroDate).Error; err != nil {
			return nil, err
		}

		// Return all expenses up until the earliest zero-settled settlement
		result := r.db.Table("expenses").
			Select("expenses.*, (?) AS fx_rate", fxRateSubQuery).
			Preload(clause.Associations).
			Preload("ExpenseSplits.User").
			Where("is_template = 0 AND paid_date < ?", earliestZeroDate).
			Order("paid_date desc").
			Find(&expenses)

		if result.Error != nil {
			return nil, result.Error
		}
		return expenses, nil
	}

	var count int64
	subqueryPreviousZero := subqueryNthZero(offset)

	r.db.Table("(?) as prev_zero", subqueryPreviousZero).Count(&count)

	if count == 0 {
		// Only one zero-settled settlement exists, return expenses after the latest zero
		subqueryLatestZero := subqueryNthZero(offset)

		result := r.db.Table("expenses").
			Select("expenses.*, (?) AS fx_rate", fxRateSubQuery).
			Preload(clause.Associations).
			Preload("ExpenseSplits.User").
			Where("is_template = 0 AND paid_date > (?)", subqueryLatestZero).
			Order("paid_date DESC").
			Find(&expenses)

		if result.Error != nil {
			return nil, result.Error
		}
		return expenses, nil
	}

	// If more than one zero-settled settlement exists, return expenses between the two zeros
	subqueryLatestZero := subqueryNthZero(offset - 1)

	result := r.db.Table("expenses").
		Select("expenses.*, (?) AS fx_rate", fxRateSubQuery).
		Preload(clause.Associations).
		Preload("ExpenseSplits.User").
		Where("is_template = 0 AND paid_date > (?) AND paid_date < (?)", subqueryPreviousZero, subqueryLatestZero).
		Order("paid_date DESC").
		Find(&expenses)

	if result.Error != nil {
		return nil, result.Error
	}

	return expenses, nil
}

func (r *expenseRepository) CreateExpense(expense *models.Expense) error {
	return r.db.Create(expense).Error
}

func (r *expenseRepository) GetByID(id uint, preloads ...string) (*models.Expense, error) {
	var expense models.Expense
	query := r.db
	for _, preload := range preloads {
		query = query.Preload(preload)
	}
	result := query.First(&expense, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &expense, nil
}

func (r *expenseRepository) UpdateExpense(expense *models.Expense) error {
	return r.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(expense).Error
}

func (r *expenseRepository) CreateScheduledExpense(
	scheduledExpense *models.ScheduledExpense,
) error {
	return r.db.Create(scheduledExpense).Error
}

func (r *expenseRepository) UpdateScheduledExpense(
	scheduledExpense *models.ScheduledExpense,
) error {
	return r.db.Save(scheduledExpense).Error
}

func (r *expenseRepository) DeleteScheduledExpense(
	scheduledExpense *models.ScheduledExpense,
) error {
	err := r.db.Delete(scheduledExpense.TemplateExpense).Error
	err = r.db.Delete(scheduledExpense).Error
	return err
}

func (r *expenseRepository) DeleteExpense(expense *models.Expense) error {
	return r.db.Select("ExpenseSplits").Delete(&expense).Error
}
