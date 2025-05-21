package router

import (
	"card-system/backend/internal/controllers"
	"card-system/backend/internal/services"

	"github.com/gin-gonic/gin"
)

// 注意参数顺序：userService → merchantService → cardService
func SetupRouter(userService services.UserService, merchantService services.MerchantService, cardService services.CardSecretService) *gin.Engine {
	r := gin.Default()

	// 创建控制器实例
	userController := controllers.NewUserController(userService)
	merchantController := controllers.NewMerchantController(merchantService) // 商户控制器
	cardController := controllers.NewCardSecretController(cardService)       // 卡密控制器

	// 设置路由
	api := r.Group("/api")
	{
		// 用户路由
		users := api.Group("/users")
		{
			users.POST("/register", userController.Register)
		}

		// 商户路由
		merchants := api.Group("/merchants")
		{
			merchants.POST("/register", merchantController.Register) // 绑定 Register 方法
		}

		// 卡密路由
		cards := api.Group("/cards")
		{
			cards.POST("/generate", cardController.GenerateCardSecrets) // 生成卡密接口
			cards.GET("/product/:product_id", cardController.GetCardSecretsByProduct)
		}
	}

	return r
}
