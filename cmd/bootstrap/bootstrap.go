package bootstrap

import (
	"log"
	"nekoacm-server/prompt"
	"os"
)

// 初始化
func Init() error {
	log.SetOutput(os.Stdout)

	if err := initConfig(); err != nil {
		return err
	}

	prompt.InitPrompt()

	if err := initLlm(); err != nil {
		return err
	}

	if err := InitServer(); err != nil {
		return err
	}

	if err := initNacos(); err != nil {
		return nil
	}

	return nil
}
