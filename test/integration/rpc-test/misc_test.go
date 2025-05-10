package rpc_test

import (
	"context"
	"neko-acm/pkg/pb"
	"testing"
	"time"
)

// 测试 Misc 服务
func TestMiscService(t *testing.T) {
	// 建立连接
	conn := getConnection(t)
	defer conn.Close()

	// 创建客户端
	client := pb.NewMiscServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// 测试 TellJoke
	t.Run("TellJoke", func(t *testing.T) {
		req := &pb.Empty{}

		resp, err := client.TellJoke(ctx, req)
		if err != nil {
			t.Fatalf("TellJoke 调用失败: %v", err)
		}

		if resp == nil {
			t.Fatal("预期返回非空响应，得到了 nil")
		}

		if resp.Content == "" {
			t.Error("预期笑话内容非空，得到了空内容")
		}

		t.Logf("成功获取笑话: %s", resp.Content[:20]+"...")
	})
}
