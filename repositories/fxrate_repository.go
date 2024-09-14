package repositories

import (
	"split/models"

	"gorm.io/gorm"
)

type FxRateRepository interface {
	Create(fxRate *models.FxRate) error
	Update(fxRate *models.FxRate) error
	Delete(id uint) error
}

type fxRateRepository struct {
	db *gorm.DB
}

func NewFxRateRepository(db *gorm.DB) FxRateRepository {
	return &fxRateRepository{db}
}

func (r *fxRateRepository) Create(fxRate *models.FxRate) error {
	return r.db.Create(fxRate).Error
}

func (r *fxRateRepository) Update(fxRate *models.FxRate) error {
	return r.db.Save(fxRate).Error
}

func (r *fxRateRepository) Delete(id uint) error {
	return r.db.Delete(&models.FxRate{}, id).Error
}
