package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"nekoacm-server/internal/application/dto"
	"nekoacm-server/internal/application/service"
	"nekoacm-server/internal/interfaces/http/vo"
	"net/http"
)

// 生成测试用例
func GenerateTestcase(c *gin.Context) {
	var req dto.TestcaseInstruction

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, vo.RespError("参数错误", nil))
		return
	}

	// 生成测试用例
	p, err := service.TestcaseGenerate(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, vo.RespOk("OK", p))
}
