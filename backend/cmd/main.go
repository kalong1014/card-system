package main

import (
	"card-system/internal/config"
	"card-system/internal/database"
	"card-system/internal/middleware"
	"card-system/internal/router"
	"card-system/pkg/logger"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig(".")
	if err != nil {
		logger.Fatalf("加载配置失败: %v", err)
	}

	// 初始化数据库
	db, err := database.ConnectDB(cfg)
	if err != nil {
		logger.Fatalf("连接数据库失败: %v", err)
	}

	// 初始化Redis
	redisClient := database.ConnectRedis(cfg)

	// 设置Gin模式
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 创建Gin引擎
	r := gin.Default()

	// 中间件
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())
	r.Use(middleware.CORS())
	r.Use(middleware.AuthMiddleware(redisClient))

	// 路由
	router.SetupRouter(r, db, redisClient)

	// 创建服务器
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.ServerPort),
		Handler: r,
	}

	// 启动服务器
	go func() {
		logger.Infof("服务器启动在端口: %d", cfg.ServerPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("启动服务器失败: %v", err)
		}
	}()

	// 优雅关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("正在关闭服务器...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("服务器关闭失败: %v", err)
	}

	logger.Info("服务器已优雅关闭")
}    