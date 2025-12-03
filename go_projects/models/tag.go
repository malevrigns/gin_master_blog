package models

import (
	"time"
)

// Tag 标签模型
type Tag struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar(100);uniqueIndex;not null"`
	Slug      string    `json:"slug" gorm:"type:varchar(100);uniqueIndex;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	
	Articles []Article `json:"articles" gorm:"many2many:article_tags;"`
}
