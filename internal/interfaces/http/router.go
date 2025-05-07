package http

import (
	"github.com/gin-gonic/gin"
	handler2 "neko-acm/internal/interfaces/http/handler"
	"neko-acm/internal/interfaces/http/middlewares"
	"neko-acm/internal/model"
	"neko-acm/pkg/config"
	"net/http"
)

func InitRoute() error {
	config := config.Conf.Server

	// index
	ginServer.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, model.RespOk("NekoACM 启动成功", nil))
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
			c.JSON(http.StatusOK, model.RespOk("NekoACM 服务可用", nil))
		})
	}

	initProblemRoute(apiRoute)
	initTestcaseRoute(apiRoute)
	initSolutionRoute(apiRoute)
	initJudgeRoute(apiRoute)
	initChatRoute(apiRoute)
	initMiscRoute(apiRoute)

	// 启动服务
	err := ginServer.Run(":" + config.Port)
	if err != nil {
		return err
	}
	return nil
}

func initProblemRoute(rg *gin.RouterGroup) {
	pr := rg.Group("/problem")

	pr.POST("/parse", handler2.ParseProblem)
	pr.POST("/translate", handler2.TranslateProblem)
	pr.POST("/generate", handler2.GenerateProblem)
}

func initTestcaseRoute(rg *gin.RouterGroup) {
	tc := rg.Group("/testcase")

	tc.POST("/generate", handler2.GenerateTestcase)
}

func initSolutionRoute(rg *gin.RouterGroup) {
	s := rg.Group("/solution")

	s.POST("/generate", handler2.GenerateSolution)
}

func initJudgeRoute(rg *gin.RouterGroup) {
	j := rg.Group("/judge")

	j.POST("/submit", handler2.JudgeSubmit)
}

func initChatRoute(rg *gin.RouterGroup) {
	c := rg.Group("/chat")

	c.POST("/assistant", handler2.ChatAssistant)
}

func initMiscRoute(rg *gin.RouterGroup) {
	m := rg.Group("/misc")

	m.GET("/joke", handler2.GenerateJoke)
}
