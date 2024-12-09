package model

// 测试用例
type Testcase struct {
	TestInput         string `json:"test_input,omitempty"`
	TestOutput        string `json:"test_output,omitempty"`
	InputExplanation  string `json:"input_explanation,omitempty"`
	OutputExplanation string `json:"output_explanation,omitempty"`
}
