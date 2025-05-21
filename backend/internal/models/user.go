package models

import "time"

// User 用户模型
type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Username  string `gorm:"size:255;not null;unique" json:"username"`
	Email     string `gorm:"size:255;not null;unique" json:"email"`
	Password  string `gorm:"size:255;not null" json:"-"`
	Role      string `gorm:"size:50;default:'user'" json:"role"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
