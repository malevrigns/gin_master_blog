package models

import (
	"time"

	"gorm.io/datatypes"
)

// Lab 实验室/专题模块
type Lab struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	Title         string         `json:"title" gorm:"type:varchar(255);not null"`
	Slug          string         `json:"slug" gorm:"type:varchar(100);uniqueIndex;not null"`
	Subtitle      string         `json:"subtitle" gorm:"type:varchar(255)"`
	Badge         string         `json:"badge" gorm:"type:varchar(50)"`
	BadgeColor    string         `json:"badge_color" gorm:"type:varchar(50)"`
	Description   string         `json:"description" gorm:"type:text"`
	Focus         string         `json:"focus" gorm:"type:varchar(255)"`
	HeroImage     string         `json:"hero_image" gorm:"type:varchar(500)"`
	Content       string         `json:"content" gorm:"type:longtext"`
	Highlights    datatypes.JSON `json:"highlights" gorm:"type:json"`
	ResourceLinks datatypes.JSON `json:"resource_links" gorm:"type:json"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

// LabResource 外链资源信息
type LabResource struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
	URL   string `json:"url"`
	Icon  string `json:"icon"`
}

// LabHighlight 模块内的主题/跳转
type LabHighlight struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Tag         string `json:"tag"`
	Category    string `json:"category"`
	Link        string `json:"link"`
}
