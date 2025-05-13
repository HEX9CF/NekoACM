package openai

import "github.com/sashabaranov/go-openai"

// 创建系统消息
func newSysMsg(c string) openai.ChatCompletionMessage {
	return openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: c,
	}
}

// 创建用户消息
func newUserMsg(c string) openai.ChatCompletionMessage {
	return openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: c,
	}
}
