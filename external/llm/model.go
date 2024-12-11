package llm

import "github.com/sashabaranov/go-openai"

// 创建系统消息
func NewSysMsg(c string) openai.ChatCompletionMessage {
	return openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: c,
	}
}

// 创建用户消息
func NewUserMsg(c string) openai.ChatCompletionMessage {
	return openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: c,
	}
}
