package repositories

import (
	"blog-system/database"
	"blog-system/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAll() ([]models.Category, error)
	FindByID(id uint) (*models.Category, error)
	Create(category *models.Category) error
	Update(category *models.Category) error
	Delete(category *models.Category) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{db: database.DB}
}

func (r *categoryRepository) FindAll() ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) FindByID(id uint) (*models.Category, error) {
	var category models.Category
	err := r.db.First(&category, id).Error
	return &category, err
}

func (r *categoryRepository) Create(category *models.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) Update(category *models.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) Delete(category *models.Category) error {
	return r.db.Delete(category).Error
}
