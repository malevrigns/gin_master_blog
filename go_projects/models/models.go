package models

import (
	"time"

	"gorm.io/datatypes"
)

// User 用户模型
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"type:varchar(100);uniqueIndex;not null"`
	Email     string    `json:"email" gorm:"type:varchar(255);uniqueIndex;not null"`
	Password  string    `json:"-" gorm:"type:varchar(255);not null"`
	Avatar    string    `json:"avatar" gorm:"type:varchar(500)"`
	Role      string    `json:"role" gorm:"type:varchar(20);default:user"` // admin, user
	Bio       string    `json:"bio" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

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

// Tag 标签模型
type Tag struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar(100);uniqueIndex;not null"`
	Slug      string    `json:"slug" gorm:"type:varchar(100);uniqueIndex;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	
	Articles []Article `json:"articles" gorm:"many2many:article_tags;"`
}

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

