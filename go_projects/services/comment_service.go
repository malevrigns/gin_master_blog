package services

import (
	"blog-system/models"
	"blog-system/repositories"
)

type CommentService interface {
	GetComments(page, pageSize int) ([]models.Comment, int64, error)
	GetPendingComments(page, pageSize int) ([]models.Comment, int64, error)
	CreateComment(input *models.Comment) (*models.Comment, error)
	UpdateCommentStatus(id uint, status string) (*models.Comment, error)
	DeleteComment(id uint) error
}

type commentService struct {
	repo repositories.CommentRepository
}

func NewCommentService() CommentService {
	return &commentService{repo: repositories.NewCommentRepository()}
}

func (s *commentService) GetComments(page, pageSize int) ([]models.Comment, int64, error) {
	return s.repo.FindAll(page, pageSize)
}

func (s *commentService) GetPendingComments(page, pageSize int) ([]models.Comment, int64, error) {
	// Assuming Repository needs update or we filter here?
	// Repository FindAll doesn't support filters yet.
	// I need to update Repository first or add a new method to Repository.
	// Let's add FindPending to Repository.
	return s.repo.FindPending(page, pageSize)
}

func (s *commentService) CreateComment(input *models.Comment) (*models.Comment, error) {
	err := s.repo.Create(input)
	return input, err
}

func (s *commentService) UpdateCommentStatus(id uint, status string) (*models.Comment, error) {
	comment, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	comment.Status = status
	err = s.repo.Update(comment)
	return comment, err
}

func (s *commentService) DeleteComment(id uint) error {
	comment, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(comment)
}
