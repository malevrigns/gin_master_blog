package repositories

import (
	"blog-system/database"
	"blog-system/models"
	"gorm.io/gorm"
)

type TagRepository interface {
	FindAll() ([]models.Tag, error)
	FindByID(id uint) (*models.Tag, error)
	FindByIds(ids []uint) ([]models.Tag, error)
	Create(tag *models.Tag) error
	Update(tag *models.Tag) error
	Delete(tag *models.Tag) error
}

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository() TagRepository {
	return &tagRepository{db: database.DB}
}

func (r *tagRepository) FindAll() ([]models.Tag, error) {
	var tags []models.Tag
	err := r.db.Find(&tags).Error
	return tags, err
}

func (r *tagRepository) FindByID(id uint) (*models.Tag, error) {
	var tag models.Tag
	err := r.db.First(&tag, id).Error
	return &tag, err
}

func (r *tagRepository) FindByIds(ids []uint) ([]models.Tag, error) {
	var tags []models.Tag
	err := r.db.Where("id IN ?", ids).Find(&tags).Error
	return tags, err
}

func (r *tagRepository) Create(tag *models.Tag) error {
	return r.db.Create(tag).Error
}

func (r *tagRepository) Update(tag *models.Tag) error {
	return r.db.Save(tag).Error
}

func (r *tagRepository) Delete(tag *models.Tag) error {
	return r.db.Delete(tag).Error
}
