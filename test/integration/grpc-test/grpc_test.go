package grpc_test

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"neko-acm/pkg/config"
	"neko-acm/pkg/pb"
	"testing"
	"time"
)

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

// 测试 Problem 服务
func TestProblemService(t *testing.T) {
	// 建立连接
	conn := getConnection(t)
	defer conn.Close()

	// 创建客户端
	client := pb.NewProblemServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
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

// 测试 Testcase 服务
func TestTestcaseService(t *testing.T) {
	// 建立连接
	conn := getConnection(t)
	defer conn.Close()

	// 创建客户端
	client := pb.NewTestcaseServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 测试 GenerateTestcase
	t.Run("GenerateTestcase", func(t *testing.T) {
		req := &pb.TestcaseInstructionRequest{
			Title:        "两数之和",
			Description:  "给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。",
			Input:        "一个整数数组 nums 和一个目标值 target。",
			Output:       "返回两个整数的数组下标。",
			SampleInput:  "nums = [2, 7, 11, 15], target = 9",
			SampleOutput: "[0, 1]",
		}

		resp, err := client.GenerateTestcase(ctx, req)
		if err != nil {
			t.Fatalf("GenerateTestcase 调用失败: %v", err)
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

// 测试 Solution 服务
func TestSolutionService(t *testing.T) {
	// 建立连接
	conn := getConnection(t)
	defer conn.Close()

	// 创建客户端
	client := pb.NewSolutionServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 测试 GenerateSolution
	t.Run("GenerateSolution", func(t *testing.T) {
		req := &pb.SolutionInstructionRequest{
			Title:        "两数之和",
			Description:  "给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。",
			Input:        "一个整数数组 nums 和一个目标值 target。",
			Output:       "返回两个整数的数组下标。",
			SampleInput:  "nums = [2, 7, 11, 15], target = 9",
			SampleOutput: "[0, 1]",
			Language:     "go",
		}

		resp, err := client.GenerateSolution(ctx, req)
		if err != nil {
			t.Fatalf("GenerateSolution 调用失败: %v", err)
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

// 测试 Judge 服务
func TestJudgeService(t *testing.T) {
	// 建立连接
	conn := getConnection(t)
	defer conn.Close()

	// 创建客户端
	client := pb.NewJudgeServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
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

// 测试 Assistant 服务
func TestAssistantService(t *testing.T) {
	// 建立连接
	conn := getConnection(t)
	defer conn.Close()

	// 创建客户端
	client := pb.NewAssistantServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
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

// 测试 Misc 服务
func TestMiscService(t *testing.T) {
	// 建立连接
	conn := getConnection(t)
	defer conn.Close()

	// 创建客户端
	client := pb.NewMiscServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
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
