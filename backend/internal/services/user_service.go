package services

import (
	"card-system/backend/internal/models"
	"card-system/backend/internal/repositories"
	"context"
)

// UserService 用户服务接口
type UserService interface {
	Register(ctx context.Context, user *models.User) error
}

// UserServiceImpl 用户服务实现
type UserServiceImpl struct {
	userRepo repositories.UserRepository
}

// NewUserService 创建用户服务实例
func NewUserService(userRepo repositories.UserRepository) UserService {
	return &UserServiceImpl{userRepo: userRepo}
}

// Register 用户注册
func (s *UserServiceImpl) Register(ctx context.Context, user *models.User) error {
	return s.userRepo.Create(ctx, user)
}
