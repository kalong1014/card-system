package database

import (
	"context"
	"fmt"
	"time"

	"card-system/backend/internal/config"
	"card-system/backend/pkg/logger"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectDB 连接MySQL数据库
func ConnectDB(cfg *config.Config) (*gorm.DB, error) {
	// 构建DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("连接数据库失败: %v", err)
		return nil, err
	}

	// 获取底层SQL连接
	sqlDB, err := db.DB()
	if err != nil {
		logger.Errorf("获取数据库连接失败: %v", err)
		return nil, err
	}

	// 配置连接池
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Minute * 30)

	// 测试连接
	if err := sqlDB.Ping(); err != nil {
		logger.Errorf("测试数据库连接失败: %v", err)
		return nil, err
	}

	logger.Info("数据库连接成功")
	return db, nil
}

// ConnectRedis 连接Redis
func ConnectRedis(cfg *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisHost + ":" + cfg.RedisPort,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

	// 测试连接
	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		logger.Fatalf("连接Redis失败: %v", err)
	}

	logger.Info("Redis连接成功")
	return client
}
