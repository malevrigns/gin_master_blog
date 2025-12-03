package services

import (
	"blog-system/models"
	"blog-system/repositories"
	"github.com/gosimple/slug"
)

type CategoryService interface {
	GetCategories() ([]models.Category, error)
	GetCategory(id uint) (*models.Category, error)
	CreateCategory(input *models.Category) (*models.Category, error)
	UpdateCategory(id uint, input *models.Category) (*models.Category, error)
	DeleteCategory(id uint) error
}

type categoryService struct {
	repo repositories.CategoryRepository
}

func NewCategoryService() CategoryService {
	return &categoryService{repo: repositories.NewCategoryRepository()}
}

func (s *categoryService) GetCategories() ([]models.Category, error) {
	return s.repo.FindAll()
}

func (s *categoryService) GetCategory(id uint) (*models.Category, error) {
	return s.repo.FindByID(id)
}

func (s *categoryService) CreateCategory(input *models.Category) (*models.Category, error) {
	input.Slug = slug.Make(input.Name)
	err := s.repo.Create(input)
	return input, err
}

func (s *categoryService) UpdateCategory(id uint, input *models.Category) (*models.Category, error) {
	category, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	category.Name = input.Name
	category.Slug = slug.Make(input.Name)
	category.Description = input.Description

	err = s.repo.Update(category)
	return category, err
}

func (s *categoryService) DeleteCategory(id uint) error {
	category, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(category)
}
