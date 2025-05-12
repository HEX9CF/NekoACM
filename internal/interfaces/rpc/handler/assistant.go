package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"nekoacm-server/internal/application/dto"
	"nekoacm-server/internal/application/service"
	"nekoacm-server/internal/interfaces/converter"
	"nekoacm-server/pkg/pb"
)

// AssistantServer 实现 AssistantService 接口
type AssistantServer struct {
	pb.UnimplementedAssistantServiceServer
}

// Chat 处理聊天请求
func (s *AssistantServer) Chat(ctx context.Context, req *pb.ChatRequest) (*pb.ChatResponse, error) {
	// 将 protobuf 请求转换为 DTO
	chatMsg := converter.ChatRequestToDTO(req)

	// 调用聊天服务
	response, err := service.AssistantChat(*chatMsg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "聊天请求失败: %v", err)
	}

	// 构建响应并返回
	return converter.ChatResponseFromDTO(&dto.ChatMsg{Content: response}), nil
}
