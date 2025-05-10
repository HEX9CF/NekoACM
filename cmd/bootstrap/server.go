package bootstrap

import (
	"log"
	"neko-acm/internal/interfaces/http"
	"neko-acm/internal/interfaces/rpc"
	"os"
	"os/signal"
	"syscall"
)

// 初始化服务器
func InitServer() error {
	// 创建错误通道
	errChan := make(chan error, 2)
	sigChan := make(chan os.Signal, 1)

	// 监听系统信号
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// 启动gRPC服务器
	go func() {
		log.Println("正在启动gRPC服务器...")
		err := rpc.InitServer()
		if err != nil {
			log.Println("初始化gRPC服务器失败！")
			errChan <- err
		}
	}()

	// 启动HTTP服务器
	go func() {
		log.Println("正在启动HTTP服务器...")
		err := http.InitServer()
		if err != nil {
			log.Println("初始化HTTP服务器失败！")
			errChan <- err
		}
	}()

	// 阻塞主线程，等待错误或终止信号
	select {
	case err := <-errChan:
		return err
	case sig := <-sigChan:
		log.Printf("接收到信号: %s，正在关闭服务器...\n", sig)
		return nil
	}
}
