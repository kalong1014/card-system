package services

import (
	"card-system/backend/internal/models"
	"card-system/backend/internal/repositories"
	"card-system/backend/pkg/logger"
	"fmt"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

// UserService 用户服务接口
type UserService interface {
	Register(username, password, email, phone string) (*models.User, error)
	Login(username, password string) (*models.User, error)
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
func (s *UserServiceImpl) Register(username, password, email, phone string) (*models.User, error) {
	// 检查用户名是否存在
	user, err := s.userRepo.GetByUsername(username)
	if err == nil {
		return nil, fmt.Errorf("用户名已存在")
	}

	// 加密密码
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("密码加密失败", zap.Error(err))
		return nil, err
	}

	newUser := &models.User{
		Username: username,
		Password: string(hashedPwd),
		Email:    email,
		Phone:    phone,
		Role:     "user",
		Status:   1,
	}

	return newUser, s.userRepo.Create(newUser)
}

// Login 用户登录
func (s *UserServiceImpl) Login(username, password string) (*models.User, error) {
	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
}
