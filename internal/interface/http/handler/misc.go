package handler

import (
	"nekoacm/internal/application/service"
	"nekoacm/internal/interface/http/vo"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 生成笑话
func GenerateJoke(c *gin.Context) {
	p, err := service.TellJoke()
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, vo.RespOk("OK", p))
}
