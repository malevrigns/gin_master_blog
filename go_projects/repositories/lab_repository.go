package repositories

import (
	"blog-system/database"
	"blog-system/models"
	"gorm.io/gorm"
)

type LabRepository interface {
	FindAll() ([]models.Lab, error)
	FindByID(id uint) (*models.Lab, error)
	FindBySlug(slug string) (*models.Lab, error)
	Create(lab *models.Lab) error
	Update(lab *models.Lab) error
	Delete(lab *models.Lab) error
}

type labRepository struct {
	db *gorm.DB
}

func NewLabRepository() LabRepository {
	return &labRepository{db: database.DB}
}

func (r *labRepository) FindAll() ([]models.Lab, error) {
	var labs []models.Lab
	err := r.db.Order("created_at DESC").Find(&labs).Error
	return labs, err
}

func (r *labRepository) FindByID(id uint) (*models.Lab, error) {
	var lab models.Lab
	err := r.db.First(&lab, id).Error
	return &lab, err
}

func (r *labRepository) FindBySlug(slug string) (*models.Lab, error) {
	var lab models.Lab
	err := r.db.Where("slug = ?", slug).First(&lab).Error
	return &lab, err
}

func (r *labRepository) Create(lab *models.Lab) error {
	return r.db.Create(lab).Error
}

func (r *labRepository) Update(lab *models.Lab) error {
	return r.db.Save(lab).Error
}

func (r *labRepository) Delete(lab *models.Lab) error {
	return r.db.Delete(lab).Error
}
