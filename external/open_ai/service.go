package open_ai

import (
	"github.com/sashabaranov/go-openai"
)

func Chat(systemContent string, userContent string) (string, error) {
	// 组合Prompt
	sysMsg := NewSysMsg(systemContent)
	userMsg := NewUserMsg(userContent)
	msgs := []openai.ChatCompletionMessage{sysMsg, userMsg}

	// 请求模型
	resp, err := RequestMessages(msgs)
	if err != nil {
		return "", err
	}

	return resp.Content, nil
}
