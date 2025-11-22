package bootstrap

import (
	"log"
	"nekoacm/internal/interface/http"
	"nekoacm/internal/interface/rpc"
)

// 初始化服务器
func InitServer() error {
	// 启动HTTP服务器
	log.Println("正在启动 HTTP 服务器...")
	if err := http.InitServer(); err != nil {
		log.Println("初始化 HTTP 服务器失败！")
		return err
	}

	return nil
}

// 初始化gRPC服务器
func InitGrpcServer() error {
	log.Println("启动 gRPC 服务器")
	if err := rpc.InitServer(); err != nil {
		log.Println("初始化 gRPC 服务器失败！")
		return err
	}

	return nil
}
