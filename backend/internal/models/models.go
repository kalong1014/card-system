// backend/internal/models/models.go
package models

import (
	"time"

	"gorm.io/gorm"
)

// 基础模型
type BaseModel struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// 用户模型
type User struct {
	BaseModel
	Username    string    `gorm:"uniqueIndex;not null" json:"username"`
	Password    string    `gorm:"not null" json:"-"`
	Email       string    `gorm:"uniqueIndex;not null" json:"email"`
	Phone       string    `gorm:"uniqueIndex;size:11" json:"phone"`
	Status      int       `gorm:"default:1" json:"status"`    // 1: 正常, 0: 禁用
	Role        string    `gorm:"default:'user'" json:"role"` // user, merchant, admin
	MerchantID  uint      `gorm:"default:0" json:"merchant_id"`
	LastLoginAt time.Time `json:"last_login_at"`
}

// 商户模型
type Merchant struct {
	BaseModel
	UserID      uint    `gorm:"uniqueIndex;not null" json:"user_id"`
	Name        string  `gorm:"not null" json:"name"`
	Contact     string  `gorm:"not null" json:"contact"`
	Phone       string  `gorm:"not null;size:11" json:"phone"`
	Email       string  `gorm:"not null" json:"email"`
	Status      int     `gorm:"default:0" json:"status"` // 0: 待审核, 1: 已审核, 2: 已拒绝
	Balance     float64 `gorm:"default:0.0" json:"balance"`
	Withdrawals float64 `gorm:"default:0.0" json:"withdrawals"`
	TotalIncome float64 `gorm:"default:0.0" json:"total_income"`
}

// 商品模型
type Product struct {
	BaseModel
	MerchantID  uint    `gorm:"not null" json:"merchant_id"`
	Name        string  `gorm:"not null" json:"name"`
	Category    string  `gorm:"not null" json:"category"`
	Price       float64 `gorm:"not null" json:"price"`
	Stock       int     `gorm:"default:0" json:"stock"`
	Sold        int     `gorm:"default:0" json:"sold"`
	Description string  `gorm:"type:text" json:"description"`
	Status      int     `gorm:"default:1" json:"status"` // 1: 上架, 0: 下架
}

// 卡密模型
type CardSecret struct {
	BaseModel
	ProductID uint      `gorm:"not null" json:"product_id"`
	Secret    string    `gorm:"uniqueIndex;not null" json:"secret"`
	Status    int       `gorm:"default:0" json:"status"` // 0: 未使用, 1: 已使用, 2: 已锁定
	OrderID   uint      `gorm:"default:0" json:"order_id"`
	UsedAt    time.Time `json:"used_at"`
	UserID    uint      `json:"user_id"`
}

// 订单模型
type Order struct {
	BaseModel
	OrderNo     string    `gorm:"uniqueIndex;not null" json:"order_no"`
	UserID      uint      `gorm:"not null" json:"user_id"`
	MerchantID  uint      `gorm:"not null" json:"merchant_id"`
	ProductID   uint      `gorm:"not null" json:"product_id"`
	Amount      float64   `gorm:"not null" json:"amount"`
	Status      int       `gorm:"default:0" json:"status"` // 0: 待支付, 1: 已支付, 2: 已完成, 3: 已取消
	PaymentType string    `gorm:"not null" json:"payment_type"`
	PaymentNo   string    `json:"payment_no"`
	PaidAt      time.Time `json:"paid_at"`
	CompletedAt time.Time `json:"completed_at"`
	CanceledAt  time.Time `json:"canceled_at"`
	Remark      string    `json:"remark"`
}

// 支付渠道模型
type PaymentChannel struct {
	BaseModel
	MerchantID uint   `gorm:"not null" json:"merchant_id"`
	Name       string `gorm:"not null" json:"name"`
	Type       string `gorm:"not null" json:"type"`    // alipay, wechat, unionpay
	Status     int    `gorm:"default:1" json:"status"` // 1: 启用, 0: 禁用
	Config     string `gorm:"type:text" json:"config"` // 配置信息，JSON格式
}
