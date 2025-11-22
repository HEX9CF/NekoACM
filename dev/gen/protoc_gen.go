//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func main() {
	// 获取项目根目录（相对于当前脚本）
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Fprintf(os.Stderr, "无法获取当前文件路径\n")
		os.Exit(1)
	}

	rootDir := filepath.Join(filepath.Dir(filename), "../..")
	protoDir := filepath.Join(rootDir, "api/proto")
	outputDir := filepath.Join(rootDir, "api/proto/pb")

	// 切换到 proto 文件目录
	if err := os.Chdir(protoDir); err != nil {
		fmt.Fprintf(os.Stderr, "无法切换到proto目录 %s: %v\n", protoDir, err)
		os.Exit(1)
	}
	fmt.Printf("处理目录: %s\n", protoDir)

	// 获取当前目录下所有 proto 文件
	protoFiles, err := filepath.Glob("*.proto")
	if err != nil {
		fmt.Fprintf(os.Stderr, "查找 proto 文件失败: %v\n", err)
		os.Exit(1)
	}

	if len(protoFiles) == 0 {
		fmt.Println("警告: 未找到 proto 文件")
		return
	}
	fmt.Println("共找到 " + fmt.Sprint(len(protoFiles)) + " 个 proto 文件")

	// 确保输出目录存在
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "创建输出目录失败: %v\n", err)
		os.Exit(1)
	}

	// 生成 Go 代码
	args := append([]string{
		"--go_out=" + outputDir,
		"--go_opt=paths=source_relative",
	}, protoFiles...)

	cmd := exec.Command("protoc", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "生成 Go 代码失败: %v\n", err)
		os.Exit(1)
	}

	// 生成 gRPC 代码
	args = append([]string{
		"--go-grpc_out=" + outputDir,
		"--go-grpc_opt=paths=source_relative",
	}, protoFiles...)

	cmd = exec.Command("protoc", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "生成 gRPC 代码失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Protocol Buffers 代码生成成功")
}
