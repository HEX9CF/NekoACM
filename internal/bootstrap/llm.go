package bootstrap

import (
	"log"
	"stuoj-ai/external/llm"
)

func initLlm() {
	err := llm.InitLlm()
	if err != nil {
		log.Println("Init LLM failed!")
		return
	}

	log.Println("Init LLM success.")
}
