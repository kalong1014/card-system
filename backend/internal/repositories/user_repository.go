package repositories

import (
	"card-system/backend/internal/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

// UserRepository 定义用户仓库接口
type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	Create(ctx context.Context, user *models.User) error
}

// UserRepositoryImpl 实现 UserRepository 接口
type UserRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepository 创建用户仓库实例
func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

// GetByEmail 根据邮箱查找用户
func (r *UserRepositoryImpl) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	result := r.db.WithContext(ctx).Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil // 返回nil表示用户不存在
		}
		return nil, result.Error
	}
	return &user, nil
}

// Create 创建用户
func (r *UserRepositoryImpl) Create(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}
