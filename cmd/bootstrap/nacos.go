package bootstrap

import (
	"log"
	"nekoacm/internal/infrastructure/nacos"
	"nekoacm/pkg/config"
)

func initNacos() error {
	if !config.Conf.Nacos.Enable {
		return nil
	}

	err := nacos.InitNacos()
	if err != nil {
		log.Println("初始化 Nacos 失败！")
		return err
	}
	log.Println("初始化 Nacos 成功")

	err = nacos.LoadConfig()
	if err != nil {
		log.Println("加载 Nacos 配置失败！")
		return err
	}
	log.Println("加载 Nacos 配置成功")

	return nil
}
