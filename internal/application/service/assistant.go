package service

import (
	"log"
	"nekoacm-server/internal/application/dto"
	"nekoacm-server/internal/infrastructure/openai"
	"nekoacm-server/prompt"
	"strings"
)

// 助手
func AssistantChat(msg dto.ChatMsg) (string, error) {
	msg.Content = strings.TrimSpace(msg.Content)
	if msg.Content == "" {
		return "", nil
	}
	log.Println("请求对话，内容长度:", len(msg.Content))

	// 请求模型
	resp, err := openai.Chat(prompt.ChatAssistant, msg.Content)
	if err != nil {
		return "", err
	}
	log.Println("请求成功，结果长度:", len(resp))

	return resp, nil
}
