package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stuoj-ai/internal/conf"
	"stuoj-ai/internal/model"
	"stuoj-ai/server/handler"
	"stuoj-ai/server/middlewares"
)

func InitRoute() error {
	config := conf.Conf.Server

	// index
	ginServer.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, model.RespOk("STUOJ AI出题模块启动成功！", nil))
	})

	// 404
	ginServer.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, model.RespError("404 Not Found", nil))
	})

	// 使用中间件
	ginServer.Use(middlewares.TokenAuth())

	// 初始化路由
	ginServer.POST("/problem", handler.ProblemDraft)
	ginServer.POST("/testcase", handler.TestcaseDraft)

	// 启动服务
	err := ginServer.Run(":" + config.Port)
	if err != nil {
		return err
	}
	return nil
}
