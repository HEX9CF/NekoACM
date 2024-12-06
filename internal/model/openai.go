package model

import "github.com/sashabaranov/go-openai"

func NewSysMsg(c string) openai.ChatCompletionMessage {
	return openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: c,
	}
}

func NewUserMsg(c string) openai.ChatCompletionMessage {
	return openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: c,
	}
}
