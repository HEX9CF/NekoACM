package handler

import (
	"github.com/gin-gonic/gin"
	"neko-acm/internal/model"
	"neko-acm/internal/service/misc"
	"net/http"
)

// 生成笑话
func GenerateJoke(c *gin.Context) {
	p, err := misc.TellJoke()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", p))
}
