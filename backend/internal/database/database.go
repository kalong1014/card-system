package database

import (
	"card-system/backend/internal/config"
	"card-system/backend/utils"
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	Redis *redis.Client
)

// InitDB 初始化数据库连接（接收 *config.Config 参数）
func InitDB(cfg *config.Config) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.Log.Fatal("Failed to connect database: %v", err)
		return err
	}

	utils.Log.Info("Database connected successfully")
	return nil
}

// InitRedis 初始化 Redis 连接（接收 *config.Config 参数）
func InitRedis(cfg *config.Config) error {
	Redis = redis.NewClient(&redis.Options{
		Addr:     cfg.RedisHost + ":" + cfg.RedisPort,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

	ctx := context.Background()
	_, err := Redis.Ping(ctx).Result()
	if err != nil {
		utils.Log.Fatal("Failed to connect redis: %v", err)
		return err
	}

	utils.Log.Info("Redis connected successfully")
	return nil
}
