package models

import (
	"time"
)

// CardSecret 卡密模型
type CardSecret struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	ProductID uint   `gorm:"not null" json:"product_id"`
	Secret    string `gorm:"size:255;not null;unique" json:"secret"`
	Status    int    `gorm:"default:0;not null" json:"status"` // 0:未使用, 1:已使用
	CreatedAt time.Time
	UpdatedAt time.Time
}
