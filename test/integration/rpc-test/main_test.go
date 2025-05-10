package rpc_test

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"neko-acm/internal/infrastructure/open_ai"
	"neko-acm/internal/interfaces/rpc"
	"neko-acm/pkg/config"
	"neko-acm/prompt"
	"os"
	"testing"
	"time"
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

	prompt.InitPrompt()

	go func() {
		err := rpc.InitServer()
		if err != nil {
			panic("gRPC服务器初始化失败：" + err.Error())
		}
	}()

	code := m.Run()

	os.Exit(code)
}

// 创建 gRPC 连接的辅助函数
func getConnection(t *testing.T) *grpc.ClientConn {
	conf := config.Conf.Grpc

	address := "localhost" + ":" + conf.Port

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("连接服务器失败: %v", err)
	}
	return conn
}
