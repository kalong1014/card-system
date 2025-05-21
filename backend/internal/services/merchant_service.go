package services

import (
	"card-system/backend/internal/models"
	"card-system/backend/internal/repositories"
	"context"
)

// MerchantService 商户服务接口
type MerchantService interface {
	Register(ctx context.Context, merchant *models.Merchant) error // 必须包含此方法
}

// MerchantServiceImpl 商户服务实现
type MerchantServiceImpl struct {
	merchantRepo repositories.MerchantRepository
}

// NewMerchantService 创建商户服务实例
func NewMerchantService(merchantRepo repositories.MerchantRepository) MerchantService {
	return &MerchantServiceImpl{merchantRepo: merchantRepo}
}

// Register 实现商户注册逻辑（示例）
func (s *MerchantServiceImpl) Register(ctx context.Context, merchant *models.Merchant) error {
	// 可选：添加校验逻辑（如邮箱唯一性检查）
	// 调用仓库层创建商户
	return s.merchantRepo.Create(ctx, merchant)
}
