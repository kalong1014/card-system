package controllers

import (
	"card-system/internal/models"
	"card-system/internal/services"
	"card-system/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CardSecretController 卡密控制器
type CardSecretController struct {
	cardSecretService services.CardSecretService
}

// NewCardSecretController 创建卡密控制器
func NewCardSecretController(cardSecretService services.CardSecretService) *CardSecretController {
	return &CardSecretController{
		cardSecretService: cardSecretService,
	}
}

// GenerateCardSecrets 生成卡密
func (c *CardSecretController) GenerateCardSecrets(ctx *gin.Context) {
	// 获取当前商户ID
	merchantID, exists := ctx.Get("merchant_id")
	if !exists {
		response.Error(ctx, http.StatusUnauthorized, "未授权的商户操作")
		return
	}

	// 解析请求参数
	var req struct {
		ProductID uint `json:"product_id" binding:"required"`
		Count     int  `json:"count" binding:"required,min=1,max=1000"`
		Length    int  `json:"length" binding:"required,min=8,max=32"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// 生成卡密
	cardSecrets, err := c.cardSecretService.GenerateCardSecrets(ctx, req.ProductID, merchantID.(uint), req.Count, req.Length)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(ctx, http.StatusCreated, gin.H{
		"card_secrets": cardSecrets,
		"count":        len(cardSecrets),
	}, "卡密生成成功")
}

// GetCardSecretsByProduct 获取商品卡密列表
func (c *CardSecretController) GetCardSecretsByProduct(ctx *gin.Context) {
	// 获取当前商户ID
	merchantID, exists := ctx.Get("merchant_id")
	if !exists {
		response.Error(ctx, http.StatusUnauthorized, "未授权的商户操作")
		return
	}

	// 获取商品ID
	productID, err := strconv.ParseUint(ctx.Param("product_id"), 10, 64)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "无效的商品ID")
		return
	}

	// 获取卡密列表
	cardSecrets, err := c.cardSecretService.GetCardSecretsByProduct(ctx, uint(productID))
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// 过滤非当前商户的卡密
	var filteredCardSecrets []*models.CardSecret
	for _, cardSecret := range cardSecrets {
		if cardSecret.MerchantID == merchantID.(uint) {
			filteredCardSecrets = append(filteredCardSecrets, cardSecret)
		}
	}

	response.Success(ctx, http.StatusOK, gin.H{
		"card_secrets": filteredCardSecrets,
		"count":        len(filteredCardSecrets),
	}, "卡密列表获取成功")
}

// UseCardSecret 使用卡密
func (c *CardSecretController) UseCardSecret(ctx *gin.Context) {
	// 获取当前用户ID
	userID, exists := ctx.Get("user_id")
	if !exists {
		response.Error(ctx, http.StatusUnauthorized, "未授权的用户操作")
		return
	}

	// 解析请求参数
	var req struct {
		Secret  string `json:"secret" binding:"required"`
		OrderID uint   `json:"order_id" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// 使用卡密
	cardSecret, err := c.cardSecretService.UseCardSecret(ctx, req.Secret, req.OrderID, userID.(string))
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(ctx, http.StatusOK, gin.H{
		"card_secret": cardSecret,
	}, "卡密使用成功")
}    