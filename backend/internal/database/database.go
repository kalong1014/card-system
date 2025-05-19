package database

import (
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

// InitDB 初始化数据库连接
func InitDB() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"your_username",
		"your_password",
		"your_host",
		"your_port",
		"your_dbname",
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.Log.Fatal("Failed to connect database: %v", err)
		return err
	}

	utils.Log.Info("Database connected successfully")
	return nil // 添加缺失的返回语句
}

// InitRedis 初始化Redis连接
func InitRedis() error {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// 修正：使用Redis包的上下文
	ctx := context.Background()
	_, err := Redis.Ping(ctx).Result()
	if err != nil {
		utils.Log.Fatal("Failed to connect redis: %v", err)
		return err
	}

	utils.Log.Info("Redis connected successfully")
	return nil
}
