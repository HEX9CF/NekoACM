package llm

import (
	openai "github.com/sashabaranov/go-openai"
	"log"
	"neko-acm/internal/conf"
)

var (
	client *openai.Client
	config conf.OpenaiConf
)

func InitLlm() error {
	config = conf.Conf.Openai

	openaiConfig := openai.DefaultConfig(config.ApiKey)
	openaiConfig.BaseURL = config.BaseUrl
	log.Println("正在连接大模型服务：" + config.BaseUrl)

	client = openai.NewClientWithConfig(openaiConfig)

	err := Test()
	if err != nil {
		log.Println("大模型服务连接失败！")
		return err
	}

	log.Println("大模型服务连接成功")
	return nil
}
