package openai

import (
	"github.com/sashabaranov/go-openai"
	"strings"
)

func Chat(systemContent string, userContent string) (string, error) {
	// 组合Prompt
	sysMsg := newSysMsg(systemContent)
	userMsg := newUserMsg(userContent)
	msgs := []openai.ChatCompletionMessage{sysMsg, userMsg}

	//fmt.Println("system:", systemContent)
	//fmt.Println("user:", userContent)

	// 请求模型
	resp, err := RequestMessages(msgs)
	if err != nil {
		return "", err
	}

	// 处理结果
	data := strings.TrimSpace(resp.Content)
	data = strings.Replace(data, "```json\n{", "{", 1)
	data = strings.Replace(data, "```\n{", "{", 1)
	data = strings.Replace(data, "}\n```", "}", 1)

	return data, nil
}
