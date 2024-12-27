package testcase

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

// 生成测试用例
func Generate(ti model.TestcaseInstruction) (model.Testcase, error) {
	var t model.Testcase

	// 说明转换为字符串
	instruction, err := utils.PrettyStruct(ti)
	if err != nil {
		return model.Testcase{}, err
	}
	log.Println("请求生成测试用例：" + instruction)

	// 组合Prompt
	sysMsg := open_ai.NewSysMsg(prompt.TestcaseGenerate)
	userMsg := open_ai.NewUserMsg(instruction)
	msgs := []openai.ChatCompletionMessage{sysMsg, userMsg}

	// 请求模型
	resp, err := open_ai.RequestMessages(msgs)
	if err != nil {
		log.Println(err)
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
