package rpc_test

import (
	"context"
	"nekoacm-common/api/proto/pb"
	"testing"
	"time"
)

// 测试 Assistant 服务
func TestAssistantService(t *testing.T) {
	// 建立连接
	conn := getConnection(t)
	defer conn.Close()

	// 创建客户端
	client := pb.NewAssistantServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// 测试 Chat
	t.Run("Chat", func(t *testing.T) {
		req := &pb.ChatRequest{
			Content: "什么是动态规划？",
		}

		resp, err := client.Chat(ctx, req)
		if err != nil {
			t.Fatalf("Chat 调用失败: %v", err)
		}

		if resp == nil {
			t.Fatal("预期返回非空响应，得到了 nil")
		}

		if resp.Content == "" {
			t.Error("预期聊天内容非空，得到了空内容")
		}

		t.Logf("成功获取助手回复")
	})
}
