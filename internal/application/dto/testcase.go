package dto

// 测试用例
type Testcase struct {
	TestInput         string `json:"test_input"`
	TestOutput        string `json:"test_output"`
	InputExplanation  string `json:"input_explanation"`
	OutputExplanation string `json:"output_explanation"`
}

// 测试用例说明
type TestcaseInstruction struct {
	Title        string   `json:"title" binding:"omitempty"`
	Description  string   `json:"description" binding:"omitempty"`
	Input        string   `json:"input" binding:"omitempty"`
	Output       string   `json:"output" binding:"omitempty"`
	SampleInput  string   `json:"sample_input" binding:"omitempty"`
	SampleOutput string   `json:"sample_output" binding:"omitempty"`
	Hint         string   `json:"hint" binding:"omitempty"`
	Tags         []string `json:"tags" binding:"omitempty"`
	Solution     string   `json:"solution" binding:"omitempty"`
}
