package repositories

import (
	"card-system/backend/internal/models"

	"gorm.io/gorm"
)

// UserRepository 用户仓储接口
type UserRepository interface {
	Create(user *models.User) error
	GetByUsername(username string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetByPhone(phone string) (*models.User, error)
	GetByID(id uint) (*models.User, error)
	Update(user *models.User) error
}

// UserRepositoryImpl 用户仓储实现
type UserRepositoryImpl struct {
	*BaseRepository
}

// NewUserRepository 创建用户仓储实例
func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{NewBaseRepository(db)}
}

// Create 创建用户
func (r *UserRepositoryImpl) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// GetByUsername 根据用户名查询用户
func (r *UserRepositoryImpl) GetByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

// 其他方法类似实现...
