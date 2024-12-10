package bootstrap

import (
	"log"
	"neko-acm/external/llm"
)

func initLlm() error {
	err := llm.InitLlm()
	if err != nil {
		log.Println("初始化大模型失败！")
		return err
	}

	log.Println("初始化大模型成功")
	return nil
}
