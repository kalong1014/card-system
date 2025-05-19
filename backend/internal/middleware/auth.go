package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

// AuthMiddleware JWT认证中间件
func AuthMiddleware(redisClient *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取Authorization头
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "未提供认证信息"})
			return
		}

		// 验证Authorization头格式
		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "认证格式不正确"})
			return
		}

		// 解析token
		tokenString := authHeaderParts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 验证签名方法
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "无效的token"})
			return
		}

		// 验证token是否有效
		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "无效的token"})
			return
		}

		// 从token中获取用户ID
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "无法解析token claims"})
			return
		}

		userID, ok := claims["user_id"].(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "无效的用户ID"})
			return
		}

		// 检查token是否在Redis黑名单中
		ctx := c.Request.Context()
		if redisClient != nil {
			blacklisted, err := redisClient.Get(ctx, "token_blacklist:"+tokenString).Bool()
			if err == nil && blacklisted {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "已失效的token"})
				return
			}
		}

		// 将用户ID添加到上下文
		c.Set("user_id", userID)

		// 如果是商户请求，验证商户ID
		if role, exists := claims["role"].(string); exists && role == "merchant" {
			merchantID, ok := claims["merchant_id"].(float64)
			if !ok {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "无效的商户ID"})
				return
			}
			c.Set("merchant_id", uint(merchantID))
		}

		c.Next()
	}
}
