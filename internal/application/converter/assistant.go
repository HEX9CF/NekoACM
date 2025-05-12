package converter

import (
	"nekoacm-server/internal/application/dto"
	"nekoacm-server/pkg/pb"
)

// ChatRequestToDTO 将 protobuf 请求转换为 DTO
func ChatRequestToDTO(req *pb.ChatRequest) *dto.ChatMsg {
	if req == nil {
		return nil
	}
	return &dto.ChatMsg{
		Content: req.GetContent(),
	}
}

// ChatResponseFromDTO 将 DTO 转换为 protobuf 响应
func ChatResponseFromDTO(msg *dto.ChatMsg) *pb.ChatResponse {
	if msg == nil {
		return nil
	}
	return &pb.ChatResponse{
		Content: msg.Content,
	}
}
