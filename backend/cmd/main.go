package main

import (
	"card-system/backend/internal/config"
	"card-system/backend/internal/database"
	"card-system/backend/internal/repositories"
	"card-system/backend/internal/router"
	"card-system/backend/internal/services"
	"card-system/backend/utils"
	"net/http"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig("../.env")
	if err != nil {
		utils.Log.Fatal("Failed to load config: %v", err)
	}

	// 初始化数据库（传递 cfg 参数）
	if err := database.InitDB(cfg); err != nil {
		utils.Log.Fatal("Failed to initialize database: %v", err)
	}

	// 初始化 Redis（传递 cfg 参数）
	if err := database.InitRedis(cfg); err != nil {
		utils.Log.Fatal("Failed to initialize redis: %v", err)
	}

	// 创建仓库实例
	userRepo := repositories.NewUserRepository(database.DB)
	merchantRepo := repositories.NewMerchantRepository(database.DB)
	cardRepo := repositories.NewCardSecretRepository(database.DB)

	// 创建服务实例
	userService := services.NewUserService(userRepo)
	merchantService := services.NewMerchantService(merchantRepo) // 商户服务
	cardService := services.NewCardSecretService(cardRepo)

	// 设置路由（确保传递所有服务）
	r := router.SetupRouter(userService, merchantService, cardService)

	// 启动服务器
	utils.Log.Info("Server started on port :%s", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, r); err != nil {
		utils.Log.Fatal("Server failed to start: %v", err)
	}
}
