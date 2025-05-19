package services

import (
	"card-system/internal/models"
	"card-system/internal/repositories"
	"card-system/pkg/logger"
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// CardSecretService 卡密服务接口
type CardSecretService interface {
	GenerateCardSecrets(ctx context.Context, productID, merchantID uint, count int, length int) ([]*models.CardSecret, error)
	GetCardSecretsByProduct(ctx context.Context, productID uint) ([]*models.CardSecret, error)
	GetCardSecretByID(ctx context.Context, id uint) (*models.CardSecret, error)
	UpdateCardSecret(ctx context.Context, cardSecret *models.CardSecret) error
	UseCardSecret(ctx context.Context, secret string, orderID uint, usedBy string) (*models.CardSecret, error)
}

// CardSecretServiceImpl 卡密服务实现
type CardSecretServiceImpl struct {
	cardSecretRepo repositories.CardSecretRepository
	productRepo    repositories.ProductRepository
}

// NewCardSecretService 创建卡密服务
func NewCardSecretService(cardSecretRepo repositories.CardSecretRepository, productRepo repositories.ProductRepository) CardSecretService {
	return &CardSecretServiceImpl{
		cardSecretRepo: cardSecretRepo,
		productRepo:    productRepo,
	}
}

// GenerateCardSecrets 生成卡密
func (s *CardSecretServiceImpl) GenerateCardSecrets(ctx context.Context, productID, merchantID uint, count int, length int) ([]*models.CardSecret, error) {
	// 检查商品是否存在
	product, err := s.productRepo.GetByID(ctx, productID)
	if err != nil {
		return nil, fmt.Errorf("商品不存在: %v", err)
	}

	if product.MerchantID != merchantID {
		return nil, fmt.Errorf("无权限为该商品生成卡密")
	}

	// 生成卡密
	cardSecrets := make([]*models.CardSecret, count)
	now := time.Now()

	for i := 0; i < count; i++ {
		secret := s.generateRandomSecret(length)
		cardSecrets[i] = &models.CardSecret{
			Secret:     secret,
			ProductID:  productID,
			Status:     0, // 未使用
			MerchantID: merchantID,
			CreatedAt:  now,
			UpdatedAt:  now,
		}
	}

	// 保存卡密
	err = s.cardSecretRepo.CreateBatch(ctx, cardSecrets)
	if err != nil {
		logger.Errorf("批量创建卡密失败: %v", err)
		return nil, fmt.Errorf("生成卡密失败: %v", err)
	}

	// 更新商品库存
	product.Stock += count
	err = s.productRepo.Update(ctx, product)
	if err != nil {
		logger.Errorf("更新商品库存失败: %v", err)
		// 这里可以考虑添加事务回滚逻辑
	}

	return cardSecrets, nil
}

// GetCardSecretsByProduct 获取商品的卡密列表
func (s *CardSecretServiceImpl) GetCardSecretsByProduct(ctx context.Context, productID uint) ([]*models.CardSecret, error) {
	return s.cardSecretRepo.GetByProductID(ctx, productID)
}

// GetCardSecretByID 获取卡密详情
func (s *CardSecretServiceImpl) GetCardSecretByID(ctx context.Context, id uint) (*models.CardSecret, error) {
	return s.cardSecretRepo.GetByID(ctx, id)
}

// UpdateCardSecret 更新卡密信息
func (s *CardSecretServiceImpl) UpdateCardSecret(ctx context.Context, cardSecret *models.CardSecret) error {
	return s.cardSecretRepo.Update(ctx, cardSecret)
}

// UseCardSecret 使用卡密
func (s *CardSecretServiceImpl) UseCardSecret(ctx context.Context, secret string, orderID uint, usedBy string) (*models.CardSecret, error) {
	// 查询卡密
	cardSecret, err := s.cardSecretRepo.GetBySecret(ctx, secret)
	if err != nil {
		return nil, fmt.Errorf("卡密不存在: %v", err)
	}

	// 检查卡密状态
	if cardSecret.Status != 0 {
		return nil, fmt.Errorf("卡密已被使用或锁定")
	}

	// 更新卡密状态
	cardSecret.Status = 1
	cardSecret.UsedAt = time.Now()
	cardSecret.UsedBy = usedBy
	cardSecret.OrderID = orderID

	err = s.cardSecretRepo.Update(ctx, cardSecret)
	if err != nil {
		return nil, fmt.Errorf("使用卡密失败: %v", err)
	}

	return cardSecret, nil
}

// generateRandomSecret 生成随机卡密
func (s *CardSecretServiceImpl) generateRandomSecret(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	sb.Grow(length)
	for i := 0; i < length; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}    