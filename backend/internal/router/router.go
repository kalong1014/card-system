package router

import (
	"card-system/backend/internal/controllers"
	"card-system/backend/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userService services.UserService, cardService services.CardSecretService) *gin.Engine {
	r := gin.Default()

	// 创建控制器实例，直接使用接口类型
	userController := controllers.NewUserController(userService)
	cardController := controllers.NewCardSecretController(cardService)

	// 设置路由
	api := r.Group("/api")
	{
		// 用户路由
		users := api.Group("/users")
		{
			users.POST("/register", userController.Register)
			// 其他用户路由...
		}

		// 卡密路由
		cards := api.Group("/cards")
		{
			cards.POST("/generate", cardController.GenerateCardSecrets)
			cards.GET("/product/:product_id", cardController.GetCardSecretsByProduct)
			// 其他卡密路由...
		}
	}

	return r
}
