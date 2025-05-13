package bootstrap

import (
	"log"
	"nekoacm-server/internal/infrastructure/openai"
	"nekoacm-server/prompt"
)

// 初始化大模型服务
func initLlm() error {
	prompt.InitPrompt()
	log.Println("NekoACM Prompt 已加载")

	err := openai.InitLlm()
	if err != nil {
		log.Println("初始化 LLM 服务失败！")
		return err
	}

	log.Println("初始化 LLM 服务成功")
	return nil
}
