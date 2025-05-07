package open_ai

import (
	"github.com/sashabaranov/go-openai"
	"log"
	"neko-acm/internal/conf"
)

var (
	client *openai.Client
	config conf.OpenaiConf
)

// 初始化大模型服务
func InitLlm() error {
	config = conf.Conf.Openai

	// 配置大模型服务
	openaiConfig := openai.DefaultConfig(config.ApiKey)
	openaiConfig.BaseURL = config.BaseUrl
	log.Println("正在连接大模型服务：" + config.BaseUrl)

	// 创建客户端
	client = openai.NewClientWithConfig(openaiConfig)

	// 测试连接
	err := Test()
	if err != nil {
		log.Println("大模型服务连接失败！")
		return err
	}

	log.Println("大模型服务连接成功")
	return nil
}
