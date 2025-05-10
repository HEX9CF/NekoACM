package converter

import (
	"neko-acm/internal/application/dto"
	"neko-acm/pkg/pb"
)

// TestcaseResponseFromDTO 将 DTO 转换为 protobuf 响应
func TestcaseResponseFromDTO(t *dto.Testcase) *pb.TestcaseResponse {
	if t == nil {
		return nil
	}
	return &pb.TestcaseResponse{
		TestInput:         t.TestInput,
		TestOutput:        t.TestOutput,
		InputExplanation:  t.InputExplanation,
		OutputExplanation: t.OutputExplanation,
	}
}

// TestcaseInstructionToDTO 将 protobuf 请求转换为 DTO
func TestcaseInstructionToDTO(req *pb.TestcaseInstructionRequest) *dto.TestcaseInstruction {
	if req == nil {
		return nil
	}
	return &dto.TestcaseInstruction{
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

// TestcaseResponseToDTO 将 protobuf 响应转换为 DTO
func TestcaseResponseToDTO(resp *pb.TestcaseResponse) *dto.Testcase {
	if resp == nil {
		return nil
	}
	return &dto.Testcase{
		TestInput:         resp.GetTestInput(),
		TestOutput:        resp.GetTestOutput(),
		InputExplanation:  resp.GetInputExplanation(),
		OutputExplanation: resp.GetOutputExplanation(),
	}
}
