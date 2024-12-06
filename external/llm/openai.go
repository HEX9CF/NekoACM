package llm

import (
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	"stuoj-ai/internal/conf"
)

func Test() {
	config := conf.Conf.Ai

	openaiConfig := openai.DefaultConfig(config.Key)
	if config.Port == "" {
		openaiConfig.BaseURL = config.Host
	} else {
		openaiConfig.BaseURL = config.Host + ":" + config.Port
	}

	client := openai.NewClientWithConfig(openaiConfig)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello!",
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
