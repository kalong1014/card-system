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

// NewUserController 创建用户控制器实例
func NewUserController(service services.UserService) *UserController {
	return &UserController{userService: service}
}

// Register 用户注册
func (c *UserController) Register(ctx *gin.Context) {
	// 实现代码...
	response.SuccessWithMessage(ctx, "注册成功")
}
