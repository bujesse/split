package repositories

import (
	"split/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FxRateRepository interface {
	Create(fxRate *models.FxRate) error
	GetByID(id uint) (*models.FxRate, error)
	// GetByName(name string) (*models.FxRate, error)
	Update(fxRate *models.FxRate) error
	GetAll() ([]models.FxRate, error)
	Delete(id uint) error
}

type fxRateRepository struct {
	db *gorm.DB
}

func NewFxRateRepository(db *gorm.DB) FxRateRepository {
	return &fxRateRepository{db}
}

func (r *fxRateRepository) GetAll() ([]models.FxRate, error) {
	var fxRates []models.FxRate
	result := r.db.Preload(clause.Associations).Order("fxRate_date desc").Find(&fxRates)
	if result.Error != nil {
		return nil, result.Error
	}
	return fxRates, nil
}

func (r *fxRateRepository) Create(fxRate *models.FxRate) error {
	return r.db.Create(fxRate).Error
}

func (r *fxRateRepository) GetByID(id uint) (*models.FxRate, error) {
	var fxRate models.FxRate
	result := r.db.Preload(clause.Associations).First(&fxRate, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &fxRate, nil
}

// func (r *fxRateRepository) GetByName(id string) (*models.FxRate, error) {
// 	var fxRate models.FxRate
// 	result := r.db.Preload(clause.Associations).First(&fxRate)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return &fxRate, nil
// }

func (r *fxRateRepository) Update(fxRate *models.FxRate) error {
	return r.db.Save(fxRate).Error
}

func (r *fxRateRepository) Delete(id uint) error {
	return r.db.Delete(&models.FxRate{}, id).Error
}
