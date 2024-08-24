package services

import (
	"split/models"
	"split/repositories"
)

type ExpenseService interface {
	CreateExpense(expense *models.Expense) error
	GetExpenseByID(id uint) (*models.Expense, error)
	UpdateExpense(expense *models.Expense) error
}

type expenseService struct {
	repo repositories.ExpenseRepository
}

func NewExpenseService(repo repositories.ExpenseRepository) ExpenseService {
	return &expenseService{repo}
}

func (s *expenseService) CreateExpense(expense *models.Expense) error {
	return s.repo.Create(expense)
}

func (s *expenseService) GetExpenseByID(id uint) (*models.Expense, error) {
	return s.repo.GetByID(id)
}

func (s *expenseService) UpdateExpense(expense *models.Expense) error {
	return s.repo.Update(expense)
}
