package repositories

import (
	"split/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SettlementRepository interface {
	Create(settlement *models.Settlement) error
	GetByID(id uint) (*models.Settlement, error)
	// GetByName(name string) (*models.Settlement, error)
	Update(settlement *models.Settlement) error
	GetAll() ([]models.Settlement, error)
	GetAllSinceLastSettlement() ([]models.Settlement, error)
	Delete(id uint) error
}

type settlementRepository struct {
	db *gorm.DB
}

func NewSettlementRepository(db *gorm.DB) SettlementRepository {
	return &settlementRepository{db}
}

func (r *settlementRepository) GetAll() ([]models.Settlement, error) {
	var settlements []models.Settlement
	result := r.db.Preload(clause.Associations).Order("settlement_date desc").Find(&settlements)
	if result.Error != nil {
		return nil, result.Error
	}
	return settlements, nil
}

func (r *settlementRepository) Create(settlement *models.Settlement) error {
	return r.db.Create(settlement).Error
}

func (r *settlementRepository) GetByID(id uint) (*models.Settlement, error) {
	var settlement models.Settlement
	result := r.db.Preload(clause.Associations).First(&settlement, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &settlement, nil
}

func (r *settlementRepository) GetAllSinceLastSettlement() ([]models.Settlement, error) {
	var settlements []models.Settlement

	subquery := r.db.Model(&models.Settlement{}).
		Select("settlement_date").
		Where("settled_to_zero = ?", true).
		Order("settlement_date desc").
		Limit(1)

	result := r.db.Preload(clause.Associations).
		Where("settlement_date > (?)", subquery).
		Order("settlement_date desc").
		Find(&settlements)

	if result.Error != nil {
		return nil, result.Error
	}

	return settlements, nil
}

// func (r *settlementRepository) GetByName(id string) (*models.Settlement, error) {
// 	var settlement models.Settlement
// 	result := r.db.Preload(clause.Associations).First(&settlement)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return &settlement, nil
// }

func (r *settlementRepository) Update(settlement *models.Settlement) error {
	return r.db.Save(settlement).Error
}

func (r *settlementRepository) Delete(id uint) error {
	return r.db.Delete(&models.Settlement{}, id).Error
}
