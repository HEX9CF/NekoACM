package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stuoj-ai/internal/model"
	"stuoj-ai/internal/service/problem"
)

func ProblemDraft(c *gin.Context) {
	var req model.ProblemInstruction

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	p, err := problem.Draft(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", p))
}
