package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"neko-acm/internal/application/service"
	"neko-acm/internal/interfaces/converter"
	"neko-acm/pkg/pb"
)

// MiscServer 实现 MiscService 接口
type MiscServer struct {
	pb.UnimplementedMiscServiceServer
}

// TellJoke 处理笑话请求
func (s *MiscServer) TellJoke(ctx context.Context, req *pb.Empty) (*pb.JokeResponse, error) {
	var response string
	var err error

	response, err = service.TellJoke()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "生成笑话失败: %v", err)
	}

	// 将结果转换为 protobuf 响应
	return converter.JokeResponseFromContent(response), nil
}
