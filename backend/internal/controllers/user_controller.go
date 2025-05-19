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
func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService: userService}
}

// Register 用户注册接口
func (c *UserController) Register(ctx *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required,min=6"`
		Email    string `json:"email" binding:"required,email"`
		Phone    string `json:"phone" binding:"required,len=11"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, 400, "参数错误")
		return
	}

	user, err := c.userService.Register(req.Username, req.Password, req.Email, req.Phone)
	if err != nil {
		response.Error(ctx, 500, "注册失败")
		return
	}

	response.Success(ctx, gin.H{"user": user})
}
