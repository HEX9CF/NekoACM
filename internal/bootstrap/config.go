package bootstrap

import (
	"log"
	"neko-acm-ai/internal/conf"
)

func initConfig() error {
	err := conf.InitConfig()
	if err != nil {
		log.Println("Init config failed!")
		return err
	}
	log.Println("Init config success.")
	return nil
}
