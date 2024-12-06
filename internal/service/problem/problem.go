package problem

import (
	"encoding/json"
	"errors"
	"github.com/sashabaranov/go-openai"
	"log"
	"stuoj-ai/external/llm"
	"stuoj-ai/internal/model"
	"stuoj-ai/prompt"
	"stuoj-ai/utils"
)

func Draft(pi model.ProblemInstruction) (model.Problem, error) {
	var p model.Problem

	// 题目说明转换为字符串
	instruction, err := utils.PrettyStruct(pi)
	if err != nil {
		return model.Problem{}, err
	}
	log.Println("生成题目：" + instruction)

	// 组合Prompt
	sysMsg := model.NewSysMsg(prompt.ProblemJsonInput +
		"\n\n" + prompt.ProblemJsonOutput)
	userMsg := model.NewUserMsg(instruction)
	msgs := []openai.ChatCompletionMessage{sysMsg, userMsg}

	// 请求模型
	resp, err := llm.RequestMessages(msgs)
	if err != nil {
		return model.Problem{}, errors.New("请求模型失败！")
	}

	// 解析结果
	err = json.Unmarshal([]byte(resp.Content), &p)

	return p, err
}
