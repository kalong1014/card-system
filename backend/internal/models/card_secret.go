package models

import (
	"time"
)

// CardSecret 卡密模型
type CardSecret struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	MerchantID uint      `gorm:"not null" json:"merchant_id"`
	ProductID  uint      `gorm:"not null" json:"product_id"`
	Code       string    `gorm:"size:32;uniqueIndex;not null" json:"code"`
	Status     int       `gorm:"default:0;not null" json:"status"` // 0:未使用, 1:已使用, 2:已过期
	ExpiresAt  time.Time `gorm:"not null" json:"expires_at"`
	UsedAt     time.Time `gorm:"default:null" json:"used_at"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
