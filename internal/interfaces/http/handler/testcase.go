package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"neko-acm/internal/model"
	"neko-acm/internal/service/testcase"
	"net/http"
)

// 生成测试用例
func GenerateTestcase(c *gin.Context) {
	var req model.TestcaseInstruction

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 生成测试用例
	p, err := testcase.Generate(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", p))
}
