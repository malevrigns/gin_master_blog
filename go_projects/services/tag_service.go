package services

import (
	"blog-system/models"
	"blog-system/repositories"
	"github.com/gosimple/slug"
)

type TagService interface {
	GetTags() ([]models.Tag, error)
	GetTag(id uint) (*models.Tag, error)
	CreateTag(input *models.Tag) (*models.Tag, error)
	UpdateTag(id uint, input *models.Tag) (*models.Tag, error)
	DeleteTag(id uint) error
}

type tagService struct {
	repo repositories.TagRepository
}

func NewTagService() TagService {
	return &tagService{repo: repositories.NewTagRepository()}
}

func (s *tagService) GetTags() ([]models.Tag, error) {
	return s.repo.FindAll()
}

func (s *tagService) GetTag(id uint) (*models.Tag, error) {
	return s.repo.FindByID(id)
}

func (s *tagService) CreateTag(input *models.Tag) (*models.Tag, error) {
	input.Slug = slug.Make(input.Name)
	err := s.repo.Create(input)
	return input, err
}

func (s *tagService) UpdateTag(id uint, input *models.Tag) (*models.Tag, error) {
	tag, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	tag.Name = input.Name
	tag.Slug = slug.Make(input.Name)

	err = s.repo.Update(tag)
	return tag, err
}

func (s *tagService) DeleteTag(id uint) error {
	tag, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(tag)
}
