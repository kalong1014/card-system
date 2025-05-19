package repositories

import "gorm.io/gorm"

// BaseRepository 基础仓储
type BaseRepository struct {
	db *gorm.DB
}

// NewBaseRepository 创建基础仓储实例
func NewBaseRepository(db *gorm.DB) *BaseRepository {
	return &BaseRepository{db: db}
}
