package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"neko-acm/internal/application/dto"
	"neko-acm/internal/application/service"
	"neko-acm/internal/interfaces/http/vo"
	"net/http"
)

// 生成题解
func GenerateSolution(c *gin.Context) {
	var req dto.SolutionInstruction

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, vo.RespError("参数错误", nil))
		return
	}

	// 生成题解
	p, err := service.SolutionGenerate(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, vo.RespOk("OK", p))
}
