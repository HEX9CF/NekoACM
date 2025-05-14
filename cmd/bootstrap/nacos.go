package bootstrap

import (
	"log"
	"nekoacm-server/internal/infrastructure/nacos"
	"nekoacm-server/pkg/config"
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

	err = nacos.Register()
	if err != nil {
		log.Println("注册 Nacos 实例失败！")
		return err
	}
	log.Println("注册 Nacos 实例成功")

	err = nacos.LoadConfig()
	if err != nil {
		log.Println("加载 Nacos 配置失败！")
		return err
	}
	log.Println("加载 Nacos 配置成功")

	return nil
}
