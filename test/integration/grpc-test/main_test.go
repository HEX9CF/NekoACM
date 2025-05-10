package grpc_test

import (
	"neko-acm/internal/infrastructure/open_ai"
	"neko-acm/internal/interfaces/grpc"
	"neko-acm/pkg/config"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	var err error

	err = os.Chdir("../../..")
	if err != nil {
		panic("切换工作目录失败：" + err.Error())
	}

	err = config.InitConfig()
	if err != nil {
		panic("配置文件加载失败：" + err.Error())
	}

	err = open_ai.InitLlm()
	if err != nil {
		panic("Llm初始化失败：" + err.Error())
	}

	go func() {
		err := grpc.InitServer()
		if err != nil {
			panic("gRPC服务器初始化失败：" + err.Error())
		}
	}()

	code := m.Run()

	os.Exit(code)
}
