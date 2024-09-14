package repositories

import (
	"split/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ExpenseRepository interface {
	CreateExpense(expense *models.Expense) error
	GetByID(id uint, preloads ...string) (*models.Expense, error)
	GetExpensesWithFxRate() ([]ExpenseWithFxRate, error)
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
