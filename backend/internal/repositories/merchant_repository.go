package repositories

import (
	"card-system/backend/internal/models"

	"gorm.io/gorm"
)

// MerchantRepository 商户仓储接口
type MerchantRepository interface {
	Create(merchant *models.Merchant) error
	GetByID(id uint) (*models.Merchant, error)
	GetByUserID(userID uint) (*models.Merchant, error)
	Update(merchant *models.Merchant) error
	List() ([]*models.Merchant, error)
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
func (r *MerchantRepositoryImpl) Create(merchant *models.Merchant) error {
	return r.db.Create(merchant).Error
}

// GetByID 根据ID获取商户
func (r *MerchantRepositoryImpl) GetByID(id uint) (*models.Merchant, error) {
	var merchant models.Merchant
	err := r.db.First(&merchant, id).Error
	if err != nil {
		return nil, err
	}
	return &merchant, nil
}

// GetByUserID 根据用户ID获取商户
func (r *MerchantRepositoryImpl) GetByUserID(userID uint) (*models.Merchant, error) {
	var merchant models.Merchant
	err := r.db.Where("user_id = ?", userID).First(&merchant).Error
	if err != nil {
		return nil, err
	}
	return &merchant, nil
}

// Update 更新商户
func (r *MerchantRepositoryImpl) Update(merchant *models.Merchant) error {
	return r.db.Save(merchant).Error
}

// List 获取所有商户
func (r *MerchantRepositoryImpl) List() ([]*models.Merchant, error) {
	var merchants []*models.Merchant
	err := r.db.Find(&merchants).Error
	if err != nil {
		return nil, err
	}
	return merchants, nil
}
