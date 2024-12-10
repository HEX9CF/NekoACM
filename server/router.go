package server

import (
	"github.com/gin-gonic/gin"
	"neko-acm/internal/conf"
	"neko-acm/internal/model"
	"neko-acm/server/handler"
	"neko-acm/server/middlewares"
	"net/http"
)

func InitRoute() error {
	config := conf.Conf.Server

	// index
	ginServer.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, model.RespOk("NekoACM 启动成功！", nil))
	})

	// 404
	ginServer.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, model.RespError("404 Not Found", nil))
	})

	// 使用中间件
	ginServer.Use(middlewares.TokenAuth())

	// 初始化路由
	apiRoute := ginServer.Group("/api")
	{
		apiRoute.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, model.RespOk("NekoACM 启动成功！", nil))
		})
		apiRoute.POST("/problem", handler.GenerateProblem)
		apiRoute.POST("/testcase", handler.GenerateTestcase)
		apiRoute.POST("/solution", handler.GenerateSolution)
	}

	// 启动服务
	err := ginServer.Run(":" + config.Port)
	if err != nil {
		return err
	}
	return nil
}
