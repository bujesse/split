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
	GetAll() ([]models.Expense, error)
	DeleteExpense(expense *models.Expense) error
}

type expenseRepository struct {
	db *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) ExpenseRepository {
	return &expenseRepository{db}
}

func (r *expenseRepository) GetAll() ([]models.Expense, error) {
	var expenses []models.Expense
	result := r.db.Preload(clause.Associations).
		Preload("ExpenseSplits.User").
		Order("paid_date desc").
		Find(&expenses)
	if result.Error != nil {
		return nil, result.Error
	}
	return expenses, nil
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
			Where("paid_date < ?", earliestZeroDate).
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
			Where("paid_date > (?)", subqueryLatestZero).
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
		Where("paid_date > (?) AND paid_date < (?)", subqueryPreviousZero, subqueryLatestZero).
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

func (r *expenseRepository) DeleteExpense(expense *models.Expense) error {
	return r.db.Select("ExpenseSplits").Delete(&expense).Error
}
