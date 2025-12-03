package models

import (
	"time"
)

// Comment 评论模型
type Comment struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Content   string    `json:"content" gorm:"type:text;not null"`
	Author    string    `json:"author" gorm:"type:varchar(100);not null"`
	Email     string    `json:"email" gorm:"type:varchar(255)"`
	Website   string    `json:"website" gorm:"type:varchar(500)"`
	IP        string    `json:"ip" gorm:"type:varchar(50)"`
	Status    string    `json:"status" gorm:"type:varchar(20);default:pending"` // pending, approved, rejected
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	
	ArticleID uint    `json:"article_id"`
	Article   Article `json:"article" gorm:"foreignKey:ArticleID"`
	ParentID  *uint   `json:"parent_id"`
	Parent    *Comment `json:"parent" gorm:"foreignKey:ParentID"`
	Replies   []Comment `json:"replies" gorm:"foreignKey:ParentID"`
}
