package controllers

import (
	"card-system/backend/internal/services"
	"card-system/backend/pkg/response"

	"github.com/gin-gonic/gin"
)

// UserController 用户控制器
type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) Register(ctx *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required,min=6"`
		Email    string `json:"email" binding:"required,email"`
		Phone    string `json:"phone" binding:"required,len=11"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, 400, "参数错误: "+err.Error())
		return
	}

	user, err := c.userService.Register(req.Username, req.Password, req.Email, req.Phone)
	if err != nil {
		response.Error(ctx, 500, "注册失败: "+err.Error())
		return
	}

	response.Success(ctx, gin.H{"user": user})
}

// 其他控制器（Merchant/Product/CardSecret等）类似实现，需补充完整
