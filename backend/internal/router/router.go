package router

import (
	"card-system/backend/internal/controllers"
	"card-system/backend/internal/middleware"
	"card-system/backend/internal/repositories"
	"card-system/backend/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func SetupRouter(r *gin.Engine, db *gorm.DB, redisClient *redis.Client) {
	// 仓储实例
	userRepo := repositories.NewUserRepository(db)

	// 服务实例
	userService := services.NewUserService(userRepo)

	// 控制器实例
	userCtrl := controllers.NewUserController(userService)

	// 路由分组
	api := r.Group("/api")
	api.POST("/users/register", userCtrl.Register)

	// 应用中间件
	r.Use(
		middleware.Logger(),
		middleware.Recovery(),
		middleware.CORS(),
	)
}
