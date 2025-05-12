package rpc_test

import (
	"context"
	"nekoacm-server/pkg/pb"
	"testing"
	"time"
)

// 测试 Judge 服务
func TestJudgeService(t *testing.T) {
	// 建立连接
	conn := getConnection(t)
	defer conn.Close()

	// 创建客户端
	client := pb.NewJudgeServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// 测试 SubmitCode
	t.Run("SubmitCode", func(t *testing.T) {
		req := &pb.SubmissionRequest{
			SourceCode: `
package main
import "fmt"
func main() {
    var a, b int
    fmt.Scanf("%d %d", &a, &b)
    fmt.Println(a + b)
}`,
			Language:       "go",
			Stdin:          "1 2",
			ExpectedOutput: "3",
		}

		resp, err := client.SubmitCode(ctx, req)
		if err != nil {
			t.Fatalf("SubmitCode 调用失败: %v", err)
		}

		if resp == nil {
			t.Fatal("预期返回非空响应，得到了 nil")
		}

		t.Logf("代码评测状态: %s", resp.Status)
	})
}
