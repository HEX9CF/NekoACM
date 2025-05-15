package converter

import (
	"nekoacm-common/api/proto/pb"
	"nekoacm-server/internal/application/dto"
)

// SolutionResponseFromDTO 将 DTO 转换为 protobuf 响应
func SolutionResponseFromDTO(s *dto.Solution) *pb.SolutionResponse {
	if s == nil {
		return nil
	}
	return &pb.SolutionResponse{
		Language:    s.Language,
		SourceCode:  s.SourceCode,
		Explanation: s.Explanation,
	}
}

// SolutionInstructionToDTO 将 protobuf 请求转换为 DTO
func SolutionInstructionToDTO(req *pb.SolutionInstructionRequest) *dto.SolutionInstruction {
	if req == nil {
		return nil
	}
	return &dto.SolutionInstruction{
		Title:        req.GetTitle(),
		Description:  req.GetDescription(),
		Input:        req.GetInput(),
		Output:       req.GetOutput(),
		SampleInput:  req.GetSampleInput(),
		SampleOutput: req.GetSampleOutput(),
		Hint:         req.GetHint(),
		Tags:         req.GetTags(),
		Solution:     req.GetSolution(),
		Language:     req.GetLanguage(),
	}
}

// SolutionResponseToDTO 将 protobuf 响应转换为 DTO
func SolutionResponseToDTO(resp *pb.SolutionResponse) *dto.Solution {
	if resp == nil {
		return nil
	}
	return &dto.Solution{
		Language:    resp.GetLanguage(),
		SourceCode:  resp.GetSourceCode(),
		Explanation: resp.GetExplanation(),
	}
}
