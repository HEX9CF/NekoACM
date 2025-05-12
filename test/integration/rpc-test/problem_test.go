package rpc_test

import (
	"context"
	"nekoacm-server/pkg/pb"
	"testing"
	"time"
)

// 测试 Problem 服务
func TestProblemService(t *testing.T) {
	// 建立连接
	conn := getConnection(t)
	defer conn.Close()

	// 创建客户端
	client := pb.NewProblemServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// 测试 GenerateProblem
	t.Run("GenerateProblem", func(t *testing.T) {
		req := &pb.ProblemInstructionRequest{
			Title: "一个简单的数组求和问题",
			Tags:  []string{"数组", "数学"},
		}

		resp, err := client.GenerateProblem(ctx, req)
		if err != nil {
			t.Fatalf("GenerateProblem 调用失败: %v", err)
		}

		if resp == nil {
			t.Fatal("预期返回非空响应，得到了 nil")
		}

		if resp.Title == "" {
			t.Error("预期标题非空，得到了空标题")
		}

		t.Logf("成功生成题目: %s", resp.Title)
	})

	// 测试 TranslateProblem
	t.Run("TranslateProblem", func(t *testing.T) {
		req := &pb.TranslateInstructionRequest{
			Title:       "Two Sum",
			Description: "Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.",
			Input:       "An array of integers nums and an integer target.",
			Output:      "Return indices of the two numbers such that they add up to target.",
			TargetLang:  "zh-CN",
		}

		resp, err := client.TranslateProblem(ctx, req)
		if err != nil {
			t.Fatalf("TranslateProblem 调用失败: %v", err)
		}

		if resp == nil {
			t.Fatal("预期返回非空响应，得到了 nil")
		}

		// 检查是否为中文
		if resp.Title == req.Title {
			t.Error("翻译失败，标题未发生变化")
		}

		t.Logf("成功翻译题目: %s", resp.Title)
	})

	// 测试 ParseProblem
	t.Run("ParseProblem", func(t *testing.T) {
		req := &pb.ProblemDataRequest{
			Data: `# 两数之和
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。

## 输入
一个整数数组 nums 和一个目标值 target。

## 输出
返回两个整数的数组下标。

## 样例输入
nums = [2, 7, 11, 15], target = 9

## 样例输出
[0, 1]

## 提示
你可以假设每种输入只会对应一个答案。`,
		}

		resp, err := client.ParseProblem(ctx, req)
		if err != nil {
			t.Fatalf("ParseProblem 调用失败: %v", err)
		}

		if resp == nil {
			t.Fatal("预期返回非空响应，得到了 nil")
		}

		if resp.Title != "两数之和" {
			t.Errorf("预期标题为'两数之和'，得到了 '%s'", resp.Title)
		}

		t.Logf("成功解析题目: %s", resp.Title)
	})
}
