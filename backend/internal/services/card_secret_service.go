package services

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"card-system/backend/internal/models"
	"card-system/backend/internal/repositories"
)

// CardSecretService 卡密服务接口
type CardSecretService interface {
	GenerateCardSecrets(ctx context.Context, merchantID uint, productID uint, count int, expireDays int) ([]*models.CardSecret, error)
	GetCardSecretsByProduct(ctx context.Context, productID uint) ([]*models.CardSecret, error)
	// 其他方法...
}

// CardSecretServiceImpl 卡密服务实现
type CardSecretServiceImpl struct {
	cardRepo repositories.CardSecretRepository
}

// NewCardSecretService 创建卡密服务实例
func NewCardSecretService(repo repositories.CardSecretRepository) CardSecretService {
	return &CardSecretServiceImpl{cardRepo: repo}
}

// GenerateCardSecrets 生成卡密
func (s *CardSecretServiceImpl) GenerateCardSecrets(ctx context.Context, merchantID uint, productID uint, count int, expireDays int) ([]*models.CardSecret, error) {
	if count <= 0 || count > 1000 {
		return nil, errors.New("生成数量必须在1-1000之间")
	}

	var cardSecrets []*models.CardSecret
	for i := 0; i < count; i++ {
		card := &models.CardSecret{
			MerchantID: merchantID,
			ProductID:  productID,
			Code:       generateRandomCode(16),
			Status:     0, // 修正：使用数字0表示未使用状态
			ExpiresAt:  time.Now().AddDate(0, 0, expireDays),
			CreatedAt:  time.Now(),
		}
		cardSecrets = append(cardSecrets, card)
	}

	err := s.cardRepo.BatchCreate(ctx, cardSecrets)
	if err != nil {
		return nil, err
	}

	return cardSecrets, nil
}

// GetCardSecretsByProduct 获取产品相关卡密
func (s *CardSecretServiceImpl) GetCardSecretsByProduct(ctx context.Context, productID uint) ([]*models.CardSecret, error) {
	return s.cardRepo.GetByProduct(ctx, productID)
}

// 生成随机卡密
func generateRandomCode(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
