package problem

import (
	"encoding/json"
	"errors"
	"github.com/sashabaranov/go-openai"
	"log"
	"neko-acm-ai/external/llm"
	"neko-acm-ai/internal/model"
	"neko-acm-ai/prompt"
	"neko-acm-ai/utils"
)

func Draft(pi model.ProblemInstruction) (model.Problem, error) {
	var p model.Problem

	// 题目说明转换为字符串
	instruction, err := utils.PrettyStruct(pi)
	if err != nil {
		return model.Problem{}, err
	}
	log.Println("请求生成题目：" + instruction)

	// 组合Prompt
	sysMsg := llm.NewSysMsg(prompt.ProblemDraft)
	userMsg := llm.NewUserMsg(instruction)
	msgs := []openai.ChatCompletionMessage{sysMsg, userMsg}

	// 请求模型
	resp, err := llm.RequestMessages(msgs)
	if err != nil {
		return model.Problem{}, errors.New("请求模型失败！")
	}
	log.Println("生成结果：" + resp.Content)

	// 解析结果
	err = json.Unmarshal([]byte(resp.Content), &p)
	if err != nil {
		log.Println(err)
		return model.Problem{}, errors.New("解析结果失败，请重试！")
	}

	return p, nil
}
