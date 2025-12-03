package models

import (
	"time"
)

// Link 友情链接模型
type Link struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	URL       string    `json:"url" gorm:"type:varchar(500);not null"`
	Logo      string    `json:"logo" gorm:"type:varchar(500)"`
	Desc      string    `json:"desc" gorm:"type:text"`
	IsVisible bool      `json:"is_visible" gorm:"default:true"`
	Sort      int       `json:"sort" gorm:"default:0"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SiteConfig 站点配置模型
type SiteConfig struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Key       string    `json:"key" gorm:"type:varchar(255);uniqueIndex;not null"`
	Value     string    `json:"value" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
