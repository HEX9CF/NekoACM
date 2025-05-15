package rpc_test

import (
	"context"
	"nekoacm-common/api/proto/pb"
	"testing"
	"time"
)

// 测试 Solution 服务
func TestSolutionService(t *testing.T) {
	// 建立连接
	conn := getConnection(t)
	defer conn.Close()

	// 创建客户端
	client := pb.NewSolutionServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// 测试 Generate
	t.Run("Generate", func(t *testing.T) {
		req := &pb.SolutionInstructionRequest{
			Title:        "两数之和",
			Description:  "给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。",
			Input:        "一个整数数组 nums 和一个目标值 target。",
			Output:       "返回两个整数的数组下标。",
			SampleInput:  "nums = [2, 7, 11, 15], target = 9",
			SampleOutput: "[0, 1]",
			Language:     "go",
		}

		resp, err := client.Generate(ctx, req)
		if err != nil {
			t.Fatalf("Generate 调用失败: %v", err)
		}

		if resp == nil {
			t.Fatal("预期返回非空响应，得到了 nil")
		}

		if resp.SourceCode == "" || resp.Language == "" {
			t.Error("预期源代码和语言非空，得到了空内容")
		}

		t.Logf("成功生成题解，语言为: %s", resp.Language)
	})
}
