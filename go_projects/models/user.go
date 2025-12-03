package models

import (
	"time"
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
