package repositories

import (
	"split/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CategoryRepository interface {
	Create(category *models.Category) error
	GetByID(id uint) (*models.Category, error)
	Update(category *models.Category) error
	GetAll() ([]models.Category, error)
	Delete(id uint) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) GetAll() ([]models.Category, error) {
	var categories []models.Category
	result := r.db.Preload(clause.Associations).Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

func (r *categoryRepository) Create(category *models.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) GetByID(id uint) (*models.Category, error) {
	var category models.Category
	result := r.db.Preload(clause.Associations).First(&category, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}

func (r *categoryRepository) Update(category *models.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) Delete(id uint) error {
	return r.db.Delete(&models.Category{}, id).Error
}
