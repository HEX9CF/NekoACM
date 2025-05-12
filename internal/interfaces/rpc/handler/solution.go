package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"nekoacm-server/internal/application/converter"
	"nekoacm-server/internal/application/service"
	"nekoacm-server/pkg/pb"
)

// SolutionServer 实现 SolutionService 接口
type SolutionServer struct {
	pb.UnimplementedSolutionServiceServer
}

// Generate 生成题解的 gRPC 处理方法
func (s *SolutionServer) Generate(ctx context.Context, req *pb.SolutionInstructionRequest) (*pb.SolutionResponse, error) {
	// 将 protobuf 请求转换为 DTO
	instruction := converter.SolutionInstructionToDTO(req)

	// 调用服务生成题解
	solution, err := service.SolutionGenerate(*instruction)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "生成题解失败: %v", err)
	}

	// 将 DTO 转换为 protobuf 响应
	return converter.SolutionResponseFromDTO(&solution), nil
}
