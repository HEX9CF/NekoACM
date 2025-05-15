package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"nekoacm-common/api/proto/pb"
	"nekoacm-server/internal/application/converter"
	"nekoacm-server/internal/application/service"
)

// JudgeServer 实现 JudgeService 接口
type JudgeServer struct {
	pb.UnimplementedJudgeServiceServer
}

// Submit 处理代码提交请求
func (s *JudgeServer) Submit(ctx context.Context, req *pb.SubmissionRequest) (*pb.JudgementResponse, error) {
	// 将 protobuf 请求转换为 DTO
	submission := converter.SubmissionRequestToDTO(req)

	// 调用评测服务
	judgement, err := service.Submit(*submission)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "代码评测失败: %v", err)
	}

	// 将结果转换为 protobuf 响应
	return converter.JudgementResponseFromDTO(&judgement), nil
}
