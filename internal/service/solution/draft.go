package solution

import (
	"encoding/json"
	"errors"
	"github.com/sashabaranov/go-openai"
	"log"
	"neko-acm/external/llm"
	"neko-acm/internal/model"
	"neko-acm/prompt"
	"neko-acm/utils"
)

func Draft(pi model.SolutionInstruction) (model.Solution, error) {
	var t model.Solution

	// 说明转换为字符串
	instruction, err := utils.PrettyStruct(pi)
	if err != nil {
		return model.Solution{}, err
	}
	log.Println("请求生成题解：" + instruction)

	// 组合Prompt
	sysMsg := llm.NewSysMsg(prompt.SolutionDraft)
	userMsg := llm.NewUserMsg(instruction)
	msgs := []openai.ChatCompletionMessage{sysMsg, userMsg}

	// 请求模型
	resp, err := llm.RequestMessages(msgs)
	if err != nil {
		return model.Solution{}, errors.New("请求模型失败！")
	}
	log.Println("生成结果：" + resp.Content)

	// 解析结果
	err = json.Unmarshal([]byte(resp.Content), &t)
	if err != nil {
		log.Println(err)
		return model.Solution{}, errors.New("解析结果失败，请重试！")
	}

	return t, nil
}
