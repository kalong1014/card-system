package repositories

import (
	"card-system/backend/internal/models"
	"context"

	"gorm.io/gorm"
)

// MerchantRepository 商户仓储接口
type MerchantRepository interface {
	Create(ctx context.Context, merchant *models.Merchant) error
	GetByID(ctx context.Context, id uint) (*models.Merchant, error)
	Update(ctx context.Context, merchant *models.Merchant) error
	List(ctx context.Context) ([]*models.Merchant, error)
}

// MerchantRepositoryImpl 商户仓储实现
type MerchantRepositoryImpl struct {
	db *gorm.DB
}

// NewMerchantRepository 创建商户仓储实例
func NewMerchantRepository(db *gorm.DB) MerchantRepository {
	return &MerchantRepositoryImpl{db: db}
}

// Create 创建商户
func (r *MerchantRepositoryImpl) Create(ctx context.Context, merchant *models.Merchant) error {
	return r.db.WithContext(ctx).Create(merchant).Error
}

// GetByID 根据ID获取商户
func (r *MerchantRepositoryImpl) GetByID(ctx context.Context, id uint) (*models.Merchant, error) {
	var merchant models.Merchant
	err := r.db.WithContext(ctx).First(&merchant, id).Error
	if err != nil {
		return nil, err
	}
	return &merchant, nil
}

// Update 更新商户
func (r *MerchantRepositoryImpl) Update(ctx context.Context, merchant *models.Merchant) error {
	return r.db.WithContext(ctx).Save(merchant).Error
}

// List 获取所有商户
func (r *MerchantRepositoryImpl) List(ctx context.Context) ([]*models.Merchant, error) {
	var merchants []*models.Merchant
	err := r.db.WithContext(ctx).Find(&merchants).Error
	if err != nil {
		return nil, err
	}
	return merchants, nil
}
