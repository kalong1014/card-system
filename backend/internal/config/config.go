// backend/internal/config/config.go
package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int

	JWTSecret string
}

func LoadConfig(path string) (*Config, error) {
	// 设置默认值
	if os.Getenv("SERVER_PORT") == "" {
		if err := os.Setenv("SERVER_PORT", "8080"); err != nil {
			return nil, fmt.Errorf("failed to set default SERVER_PORT: %v", err)
		}
	}

	if os.Getenv("DB_PORT") == "" {
		if err := os.Setenv("DB_PORT", "3306"); err != nil {
			return nil, fmt.Errorf("failed to set default DB_PORT: %v", err)
		}
	}

	if os.Getenv("REDIS_PORT") == "" {
		if err := os.Setenv("REDIS_PORT", "6379"); err != nil {
			return nil, fmt.Errorf("failed to set default REDIS_PORT: %v", err)
		}
	}

	if os.Getenv("DB_SSL_MODE") == "" {
		if err := os.Setenv("DB_SSL_MODE", "disable"); err != nil {
			return nil, fmt.Errorf("failed to set default DB_SSL_MODE: %v", err)
		}
	}

	// 加载 .env 文件（如果存在）
	if err := godotenv.Load(path); err != nil {
		fmt.Println("警告: 未找到.env文件，将使用系统环境变量或默认值")
	}

	config := &Config{
		ServerPort:    os.Getenv("SERVER_PORT"),
		DBHost:        os.Getenv("DB_HOST"),
		DBPort:        os.Getenv("DB_PORT"),
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBName:        os.Getenv("DB_NAME"),
		DBSSLMode:     os.Getenv("DB_SSL_MODE"),
		RedisHost:     os.Getenv("REDIS_HOST"),
		RedisPort:     os.Getenv("REDIS_PORT"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		JWTSecret:     os.Getenv("JWT_SECRET"),
	}

	// 解析 RedisDB 为整数，默认值 0
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err == nil {
		config.RedisDB = redisDB
	} else {
		config.RedisDB = 0 // 默认值
	}

	// 验证必要的配置（仅保留关键参数）
	if config.ServerPort == "" {
		return nil, fmt.Errorf("SERVER_PORT 环境变量未设置")
	}

	// 只要求必须设置 DB_HOST、DB_USER、DB_NAME，其他使用默认值
	if config.DBHost == "" || config.DBUser == "" || config.DBName == "" {
		return nil, fmt.Errorf("数据库配置不完整: DB_HOST、DB_USER、DB_NAME 是必需的")
	}

	return config, nil
}
