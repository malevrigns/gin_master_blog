package models

import (
	"time"
)

// Article 文章模型
type Article struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"type:varchar(255);not null"`
	Slug        string    `json:"slug" gorm:"type:varchar(255);uniqueIndex;not null"`
	Content     string    `json:"content" gorm:"type:longtext;not null"`
	Excerpt     string    `json:"excerpt" gorm:"type:text"`
	CoverImage  string    `json:"cover_image" gorm:"type:varchar(500)"`
	Views       int       `json:"views" gorm:"default:0"`
	Likes       int       `json:"likes" gorm:"default:0"`
	Status      string    `json:"status" gorm:"type:varchar(20);default:draft"` // draft, published
	IsTop       bool      `json:"is_top" gorm:"default:false"`
	PublishedAt *time.Time `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	
	// 关联
	AuthorID   uint      `json:"author_id"`
	Author     User      `json:"author" gorm:"foreignKey:AuthorID"`
	CategoryID uint      `json:"category_id"`
	Category   Category  `json:"category" gorm:"foreignKey:CategoryID"`
	Tags       []Tag     `json:"tags" gorm:"many2many:article_tags;"`
	Comments   []Comment `json:"comments" gorm:"foreignKey:ArticleID"`
}
