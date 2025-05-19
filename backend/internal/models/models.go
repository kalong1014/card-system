package models

import (
	"gorm.io/gorm"
	"time"
)

// 用户模型
type User struct {
	gorm.Model
	Username     string     `gorm:"uniqueIndex;not null" json:"username"`
	Password     string     `gorm:"not null" json:"-"`
	Email        string     `gorm:"uniqueIndex;not null" json:"email"`
	Phone        string     `json:"phone"`
	Avatar       string     `json:"avatar"`
	Status       int        `gorm:"default:1" json:"status"` // 1: 正常, 0: 禁用
	Role         int        `gorm:"default:2" json:"role"`   // 1: 管理员, 2: 普通用户, 3: 商户
	LastLoginAt  *time.Time `json:"last_login_at"`
	LastLoginIP  string     `json:"last_login_ip"`
	MerchantID   uint       `json:"merchant_id"`
	Merchant     Merchant   `gorm:"foreignKey:MerchantID" json:"merchant,omitempty"`
}

// 商户模型
type Merchant struct {
	gorm.Model
	Name            string     `gorm:"uniqueIndex;not null" json:"name"`
	ContactPerson   string     `json:"contact_person"`
	ContactPhone    string     `json:"contact_phone"`
	ContactEmail    string     `json:"contact_email"`
	Status          int        `gorm:"default:0" json:"status"` // 0: 待审核, 1: 已通过, 2: 已拒绝
	Level           int        `gorm:"default:1" json:"level"`  // 1: 基础版, 2: 高级版, 3: 旗舰版
	Domain          string     `json:"domain"`                   // 二级域名
	CustomDomain    string     `json:"custom_domain"`            // 自定义域名
	Description     string     `json:"description"`
	Logo            string     `json:"logo"`
	Banner          string     `json:"banner"`
	PageConfig      string     `json:"page_config"` // JSON格式的页面配置
	UserID          uint       `json:"user_id"`
	User            User       `gorm:"foreignKey:UserID" json:"user,omitempty"`
	PaymentChannels []PaymentChannel `gorm:"foreignKey:MerchantID" json:"payment_channels,omitempty"`
}

// 商品模型
type Product struct {
	gorm.Model
	Name         string  `gorm:"not null" json:"name"`
	Description  string  `json:"description"`
	Price        float64 `gorm:"not null" json:"price"`
	OriginalPrice float64 `json:"original_price"`
	Stock        int     `gorm:"default:0" json:"stock"`
	Status       int     `gorm:"default:1" json:"status"` // 1: 上架, 0: 下架
	IsHot        bool    `gorm:"default:false" json:"is_hot"`
	IsNew        bool    `gorm:"default:false" json:"is_new"`
	Image        string  `json:"image"`
	Images       string  `json:"images"` // JSON数组
	CategoryID   uint    `json:"category_id"`
	Category     Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	MerchantID   uint    `json:"merchant_id"`
	Merchant     Merchant `gorm:"foreignKey:MerchantID" json:"merchant,omitempty"`
	CardSecrets  []CardSecret `gorm:"foreignKey:ProductID" json:"card_secrets,omitempty"`
}

// 卡密模型
type CardSecret struct {
	gorm.Model
	Secret      string    `gorm:"uniqueIndex;not null" json:"secret"`
	ProductID   uint      `json:"product_id"`
	Product     Product   `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	OrderID     uint      `json:"order_id"`
	Order       Order     `gorm:"foreignKey:OrderID" json:"order,omitempty"`
	Status      int       `gorm:"default:0" json:"status"` // 0: 未使用, 1: 已使用, 2: 已锁定
	UsedAt      *time.Time `json:"used_at"`
	UsedBy      string    `json:"used_by"`
	Remark      string    `json:"remark"`
	MerchantID  uint      `json:"merchant_id"`
	Merchant    Merchant  `gorm:"foreignKey:MerchantID" json:"merchant,omitempty"`
}

// 订单模型
type Order struct {
	gorm.Model
	OrderNo      string    `gorm:"uniqueIndex;not null" json:"order_no"`
	UserID       uint      `json:"user_id"`
	User         User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	MerchantID   uint      `json:"merchant_id"`
	Merchant     Merchant  `gorm:"foreignKey:MerchantID" json:"merchant,omitempty"`
	ProductID    uint      `json:"product_id"`
	Product      Product   `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	CardSecretID uint      `json:"card_secret_id"`
	CardSecret   CardSecret `gorm:"foreignKey:CardSecretID" json:"card_secret,omitempty"`
	Amount       float64   `gorm:"not null" json:"amount"`
	PaymentType  string    `json:"payment_type"`
	PaymentNo    string    `json:"payment_no"`
	Status       int       `gorm:"default:0" json:"status"` // 0: 未支付, 1: 已支付, 2: 已完成, 3: 已取消
	PaidAt       *time.Time `json:"paid_at"`
	CompletedAt  *time.Time `json:"completed_at"`
	CanceledAt   *time.Time `json:"canceled_at"`
	IP           string    `json:"ip"`
	UserAgent    string    `json:"user_agent"`
}

// 支付渠道模型
type PaymentChannel struct {
	gorm.Model
	Name          string `gorm:"not null" json:"name"`
	Type          string `gorm:"not null" json:"type"` // wechat, alipay, crypto, international
	Status        int    `gorm:"default:1" json:"status"` // 1: 启用, 0: 禁用
	Config        string `json:"config"` // JSON格式的配置
	MerchantID    uint   `json:"merchant_id"`
	Merchant      Merchant `gorm:"foreignKey:MerchantID" json:"merchant,omitempty"`
}

// 页面模板模型
type PageTemplate struct {
	gorm.Model
	Name        string `gorm:"not null" json:"name"`
	Description string `json:"description"`
	Preview     string `json:"preview"`
	Content     string `json:"content"` // JSON格式的页面内容
	IsSystem    bool   `gorm:"default:false" json:"is_system"`
	MerchantID  uint   `json:"merchant_id"`
	Merchant    Merchant `gorm:"foreignKey:MerchantID" json:"merchant,omitempty"`
}    