package server

import (
	"github.com/gin-gonic/gin"
	"neko-acm-ai/internal/conf"
	"neko-acm-ai/internal/model"
	"neko-acm-ai/server/handler"
	"neko-acm-ai/server/middlewares"
	"net/http"
)

func InitRoute() error {
	config := conf.Conf.Server

	// index
	ginServer.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, model.RespOk("Neko ACM AI 启动成功！", nil))
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
	ginServer.POST("/solution", handler.SolutionDraft)

	// 启动服务
	err := ginServer.Run(":" + config.Port)
	if err != nil {
		return err
	}
	return nil
}
