package converter

import (
	"neko-acm/internal/application/dto"
	"neko-acm/pkg/pb"
)

// SubmissionRequestToDTO 将 protobuf 请求转换为 DTO
func SubmissionRequestToDTO(req *pb.SubmissionRequest) *dto.Submission {
	if req == nil {
		return nil
	}
	return &dto.Submission{
		SourceCode:     req.GetSourceCode(),
		Language:       req.GetLanguage(),
		Stdin:          req.GetStdin(),
		ExpectedOutput: req.GetExpectedOutput(),
	}
}

// JudgementResponseFromDTO 将 DTO 转换为 protobuf 响应
func JudgementResponseFromDTO(j *dto.Judgement) *pb.JudgementResponse {
	if j == nil {
		return nil
	}
	return &pb.JudgementResponse{
		Stdout:        j.Stdout,
		Stderr:        j.Stderr,
		CompileOutput: j.CompileOutput,
		Message:       j.Message,
		Status:        j.Status,
	}
}
