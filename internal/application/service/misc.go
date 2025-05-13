package service

import (
	"log"
	"nekoacm-server/internal/infrastructure/openai"
	"nekoacm-server/prompt"
)

// 生成笑话
func TellJoke() (string, error) {
	log.Println("请求生成笑话...")

	// 请求模型
	resp, err := openai.Chat(prompt.TellJoke, prompt.TellJokeUser)
	if err != nil {
		return "", err
	}
	log.Println("生成结果：" + resp)

	return resp, nil
}
