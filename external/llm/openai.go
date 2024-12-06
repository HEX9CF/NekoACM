package llm

import (
	openai "github.com/sashabaranov/go-openai"
	"log"
	"stuoj-ai/internal/conf"
)

var (
	client *openai.Client
	config conf.OpenaiConf
)

func InitLlm() error {
	config = conf.Conf.Openai

	openaiConfig := openai.DefaultConfig(config.ApiKey)
	openaiConfig.BaseURL = config.BaseUrl
	log.Println("Connecting to AI LLM service: " + config.BaseUrl)

	client = openai.NewClientWithConfig(openaiConfig)
	log.Println("AI LLM service connected.")

	return nil
}
