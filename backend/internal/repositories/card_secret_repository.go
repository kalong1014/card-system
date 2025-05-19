package repositories

import (
	"card-system/backend/internal/models"
	"context"

	"gorm.io/gorm"
)

// CardSecretRepository 定义卡密仓库接口
type CardSecretRepository interface {
	BatchCreate(ctx context.Context, cardSecrets []*models.CardSecret) error
	GetByProduct(ctx context.Context, productID uint) ([]*models.CardSecret, error)
	// 其他方法...
}

// CardSecretRepositoryImpl 卡密仓库实现
type CardSecretRepositoryImpl struct {
	db *gorm.DB
}

// NewCardSecretRepository 创建卡密仓库实例
func NewCardSecretRepository(db *gorm.DB) CardSecretRepository {
	return &CardSecretRepositoryImpl{db: db}
}

// BatchCreate 批量创建卡密
func (r *CardSecretRepositoryImpl) BatchCreate(ctx context.Context, cardSecrets []*models.CardSecret) error {
	return r.db.WithContext(ctx).Create(cardSecrets).Error
}

// GetByProduct 根据产品ID获取卡密
func (r *CardSecretRepositoryImpl) GetByProduct(ctx context.Context, productID uint) ([]*models.CardSecret, error) {
	var cardSecrets []*models.CardSecret
	err := r.db.WithContext(ctx).Where("product_id = ?", productID).Find(&cardSecrets).Error
	return cardSecrets, err
}
