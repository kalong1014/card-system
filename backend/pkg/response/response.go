package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIResponse API统一响应结构
type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Success 返回成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, APIResponse{
		Code:    0,
		Message: "操作成功",
		Data:    data,
	})
}

// SuccessWithMessage 返回带消息的成功响应
func SuccessWithMessage(c *gin.Context, message string) {
	c.JSON(http.StatusOK, APIResponse{
		Code:    0,
		Message: message,
	})
}

// Error 返回错误响应
func Error(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, APIResponse{
		Code:    code,
		Message: message,
	})
}

// ErrorWithStatus 返回带HTTP状态码的错误响应
func ErrorWithStatus(c *gin.Context, httpStatus, code int, message string) {
	c.JSON(httpStatus, APIResponse{
		Code:    code,
		Message: message,
	})
}
