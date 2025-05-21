package services

import (
	"card-system/backend/internal/models"
	"card-system/backend/internal/repositories"
	"context"
)

// CardSecretService 卡密服务接口
type CardSecretService interface {
	GenerateCardSecrets(ctx context.Context, productID uint, count int) error
	GetCardSecretsByProduct(ctx context.Context, productID uint) ([]*models.CardSecret, error)
}

// CardSecretServiceImpl 卡密服务实现
type CardSecretServiceImpl struct {
	cardRepo repositories.CardSecretRepository
}

// NewCardSecretService 创建卡密服务实例
func NewCardSecretService(cardRepo repositories.CardSecretRepository) CardSecretService {
	return &CardSecretServiceImpl{cardRepo: cardRepo}
}

// GenerateCardSecrets 生成卡密
func (s *CardSecretServiceImpl) GenerateCardSecrets(ctx context.Context, productID uint, count int) error {
	var cardSecrets []*models.CardSecret
	for i := 0; i < count; i++ {
		// 生成卡密逻辑...
		cardSecret := &models.CardSecret{
			ProductID: productID,
			Secret:    "generated_secret", // 替换为实际生成逻辑
		}
		cardSecrets = append(cardSecrets, cardSecret)
	}
	return s.cardRepo.BatchCreate(ctx, cardSecrets)
}

// GetCardSecretsByProduct 根据产品ID获取卡密列表
func (s *CardSecretServiceImpl) GetCardSecretsByProduct(ctx context.Context, productID uint) ([]*models.CardSecret, error) {
	return s.cardRepo.GetByProduct(ctx, productID)
}
