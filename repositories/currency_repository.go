package repositories

import (
	"split/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CurrencyRepository interface {
	Create(currency *models.Currency) error
	GetByCode(code string) (*models.Currency, error)
	Update(currency *models.Currency) error
	GetAll() ([]models.Currency, error)
	Delete(code string) error
}

type currencyRepository struct {
	db *gorm.DB
}

func NewCurrencyRepository(db *gorm.DB) CurrencyRepository {
	return &currencyRepository{db}
}

func (r *currencyRepository) GetAll() ([]models.Currency, error) {
	var currencies []models.Currency
	result := r.db.Preload(clause.Associations).Order("Name asc").Find(&currencies)
	if result.Error != nil {
		return nil, result.Error
	}
	return currencies, nil
}

func (r *currencyRepository) Create(currency *models.Currency) error {
	return r.db.Create(currency).Error
}

func (r *currencyRepository) GetByCode(code string) (*models.Currency, error) {
	var currency models.Currency
	result := r.db.Preload(clause.Associations).First(&currency, code)
	if result.Error != nil {
		return nil, result.Error
	}
	return &currency, nil
}

func (r *currencyRepository) Update(currency *models.Currency) error {
	return r.db.Save(currency).Error
}

func (r *currencyRepository) Delete(code string) error {
	return r.db.Delete(&models.Currency{}, code).Error
}
