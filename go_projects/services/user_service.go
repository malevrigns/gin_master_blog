package services

import (
	"blog-system/models"
	"blog-system/repositories"
	"blog-system/utils"
	"errors"
)

type UserService interface {
	GetUser(id uint) (*models.User, error)
	Register(username, email, password string) (*models.User, error)
	Login(username, password string) (string, *models.User, error)
	UpdateProfile(id uint, input *models.User) (*models.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService() UserService {
	return &userService{repo: repositories.NewUserRepository()}
}

func (s *userService) GetUser(id uint) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) Register(username, email, password string) (*models.User, error) {
	// Check if user exists
	if _, err := s.repo.FindByUsername(username); err == nil {
		return nil, errors.New("username already exists")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: username,
		Email:    email,
		Password: hashedPassword,
		Role:     "user",
	}

	err = s.repo.Create(user)
	return user, err
}

func (s *userService) Login(username, password string) (string, *models.User, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", nil, errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

func (s *userService) UpdateProfile(id uint, input *models.User) (*models.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if input.Avatar != "" {
		user.Avatar = input.Avatar
	}
	if input.Bio != "" {
		user.Bio = input.Bio
	}

	err = s.repo.Update(user)
	return user, err
}
