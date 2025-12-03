package models

import (
	"time"
)

// Music 音乐模型
type Music struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"type:varchar(255);not null"`
	Artist      string    `json:"artist" gorm:"type:varchar(255)"`
	Cover       string    `json:"cover" gorm:"type:varchar(500)"`
	URL         string    `json:"url" gorm:"type:varchar(500);not null"`
	Duration    int       `json:"duration"` // 秒
	Lrc         string    `json:"lrc" gorm:"type:varchar(500)"`      // 歌词文件路径
	IsPublic    bool      `json:"is_public" gorm:"default:true"`
	PlayCount   int       `json:"play_count" gorm:"default:0"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Playlist 播放列表模型
type Playlist struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	Cover     string    `json:"cover" gorm:"type:varchar(500)"`
	IsPublic  bool      `json:"is_public" gorm:"default:true"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	
	Musics []Music `json:"musics" gorm:"many2many:playlist_musics;"`
}
