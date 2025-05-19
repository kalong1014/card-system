package services

import (
	"card-system/backend/internal/models"
	"card-system/backend/internal/repositories"
	"context"
)

// UserService 用户服务接口
type UserService interface {
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	// 其他方法...
}

// UserServiceImpl 用户服务实现
type UserServiceImpl struct {
	userRepo repositories.UserRepository
}

// NewUserService 创建用户服务实例
func NewUserService(repo repositories.UserRepository) UserService {
	return &UserServiceImpl{userRepo: repo}
}

// GetUserByEmail 根据邮箱获取用户
func (s *UserServiceImpl) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return s.userRepo.GetByEmail(ctx, email)
}

// 其他方法实现...
