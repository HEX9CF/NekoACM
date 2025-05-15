package bootstrap

import (
	"log"
	"nekoacm-server/internal/infrastructure/nacos"
	"nekoacm-server/internal/interfaces/rpc"
	"nekoacm-server/pkg/config"
)

// 初始化服务器
func initServer() error {
	if config.Conf.Nacos.Enable {
		if err := nacos.NacosClient.Register(); err != nil {
			log.Println("注册服务到 Nacos 失败！")
			return err
		}
	}

	log.Println("启动 gRPC 服务器")
	if err := rpc.InitServer(); err != nil {
		log.Println("初始化 gRPC 服务器失败！")
		return err
	}

	return nil
}
