package converter

import (
	"nekoacm-server/pkg/pb"
)

// JokeResponseFromContent 将字符串内容转换为 protobuf 响应
func JokeResponseFromContent(content string) *pb.JokeResponse {
	return &pb.JokeResponse{
		Content: content,
	}
}
