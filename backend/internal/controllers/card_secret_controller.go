package controllers

import (
	"card-system/backend/internal/services"
	"card-system/backend/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CardSecretController 卡密控制器
type CardSecretController struct {
	cardSecretService services.CardSecretService
}

// NewCardSecretController 创建卡密控制器实例
func NewCardSecretController(cardSecretService services.CardSecretService) *CardSecretController {
	return &CardSecretController{cardSecretService: cardSecretService}
}

// GenerateCardSecrets 生成卡密
func (c *CardSecretController) GenerateCardSecrets(ctx *gin.Context) {
	// 获取商品ID
	productIDStr := ctx.Param("product_id")
	productID, err := strconv.ParseUint(productIDStr, 10, 64)
	if err != nil {
		response.Error(ctx, 40001, "无效的商品ID")
		return
	}

	// 获取请求参数
	var req struct {
		Count int `json:"count" binding:"required,min=1,max=1000"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, 40002, "参数错误: "+err.Error())
		return
	}

	// 获取当前商户ID
	merchantID, exists := ctx.Get("merchant_id")
	if !exists {
		response.Error(ctx, 40301, "权限不足")
		return
	}

	// 生成卡密
	cardSecrets, err := c.cardSecretService.GenerateCardSecrets(uint(productID), uint(merchantID.(uint)), req.Count)
	if err != nil {
		response.Error(ctx, 50001, "生成卡密失败: "+err.Error())
		return
	}

	response.Success(ctx, gin.H{
		"count":        len(cardSecrets),
		"card_secrets": cardSecrets,
	})
}

// GetCardSecretsByProduct 获取商品的卡密列表
func (c *CardSecretController) GetCardSecretsByProduct(ctx *gin.Context) {
	// 获取商品ID
	productIDStr := ctx.Param("product_id")
	productID, err := strconv.ParseUint(productIDStr, 10, 64)
	if err != nil {
		response.Error(ctx, 40001, "无效的商品ID")
		return
	}

	// 获取分页参数
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	// 获取当前商户ID
	merchantID, exists := ctx.Get("merchant_id")
	if !exists {
		response.Error(ctx, 40301, "权限不足")
		return
	}

	// 获取卡密列表
	cardSecrets, total, err := c.cardSecretService.GetCardSecretsByProduct(uint(productID), uint(merchantID.(uint)), page, pageSize)
	if err != nil {
		response.Error(ctx, 50001, "获取卡密列表失败: "+err.Error())
		return
	}

	response.Success(ctx, gin.H{
		"total":        total,
		"page":         page,
		"page_size":    pageSize,
		"card_secrets": cardSecrets,
	})
}
