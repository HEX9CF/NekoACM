package testcase

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

func Draft(pi model.TestcaseInstruction) (model.Testcase, error) {
	var t model.Testcase

	// 说明转换为字符串
	instruction, err := utils.PrettyStruct(pi)
	if err != nil {
		return model.Testcase{}, err
	}
	log.Println("请求生成测试样例：" + instruction)

	// 组合Prompt
	sysMsg := llm.NewSysMsg(prompt.TestcaseDraft)
	userMsg := llm.NewUserMsg(instruction)
	msgs := []openai.ChatCompletionMessage{sysMsg, userMsg}

	// 请求模型
	resp, err := llm.RequestMessages(msgs)
	if err != nil {
		return model.Testcase{}, errors.New("请求模型失败！")
	}
	log.Println("生成结果：" + resp.Content)

	// 解析结果
	err = json.Unmarshal([]byte(resp.Content), &t)
	if err != nil {
		log.Println(err)
		return model.Testcase{}, errors.New("解析结果失败，请重试！")
	}

	return t, nil
}
