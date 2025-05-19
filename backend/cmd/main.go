package main

import (
	"card-system/backend/internal/database"
	"card-system/backend/internal/repositories"
	"card-system/backend/internal/router"
	"card-system/backend/internal/services"
	"card-system/backend/utils"
	"net/http"
)

func main() {
	// 初始化数据库连接
	if err := database.InitDB(); err != nil {
		utils.Log.Fatal("Failed to initialize database: %v", err)
	}

	// 初始化Redis连接
	if err := database.InitRedis(); err != nil {
		utils.Log.Fatal("Failed to initialize redis: %v", err)
	}

	// 创建仓库实例
	userRepo := repositories.NewUserRepository(database.DB)
	cardRepo := repositories.NewCardSecretRepository(database.DB)

	// 创建服务实例
	userService := services.NewUserService(userRepo)
	cardService := services.NewCardSecretService(cardRepo)

	// 设置路由
	r := router.SetupRouter(userService, cardService)

	// 启动服务器
	utils.Log.Info("Server started on port :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		utils.Log.Fatal("Server failed to start: %v", err)
	}
}
