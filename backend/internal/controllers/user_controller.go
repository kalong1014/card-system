package controllers

import (
	"card-system/backend/internal/services"

	"github.com/gin-gonic/gin"
)

// UserController 用户控制器
type UserController struct {
	userService services.UserService
}

// NewUserController 创建用户控制器实例
func NewUserController(service services.UserService) *UserController { // 改为接收接口类型
	return &UserController{userService: service}
}

// Register 用户注册
func (c *UserController) Register(ctx *gin.Context) {
	// 实现代码...

	ctx.JSON(200, gin.H{
		"message": "注册成功"})
}

// 其他方法...
