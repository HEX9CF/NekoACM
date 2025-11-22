package bootstrap

import (
	"log"
	"nekoacm/internal/interfaces/rpc"
)

// 初始化服务器
func initServer() error {
	log.Println("启动 gRPC 服务器")
	if err := rpc.InitServer(); err != nil {
		log.Println("初始化 gRPC 服务器失败！")
		return err
	}

	return nil
}
