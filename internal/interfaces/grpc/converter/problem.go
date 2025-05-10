package converter

import (
	"neko-acm/internal/application/dto"
	"neko-acm/pkg/pb"
)

// ProblemResponseFromModel 将内部模型转换为 protobuf 响应
func ProblemResponseFromModel(p *dto.Problem) *pb.ProblemResponse {
	if p == nil {
		return nil
	}
	return &pb.ProblemResponse{
		Title:        p.Title,
		Description:  p.Description,
		Input:        p.Input,
		Output:       p.Output,
		SampleInput:  p.SampleInput,
		SampleOutput: p.SampleOutput,
		Hint:         p.Hint,
		Tags:         p.Tags,
	}
}

// ProblemInstructionToModel 将 protobuf 请求转换为内部模型
func ProblemInstructionToModel(req *pb.ProblemInstructionRequest) *dto.ProblemInstruction {
	if req == nil {
		return nil
	}
	return &dto.ProblemInstruction{
		Title:        req.GetTitle(),
		Description:  req.GetDescription(),
		Input:        req.GetInput(),
		Output:       req.GetOutput(),
		SampleInput:  req.GetSampleInput(),
		SampleOutput: req.GetSampleOutput(),
		Hint:         req.GetHint(),
		Tags:         req.GetTags(),
		Solution:     req.GetSolution(),
	}
}

// TranslateInstructionToModel 将 protobuf 翻译请求转换为内部模型
func TranslateInstructionToModel(req *pb.TranslateInstructionRequest) *dto.TranslateInstruction {
	if req == nil {
		return nil
	}
	return &dto.TranslateInstruction{
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Input:       req.GetInput(),
		Output:      req.GetOutput(),
		Hint:        req.GetHint(),
		Tags:        req.GetTags(),
		TargetLang:  req.GetTargetLang(),
	}
}

// ProblemDataToModel 将 protobuf 数据请求转换为内部模型
func ProblemDataToModel(req *pb.ProblemDataRequest) *dto.ProblemData {
	if req == nil {
		return nil
	}
	return &dto.ProblemData{
		Data: req.GetData(),
	}
}
