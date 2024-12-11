package bootstrap

import (
	"log"
	"neko-acm/external/llm"
)

// 初始化大模型服务
func initLlm() error {
	err := llm.InitLlm()
	if err != nil {
		log.Println("初始化大模型服务失败！")
		return err
	}

	log.Println("初始化大模型服务成功")
	return nil
}
