package controllers

import (
	"card-system/backend/internal/models"
	"card-system/backend/internal/services"
	"card-system/backend/pkg/response"

	"github.com/gin-gonic/gin"
)

// MerchantController 商户控制器
type MerchantController struct {
	merchantService services.MerchantService
}

// NewMerchantController 创建商户控制器实例
func NewMerchantController(service services.MerchantService) *MerchantController {
	return &MerchantController{merchantService: service}
}

// Register 处理商户注册请求
func (c *MerchantController) Register(ctx *gin.Context) {
	// 1. 解析请求体中的商户信息
	var req models.Merchant
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, 400, "请求参数格式错误")
		return
	}

	// 2. 调用服务层注册商户
	if err := c.merchantService.Register(ctx, &req); err != nil {
		response.Error(ctx, 500, "商户注册失败: "+err.Error())
		return
	}

	// 3. 返回成功响应
	response.SuccessWithMessage(ctx, "商户注册成功")
}
