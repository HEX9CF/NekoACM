package solution

import (
	"encoding/json"
	"errors"
	"github.com/sashabaranov/go-openai"
	"log"
	"neko-acm/external/open_ai"
	"neko-acm/internal/model"
	"neko-acm/prompt"
	"neko-acm/utils"
)

// 生成题解
func Generate(si model.SolutionInstruction) (model.Solution, error) {
	var s model.Solution

	// 说明转换为字符串
	instruction, err := utils.PrettyStruct(si)
	if err != nil {
		return model.Solution{}, err
	}
	log.Println("请求生成题解：" + instruction)

	// 组合Prompt
	sysMsg := open_ai.NewSysMsg(prompt.SolutionGenerate)
	userMsg := open_ai.NewUserMsg(instruction)
	msgs := []openai.ChatCompletionMessage{sysMsg, userMsg}

	// 请求模型
	resp, err := open_ai.RequestMessages(msgs)
	if err != nil {
		log.Println(err)
		return model.Solution{}, errors.New("请求模型失败！")
	}
	log.Println("生成结果：" + resp.Content)

	// 解析结果
	err = json.Unmarshal([]byte(resp.Content), &s)
	if err != nil {
		log.Println(err)
		return model.Solution{}, errors.New("解析结果失败，请重试！")
	}

	return s, nil
}
