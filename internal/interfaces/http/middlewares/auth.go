package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"neko-acm/internal/conf"
	"neko-acm/internal/model"
	"net/http"
)

// Token验证中间件
func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		// 验证Token
		if token != "Bearer "+conf.Conf.Server.Token {
			log.Println("Token 错误，拒绝请求")
			c.JSON(http.StatusUnauthorized, model.RespError("Token 错误，拒绝请求", nil))
			c.Abort()
			return
		}

		// 放行
		c.Next()
	}
}
