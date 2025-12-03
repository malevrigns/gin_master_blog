package repositories

import (
	"blog-system/database"
	"blog-system/models"
	"gorm.io/gorm"
)

type MusicRepository interface {
	FindAll() ([]models.Music, error)
	FindByID(id uint) (*models.Music, error)
	Create(music *models.Music) error
	Update(music *models.Music) error
	Delete(music *models.Music) error
	
	FindPlaylists() ([]models.Playlist, error)
	FindPlaylistByID(id uint) (*models.Playlist, error)
}

type musicRepository struct {
	db *gorm.DB
}

func NewMusicRepository() MusicRepository {
	return &musicRepository{db: database.DB}
}

func (r *musicRepository) FindAll() ([]models.Music, error) {
	var music []models.Music
	err := r.db.Order("created_at DESC").Find(&music).Error
	return music, err
}

func (r *musicRepository) FindByID(id uint) (*models.Music, error) {
	var music models.Music
	err := r.db.First(&music, id).Error
	return &music, err
}

func (r *musicRepository) Create(music *models.Music) error {
	return r.db.Create(music).Error
}

func (r *musicRepository) Update(music *models.Music) error {
	return r.db.Save(music).Error
}

func (r *musicRepository) Delete(music *models.Music) error {
	return r.db.Delete(music).Error
}

func (r *musicRepository) FindPlaylists() ([]models.Playlist, error) {
	var playlists []models.Playlist
	err := r.db.Preload("Musics").Find(&playlists).Error
	return playlists, err
}

func (r *musicRepository) FindPlaylistByID(id uint) (*models.Playlist, error) {
	var playlist models.Playlist
	err := r.db.Preload("Musics").First(&playlist, id).Error
	return &playlist, err
}
