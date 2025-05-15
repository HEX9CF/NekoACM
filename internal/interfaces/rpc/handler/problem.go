package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"nekoacm-common/api/proto/pb"
	"nekoacm-server/internal/application/converter"
	"nekoacm-server/internal/application/service"
)

// implement UnimplementedProblemServiceServer
type ProblemServer struct {
	pb.UnimplementedProblemServiceServer
}

func (s *ProblemServer) Generate(ctx context.Context, req *pb.ProblemInstructionRequest) (*pb.ProblemResponse, error) {
	// 将proto请求转换为内部模型
	instruction := converter.ProblemInstructionToDTO(req)

	// 调用服务生成题目
	p, err := service.ProblemGenerate(*instruction)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "生成题目失败: %v", err)
	}

	// 将内部模型转换为proto响应
	return converter.ProblemResponseFromDTO(&p), nil
}

func (s *ProblemServer) Translate(ctx context.Context, req *pb.TranslateInstructionRequest) (*pb.ProblemResponse, error) {
	// 将proto请求转换为内部模型
	instruction := converter.TranslateInstructionToDTO(req)

	// 调用服务翻译题目
	p, err := service.ProblemTranslate(*instruction)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "翻译题目失败: %v", err)
	}

	// 将内部模型转换为proto响应
	return converter.ProblemResponseFromDTO(&p), nil
}

func (s *ProblemServer) Parse(ctx context.Context, req *pb.ProblemDataRequest) (*pb.ProblemResponse, error) {
	// 将proto请求转换为内部模型
	data := converter.ProblemDataToDTO(req)

	// 调用服务解析题目
	p, err := service.ProblemParse(*data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "解析题目失败: %v", err)
	}

	// 将内部模型转换为proto响应
	return converter.ProblemResponseFromDTO(&p), nil
}
