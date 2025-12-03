package models

import (
	"time"
)

// Category 分类模型
type Category struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"type:varchar(100);uniqueIndex;not null"`
	Slug        string    `json:"slug" gorm:"type:varchar(100);uniqueIndex;not null"`
	Description string    `json:"description" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	
	Articles []Article `json:"articles" gorm:"foreignKey:CategoryID"`
}
