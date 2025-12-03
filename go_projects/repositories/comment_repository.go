package repositories

import (
	"blog-system/database"
	"blog-system/models"
	"gorm.io/gorm"
)

type CommentRepository interface {
	FindAll(page, pageSize int) ([]models.Comment, int64, error)
	FindPending(page, pageSize int) ([]models.Comment, int64, error)
	FindByID(id uint) (*models.Comment, error)
	Create(comment *models.Comment) error
	Update(comment *models.Comment) error
	Delete(comment *models.Comment) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository() CommentRepository {
	return &commentRepository{db: database.DB}
}

func (r *commentRepository) FindAll(page, pageSize int) ([]models.Comment, int64, error) {
	var comments []models.Comment
	var total int64
	
	query := r.db.Preload("Article")
	query.Model(&models.Comment{}).Count(&total)

	offset := (page - 1) * pageSize
	err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&comments).Error
	return comments, total, err
}

func (r *commentRepository) FindPending(page, pageSize int) ([]models.Comment, int64, error) {
	var comments []models.Comment
	var total int64
	
	query := r.db.Preload("Article").Where("status = ?", "pending")
	query.Model(&models.Comment{}).Count(&total)

	offset := (page - 1) * pageSize
	err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&comments).Error
	return comments, total, err
}

func (r *commentRepository) FindByID(id uint) (*models.Comment, error) {
	var comment models.Comment
	err := r.db.First(&comment, id).Error
	return &comment, err
}

func (r *commentRepository) Create(comment *models.Comment) error {
	return r.db.Create(comment).Error
}

func (r *commentRepository) Update(comment *models.Comment) error {
	return r.db.Save(comment).Error
}

func (r *commentRepository) Delete(comment *models.Comment) error {
	return r.db.Delete(comment).Error
}
