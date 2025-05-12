package rpc_test

import (
	"context"
	"nekoacm-server/pkg/pb"
	"testing"
	"time"
)

// 测试 Testcase 服务
func TestTestcaseService(t *testing.T) {
	// 建立连接
	conn := getConnection(t)
	defer conn.Close()

	// 创建客户端
	client := pb.NewTestcaseServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// 测试 Generate
	t.Run("Generate", func(t *testing.T) {
		req := &pb.TestcaseInstructionRequest{
			Title:        "两数之和",
			Description:  "给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。",
			Input:        "一个整数数组 nums 和一个目标值 target。",
			Output:       "返回两个整数的数组下标。",
			SampleInput:  "nums = [2, 7, 11, 15], target = 9",
			SampleOutput: "[0, 1]",
		}

		resp, err := client.Generate(ctx, req)
		if err != nil {
			t.Fatalf("Generate 调用失败: %v", err)
		}

		if resp == nil {
			t.Fatal("预期返回非空响应，得到了 nil")
		}

		if resp.TestInput == "" || resp.TestOutput == "" {
			t.Error("预期测试输入和输出非空，得到了空内容")
		}

		t.Logf("成功生成测试用例")
	})
}
