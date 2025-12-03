package services

import (
	"blog-system/models"
	"blog-system/repositories"
	"time"

	"github.com/gosimple/slug"
	"strconv"
)

type ArticleService interface {
	GetArticles(page, pageSize int, filters map[string]interface{}) ([]models.Article, int64, int, int, error)
	GetArticle(id string) (*models.Article, error)
	GetArticleBySlug(slug string) (*models.Article, error)
	CreateArticle(input *models.Article, tagIDs []uint) (*models.Article, error)
	UpdateArticle(id string, input *models.Article, tagIDs []uint) (*models.Article, error)
	DeleteArticle(id string) error
	IncrementViews(id string) error
	LikeArticle(id string) (int, error)
}

type articleService struct {
	articleRepo repositories.ArticleRepository
	tagRepo     repositories.TagRepository
}

func NewArticleService() ArticleService {
	return &articleService{
		articleRepo: repositories.NewArticleRepository(),
		tagRepo:     repositories.NewTagRepository(),
	}
}

func (s *articleService) GetArticles(page, pageSize int, filters map[string]interface{}) ([]models.Article, int64, int, int, error) {
	articles, total, err := s.articleRepo.FindAll(page, pageSize, filters)
	return articles, total, page, pageSize, err
}

func (s *articleService) GetArticle(id string) (*models.Article, error) {
	return s.articleRepo.FindByID(id)
}

func (s *articleService) GetArticleBySlug(slug string) (*models.Article, error) {
	return s.articleRepo.FindBySlug(slug)
}

func (s *articleService) CreateArticle(input *models.Article, tagIDs []uint) (*models.Article, error) {
	if input.Status == "" {
		input.Status = "draft"
	}

	articleSlug := slug.Make(input.Title)
	count, _ := s.articleRepo.CountBySlug(articleSlug)
	if count > 0 {
		articleSlug = articleSlug + "-" + strconv.FormatInt(time.Now().Unix(), 10)
	}
	input.Slug = articleSlug

	if input.Status == "published" {
		now := time.Now()
		input.PublishedAt = &now
	}

	// Handle Tags
	if len(tagIDs) > 0 {
		tags, _ := s.tagRepo.FindByIds(tagIDs)
		input.Tags = tags
	}

	err := s.articleRepo.Create(input)
	return input, err
}

func (s *articleService) UpdateArticle(id string, input *models.Article, tagIDs []uint) (*models.Article, error) {
	article, err := s.articleRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if input.Title != "" {
		article.Title = input.Title
		article.Slug = slug.Make(input.Title)
	}
	if input.Content != "" {
		article.Content = input.Content
	}
	if input.Excerpt != "" {
		article.Excerpt = input.Excerpt
	}
	if input.CoverImage != "" {
		article.CoverImage = input.CoverImage
	}
	if input.CategoryID != 0 {
		article.CategoryID = input.CategoryID
	}
	if input.Status != "" {
		article.Status = input.Status
		if input.Status == "published" && article.PublishedAt == nil {
			now := time.Now()
			article.PublishedAt = &now
		}
	}
	article.IsTop = input.IsTop

	// Update Tags
	if tagIDs != nil {
		tags, _ := s.tagRepo.FindByIds(tagIDs)
		article.Tags = tags
	}

	err = s.articleRepo.Update(article)
	return article, err
}

func (s *articleService) DeleteArticle(id string) error {
	article, err := s.articleRepo.FindByID(id)
	if err != nil {
		return err
	}
	return s.articleRepo.Delete(article)
}

func (s *articleService) IncrementViews(id string) error {
	article, err := s.articleRepo.FindByID(id)
	if err != nil {
		return err
	}
	article.Views++
	return s.articleRepo.Update(article)
}

func (s *articleService) LikeArticle(id string) (int, error) {
	article, err := s.articleRepo.FindByID(id)
	if err != nil {
		return 0, err
	}
	article.Likes++
	err = s.articleRepo.Update(article)
	return article.Likes, err
}
