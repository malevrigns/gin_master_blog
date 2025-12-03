package repositories

import (
	"blog-system/database"
	"blog-system/models"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	FindAll(page, pageSize int, filters map[string]interface{}) ([]models.Article, int64, error)
	FindByID(id string) (*models.Article, error)
	FindBySlug(slug string) (*models.Article, error)
	Create(article *models.Article) error
	Update(article *models.Article) error
	Delete(article *models.Article) error
	CountBySlug(slug string) (int64, error)
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository() ArticleRepository {
	return &articleRepository{db: database.DB}
}

func (r *articleRepository) FindAll(page, pageSize int, filters map[string]interface{}) ([]models.Article, int64, error) {
	var articles []models.Article
	var total int64
	
	query := r.db.Preload("Author").Preload("Category").Preload("Tags")

	if status, ok := filters["status"]; ok && status != "" {
		query = query.Where("status = ?", status)
	}

	if categoryID, ok := filters["category_id"]; ok && categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}

	if tagSlug, ok := filters["tag_slug"]; ok && tagSlug != "" {
		query = query.Joins("JOIN article_tags ON article_tags.article_id = articles.id").
			Joins("JOIN tags ON tags.id = article_tags.tag_id").
			Where("tags.slug = ?", tagSlug)
	}

	if search, ok := filters["search"]; ok && search != "" {
		query = query.Where("title LIKE ? OR content LIKE ?", "%"+search.(string)+"%", "%"+search.(string)+"%")
	}

	query.Model(&models.Article{}).Count(&total)

	offset := (page - 1) * pageSize
	err := query.Order("is_top DESC, published_at DESC").Offset(offset).Limit(pageSize).Find(&articles).Error
	return articles, total, err
}

func (r *articleRepository) FindByID(id string) (*models.Article, error) {
	var article models.Article
	err := r.db.Preload("Author").Preload("Category").Preload("Tags").
		Preload("Comments", "status = ?", "approved").Preload("Comments.Replies").
		First(&article, id).Error
	return &article, err
}

func (r *articleRepository) FindBySlug(slug string) (*models.Article, error) {
	var article models.Article
	err := r.db.Where("slug = ?", slug).First(&article).Error
	return &article, err
}

func (r *articleRepository) Create(article *models.Article) error {
	return r.db.Create(article).Error
}

func (r *articleRepository) Update(article *models.Article) error {
	return r.db.Save(article).Error
}

func (r *articleRepository) Delete(article *models.Article) error {
	return r.db.Delete(article).Error
}

func (r *articleRepository) CountBySlug(slug string) (int64, error) {
	var count int64
	err := r.db.Model(&models.Article{}).Where("slug = ?", slug).Count(&count).Error
	return count, err
}
