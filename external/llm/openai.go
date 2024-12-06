package llm

import (
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	"log"
	"stuoj-ai/internal/conf"
)

var (
	Client *openai.Client
	config conf.OpenaiConf
)

func InitLlm() error {
	config = conf.Conf.Openai

	openaiConfig := openai.DefaultConfig(config.ApiKey)
	openaiConfig.BaseURL = config.BaseUrl
	log.Println("Connecting to AI LLM service: " + config.BaseUrl)

	Client = openai.NewClientWithConfig(openaiConfig)
	log.Println("AI LLM service connected.")

	return nil
}

func Test() {
	resp, err := Client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: config.Model,
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
