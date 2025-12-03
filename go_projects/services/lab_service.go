package services

import (
	"blog-system/models"
	"blog-system/repositories"
	"github.com/gosimple/slug"
)

type LabService interface {
	GetLabs() ([]models.Lab, error)
	GetLab(id uint) (*models.Lab, error)
	GetLabBySlug(slug string) (*models.Lab, error)
	CreateLab(input *models.Lab) (*models.Lab, error)
	UpdateLab(id uint, input *models.Lab) (*models.Lab, error)
	DeleteLab(id uint) error
}

type labService struct {
	repo repositories.LabRepository
}

func NewLabService() LabService {
	return &labService{repo: repositories.NewLabRepository()}
}

func (s *labService) GetLabs() ([]models.Lab, error) {
	return s.repo.FindAll()
}

func (s *labService) GetLab(id uint) (*models.Lab, error) {
	return s.repo.FindByID(id)
}

func (s *labService) GetLabBySlug(slug string) (*models.Lab, error) {
	return s.repo.FindBySlug(slug)
}

func (s *labService) CreateLab(input *models.Lab) (*models.Lab, error) {
	input.Slug = slug.Make(input.Title)
	err := s.repo.Create(input)
	return input, err
}

func (s *labService) UpdateLab(id uint, input *models.Lab) (*models.Lab, error) {
	lab, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	lab.Title = input.Title
	lab.Slug = slug.Make(input.Title)
	lab.Subtitle = input.Subtitle
	lab.Badge = input.Badge
	lab.BadgeColor = input.BadgeColor
	lab.Description = input.Description
	lab.Focus = input.Focus
	lab.HeroImage = input.HeroImage
	lab.Content = input.Content
	lab.Highlights = input.Highlights
	lab.ResourceLinks = input.ResourceLinks

	err = s.repo.Update(lab)
	return lab, err
}

func (s *labService) DeleteLab(id uint) error {
	lab, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(lab)
}
