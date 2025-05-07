package bootstrap

import (
	"log"
	"neko-acm/pkg/config"
)

// 初始化配置
func initConfig() error {
	err := config.InitConfig()
	if err != nil {
		log.Println("初始化配置失败！")
		return err
	}
	log.Println("初始化配置成功")
	return nil
}
