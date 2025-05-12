package bootstrap

import (
	"log"
	"nekoacm-server/internal/infrastructure/open_ai"
)

// 初始化大模型服务
func initLlm() error {
	err := open_ai.InitLlm()
	if err != nil {
		log.Println("初始化大模型服务失败！")
		return err
	}

	log.Println("初始化大模型服务成功")
	return nil
}
