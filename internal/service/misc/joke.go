package misc

import (
	"errors"
	"github.com/sashabaranov/go-openai"
	"log"
	"neko-acm/external/open_ai"
	"neko-acm/prompt"
)

// 生成笑话
func TellJoke() (string, error) {
	log.Println("请求生成笑话...")

	// 组合Prompt
	sysMsg := open_ai.NewSysMsg(prompt.JokeTeller)
	userMsg := open_ai.NewUserMsg(prompt.TellJoke)
	msgs := []openai.ChatCompletionMessage{sysMsg, userMsg}

	// 请求模型
	resp, err := open_ai.RequestMessages(msgs)
	if err != nil {
		log.Println(err)
		return "", errors.New("请求模型失败！")
	}
	log.Println("生成结果：" + resp.Content)

	return resp.Content, nil
}
