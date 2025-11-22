package openai

import (
	"log"
	"nekoacm/pkg/config"

	"github.com/sashabaranov/go-openai"
)

var (
	client *openai.Client
	conf   config.OpenaiConf
)

// 初始化大模型服务
func InitLlm() error {
	conf = config.Conf.Openai

	// 配置大模型服务
	openaiConfig := openai.DefaultConfig(conf.ApiKey)
	openaiConfig.BaseURL = conf.BaseUrl
	log.Println("正在连接 LLM 服务：" + conf.BaseUrl)

	// 创建客户端
	client = openai.NewClientWithConfig(openaiConfig)

	// 测试连接
	err := ping()
	if err != nil {
		log.Println("LLM 服务连接失败！")
		return err
	}

	log.Println("LLM 服务连接成功")
	return nil
}
