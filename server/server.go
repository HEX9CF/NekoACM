package server

import "github.com/gin-gonic/gin"

var (
	ginServer *gin.Engine
)

func InitServer() error {
	ginServer = gin.Default()
	err := InitRoute()
	if err != nil {
		return err
	}

	return nil
}
