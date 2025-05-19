package main

import (
	"card-system/backend/internal/config"
	"card-system/backend/internal/database"
	"card-system/backend/internal/router"
	"card-system/backend/pkg/logger"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		logger.Fatal("配置加载失败", zap.Error(err))
	}

	db, err := database.ConnectDB(cfg)
	if err != nil {
		logger.Fatal("数据库连接失败", zap.Error(err))
	}

	redisClient := database.ConnectRedis(cfg)
	defer redisClient.Close()

	r := gin.New()
	router.SetupRouter(r, db, redisClient)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.ServerPort),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("服务器启动失败", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("服务器关闭失败", zap.Error(err))
	}

	logger.Info("服务器已优雅关闭")
}
