package repositories

import (
	"split/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CurrencyRepository interface {
	Create(currency *models.Currency) error
	GetByID(id uint) (*models.Currency, error)
	Update(currency *models.Currency) error
	GetAll() ([]models.Currency, error)
	Delete(id uint) error
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

func (r *currencyRepository) GetByID(id uint) (*models.Currency, error) {
	var currency models.Currency
	result := r.db.Preload(clause.Associations).First(&currency, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &currency, nil
}

func (r *currencyRepository) Update(currency *models.Currency) error {
	return r.db.Save(currency).Error
}

func (r *currencyRepository) Delete(id uint) error {
	return r.db.Delete(&models.Currency{}, id).Error
}
