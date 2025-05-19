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
	// 替换为你的实际数据库配置
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root",        // 数据库用户名
		"Aa123789@",   // 数据库密码
		"localhost",   // 数据库主机
		"3306",        // 数据库端口（通常MySQL是3306）
		"card_system", // 数据库名称
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

// InitRedis 初始化Redis连接
func InitRedis() error {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
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
