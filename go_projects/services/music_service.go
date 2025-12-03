package services

import (
	"blog-system/models"
	"blog-system/repositories"
)

type MusicService interface {
	GetMusicList() ([]models.Music, error)
	GetMusic(id uint) (*models.Music, error)
	CreateMusic(input *models.Music) (*models.Music, error)
	UpdateMusic(id uint, input *models.Music) (*models.Music, error)
	DeleteMusic(id uint) error
	
	GetPlaylists() ([]models.Playlist, error)
	GetPlaylist(id uint) (*models.Playlist, error)
}

type musicService struct {
	repo repositories.MusicRepository
}

func NewMusicService() MusicService {
	return &musicService{repo: repositories.NewMusicRepository()}
}

func (s *musicService) GetMusicList() ([]models.Music, error) {
	return s.repo.FindAll()
}

func (s *musicService) GetMusic(id uint) (*models.Music, error) {
	return s.repo.FindByID(id)
}

func (s *musicService) CreateMusic(input *models.Music) (*models.Music, error) {
	err := s.repo.Create(input)
	return input, err
}

func (s *musicService) UpdateMusic(id uint, input *models.Music) (*models.Music, error) {
	music, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	music.Title = input.Title
	music.Artist = input.Artist
	music.Cover = input.Cover
	music.URL = input.URL
	music.Duration = input.Duration
	music.Lrc = input.Lrc
	music.IsPublic = input.IsPublic

	err = s.repo.Update(music)
	return music, err
}

func (s *musicService) DeleteMusic(id uint) error {
	music, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(music)
}

func (s *musicService) GetPlaylists() ([]models.Playlist, error) {
	return s.repo.FindPlaylists()
}

func (s *musicService) GetPlaylist(id uint) (*models.Playlist, error) {
	return s.repo.FindPlaylistByID(id)
}
