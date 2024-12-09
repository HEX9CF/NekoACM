package bootstrap

import (
	"log"
	"neko-acm/external/llm"
)

func initLlm() error {
	err := llm.InitLlm()
	if err != nil {
		log.Println("Init LLM failed!")
		return err
	}

	log.Println("Init LLM success.")
	return nil
}
