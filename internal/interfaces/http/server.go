package http

import (
	"github.com/gin-gonic/gin"
)

var (
	ginServer *gin.Engine
)

// 初始化服务器
func InitServer() error {
	ginServer = gin.Default()
	err := InitRoute()
	if err != nil {
		return err
	}

	return nil
}
