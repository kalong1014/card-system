package models

import "time"

// Merchant 商户模型
type Merchant struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"size:255;not null" json:"name"`
	Email     string `gorm:"size:255;not null;unique" json:"email"`
	Phone     string `gorm:"size:20" json:"phone"`
	Address   string `gorm:"size:255" json:"address"`
	Status    int    `gorm:"default:1;not null" json:"status"` // 1:活跃, 0:禁用
	CreatedAt time.Time
	UpdatedAt time.Time
}
