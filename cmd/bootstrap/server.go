package bootstrap

import (
	"log"
	"neko-acm/internal/interfaces/grpc"
)

// 初始化服务器
func InitServer() error {
	err := grpc.InitServer()
	if err != nil {
		log.Println("初始化服务器失败！")
		return err
	}

	log.Println("初始化服务器成功")
	return nil
}
