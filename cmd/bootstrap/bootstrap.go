package bootstrap

import (
	"log"
	"neko-acm/prompt"
	"os"
)

// 初始化
func Init() error {
	log.SetOutput(os.Stdout)
	if err := initConfig(); err != nil {
		return err
	}
	if err := initLlm(); err != nil {
		return err
	}
	prompt.InitPrompt()
	return nil
}
