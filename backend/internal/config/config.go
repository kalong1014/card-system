package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config 应用配置结构
type Config struct {
	// 服务器配置
	ServerPort string

	// 数据库配置
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	// Redis配置
	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int

	// JWT配置
	JWTSecret string
}

// LoadConfig 从环境变量加载配置
func LoadConfig(path string) (*Config, error) {
	// 加载.env文件
	if err := godotenv.Load(path); err != nil {
		fmt.Println("警告: 未找到.env文件，将使用系统环境变量")
	}

	config := &Config{
		ServerPort: os.Getenv("SERVER_PORT"),

		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBSSLMode:  os.Getenv("DB_SSL_MODE"),

		RedisHost:     os.Getenv("REDIS_HOST"),
		RedisPort:     os.Getenv("REDIS_PORT"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		RedisDB:       parseInt(os.Getenv("REDIS_DB"), 0),

		JWTSecret: os.Getenv("JWT_SECRET"),
	}

	// 验证必要的配置项
	if config.ServerPort == "" {
		return nil, fmt.Errorf("SERVER_PORT环境变量未设置")
	}

	if config.DBHost == "" || config.DBPort == "" || config.DBUser == "" || config.DBName == "" {
		return nil, fmt.Errorf("数据库配置不完整")
	}

	return config, nil
}

// parseInt 将字符串转换为整数，失败时返回默认值
func parseInt(s string, defaultValue int) int {
	if s == "" {
		return defaultValue
	}
	var result int
	fmt.Sscanf(s, "%d", &result)
	return result
}
