package open_ai

import (
	"context"
	"github.com/sashabaranov/go-openai"
)

// 请求消息
func RequestMessage(msg openai.ChatCompletionMessage) (openai.ChatCompletionMessage, error) {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    conf.Model,
			Messages: []openai.ChatCompletionMessage{msg},
		},
	)

	if err != nil {
		return openai.ChatCompletionMessage{}, err
	}

	return resp.Choices[0].Message, nil
}

// 请求消息
func RequestMessages(msgs []openai.ChatCompletionMessage) (openai.ChatCompletionMessage, error) {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    conf.Model,
			Messages: msgs,
		},
	)

	if err != nil {
		return openai.ChatCompletionMessage{}, err
	}

	return resp.Choices[0].Message, nil
}
