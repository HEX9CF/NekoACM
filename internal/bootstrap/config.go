package bootstrap

import (
	"log"
	"neko-acm/internal/conf"
)

func initConfig() error {
	err := conf.InitConfig()
	if err != nil {
		log.Println("初始化配置失败！")
		return err
	}
	log.Println("初始化配置成功")
	return nil
}
