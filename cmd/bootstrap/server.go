package bootstrap

import (
	"log"
	"nekoacm-server/internal/interfaces/rpc"
)

// 初始化服务器
func InitServer() error {
	// 启动gRPC服务器
	log.Println("启动gRPC服务器")

	err := rpc.InitServer()
	if err != nil {
		log.Println("初始化gRPC服务器失败！")
		return err
	}

	return nil
}
