package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"nekoacm-server/internal/application/converter"
	"nekoacm-server/internal/application/service"
	"nekoacm-server/pkg/pb"
)

// TestcaseServer 实现 TestcaseService 接口
type TestcaseServer struct {
	pb.UnimplementedTestcaseServiceServer
}

// Generate 生成测试用例的 gRPC 处理方法
func (s *TestcaseServer) Generate(ctx context.Context, req *pb.TestcaseInstructionRequest) (*pb.TestcaseResponse, error) {
	// 将 protobuf 请求转换为 DTO
	instruction := converter.TestcaseInstructionToDTO(req)

	// 调用服务生成测试用例
	testcase, err := service.TestcaseGenerate(*instruction)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "生成测试用例失败: %v", err)
	}

	// 将 DTO 转换为 protobuf 响应
	return converter.TestcaseResponseFromDTO(&testcase), nil
}
