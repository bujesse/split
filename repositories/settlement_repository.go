package repositories

import (
	"split/models"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SettlementRepository interface {
	Create(settlement *models.Settlement) error
	GetByID(id uint) (*models.Settlement, error)
	// GetByName(name string) (*models.Settlement, error)
	Update(settlement *models.Settlement) error
	GetAll() ([]models.Settlement, error)
	GetAllSinceLastZeroSettlement() ([]models.Settlement, error)
	GetSettlementsBetweenZeros(offset int) ([]models.Settlement, error)
	GetNumZeroSettlements() (int64, error)
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

func (r *settlementRepository) GetAllSinceLastZeroSettlement() ([]models.Settlement, error) {
	var settlements []models.Settlement

	totalZeroSettlements, _ := r.GetNumZeroSettlements()
	if totalZeroSettlements == 0 {
		return nil, nil
	}

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

// Return all settlements between the n-1 and n zero-settled settlements.
// If the offset is greater than the total number of zero-settled settlements,
// return all settlements up until the earliest zero-settled settlement
func (r *settlementRepository) GetSettlementsBetweenZeros(
	offset int,
) ([]models.Settlement, error) {
	var settlements []models.Settlement
	var newerZeroDate, olderZeroDate time.Time

	subqueryNthZero := func(n int) *gorm.DB {
		return r.db.Model(&models.Settlement{}).
			Select("settlement_date").
			Where("settled_to_zero = ?", true).
			Order("settlement_date desc").
			Offset(n).
			Limit(1)
	}

	totalZeroSettlements, _ := r.GetNumZeroSettlements()

	// If the offset is greater than or equal to the total number of zero-settled settlements,
	// return all settlements up until the earliest zero-settled settlement
	if int64(offset+1) > totalZeroSettlements {
		if err := subqueryNthZero(int(totalZeroSettlements - 1)).Scan(&olderZeroDate).Error; err != nil {
			return nil, err
		}

		result := r.db.Preload(clause.Associations).
			Where("settlement_date <= ?", olderZeroDate).
			Order("settlement_date desc").
			Find(&settlements)

		if result.Error != nil {
			return nil, result.Error
		}
		return settlements, nil
	}

	if err := subqueryNthZero(offset - 1).Scan(&newerZeroDate).Error; err != nil {
		return nil, err
	}

	if err := subqueryNthZero(offset).Scan(&olderZeroDate).Error; err != nil {
		return nil, err
	}

	result := r.db.Preload(clause.Associations).
		Where("settlement_date > ? AND settlement_date <= ?", olderZeroDate, newerZeroDate).
		Order("settlement_date desc").
		Find(&settlements)

	if result.Error != nil {
		return nil, result.Error
	}

	return settlements, nil
}

func (r *settlementRepository) GetNumZeroSettlements() (int64, error) {
	var totalZeroSettlements int64
	r.db.Model(&models.Settlement{}).Where("settled_to_zero = ?", true).Count(&totalZeroSettlements)
	return totalZeroSettlements, nil
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
