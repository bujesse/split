package repositories

import (
	"split/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ExpenseRepository interface {
	CreateExpense(expense *models.Expense) error
	GetByID(id uint, preloads ...string) (*models.Expense, error)
	UpdateExpense(expense *models.Expense) error
	GetAll() ([]models.Expense, error)
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
		Order("created_at desc").
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
