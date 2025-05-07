package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"neko-acm/internal/model"
	"neko-acm/internal/service/problem"
	"net/http"
)

// 生成题目
func GenerateProblem(c *gin.Context) {
	var req model.ProblemInstruction

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 生成题目
	p, err := problem.Generate(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", p))
}

// 翻译题目
func TranslateProblem(c *gin.Context) {
	var req model.TranslateInstruction

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 翻译题目
	p, err := problem.Translate(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", p))
}

// 解析题目
func ParseProblem(c *gin.Context) {
	var req model.ProblemData

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 解析题目
	p, err := problem.Parse(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", p))
}
