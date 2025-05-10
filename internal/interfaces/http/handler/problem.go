package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"neko-acm/internal/application/dto"
	"neko-acm/internal/application/service"
	"neko-acm/internal/interfaces/http/vo"
	"net/http"
)

// 生成题目
func GenerateProblem(c *gin.Context) {
	var req dto.ProblemInstruction

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, vo.RespError("参数错误", nil))
		return
	}

	// 生成题目
	p, err := service.ProblemGenerate(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, vo.RespOk("OK", p))
}

// 翻译题目
func TranslateProblem(c *gin.Context) {
	var req dto.TranslateInstruction

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, vo.RespError("参数错误", nil))
		return
	}

	// 翻译题目
	p, err := service.ProblemTranslate(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, vo.RespOk("OK", p))
}

// 解析题目
func ParseProblem(c *gin.Context) {
	var req dto.ProblemData

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, vo.RespError("参数错误", nil))
		return
	}

	// 解析题目
	p, err := service.ProblemParse(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, vo.RespOk("OK", p))
}
