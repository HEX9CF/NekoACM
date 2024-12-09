package llm

import (
	openai "github.com/sashabaranov/go-openai"
	"log"
	"neko-acm-ai/internal/conf"
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

	err := Test()
	if err != nil {
		log.Println("AI LLM service connection failed.")
		return err
	}

	log.Println("AI LLM service connected.")
	return nil
}
