package middlewares

import (
	"github.com/gin-gonic/gin"
	"neko-acm/internal/conf"
	"neko-acm/internal/model"
	"net/http"
)

func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token != "Bearer "+conf.Conf.Server.Token {
			c.JSON(http.StatusUnauthorized, model.RespError("Token错误", nil))
			c.Abort()
			return
		}

		// 放行
		c.Next()
	}
}
