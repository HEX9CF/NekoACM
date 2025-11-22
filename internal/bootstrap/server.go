package bootstrap

import (
	"log"
	"nekoacm/internal/interfaces/http"
)

// 初始化服务器
func InitServer() error {
	// 启动HTTP服务器
	log.Println("正在启动 HTTP 服务器...")
	if err := http.InitServer(); err != nil {
		log.Println("初始化 HTTP 服务器失败！")
		return err
	}

	//log.Println("启动 gRPC 服务器")
	//if err := rpc.InitServer(); err != nil {
	//	log.Println("初始化 gRPC 服务器失败！")
	//	return err
	//}

	return nil
}
