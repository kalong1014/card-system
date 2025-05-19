package controllers

import (
	"card-system/backend/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CardSecretController 卡密控制器
type CardSecretController struct {
	cardService services.CardSecretService
}

// NewCardSecretController 创建卡密控制器实例
func NewCardSecretController(service services.CardSecretService) *CardSecretController {
	return &CardSecretController{cardService: service}
}

// GenerateCardSecrets 生成卡密
func (c *CardSecretController) GenerateCardSecrets(ctx *gin.Context) {
	// 实现代码...
}

// GetCardSecretsByProduct 根据产品ID获取卡密列表
func (c *CardSecretController) GetCardSecretsByProduct(ctx *gin.Context) {
	// 从URL参数获取产品ID
	productIDStr := ctx.Param("product_id")
	productID, err := strconv.ParseUint(productIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	// 调用服务层获取卡密列表
	cardSecrets, err := c.cardService.GetCardSecretsByProduct(ctx, uint(productID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get card secrets"})
		return
	}

	// 返回成功响应
	ctx.JSON(http.StatusOK, gin.H{"data": cardSecrets})
}
