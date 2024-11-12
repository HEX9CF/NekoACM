package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stuoj-ai/internal/model"
)

func InitRoute() error {
	// index
	ginServer.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, model.RespOk("启动成功！", nil))
	})

	// 404
	ginServer.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, model.RespError("404 Not Found", nil))
	})

	// 初始化路由

	// 启动服务
	err := ginServer.Run(":8080")
	if err != nil {
		return err
	}

	return nil
}
