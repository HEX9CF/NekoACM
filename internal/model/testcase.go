package model

// 测试用例
type Testcase struct {
	TestInput         string `json:"test_input,omitempty"`
	TestOutput        string `json:"test_output,omitempty"`
	InputExplanation  string `json:"input_explanation,omitempty"`
	OutputExplanation string `json:"output_explanation,omitempty"`
}

// 测试用例说明
type TestcaseInstruction struct {
	Title        string   `json:"title,omitempty" binding:"omitempty"`
	Description  string   `json:"description,omitempty" binding:"omitempty"`
	Input        string   `json:"input,omitempty" binding:"omitempty"`
	Output       string   `json:"output,omitempty" binding:"omitempty"`
	SampleInput  string   `json:"sample_input,omitempty" binding:"omitempty"`
	SampleOutput string   `json:"sample_output,omitempty" binding:"omitempty"`
	Hint         string   `json:"hint,omitempty" binding:"omitempty"`
	Tags         []string `json:"tags,omitempty" binding:"omitempty"`
	Solution     string   `json:"solution,omitempty" binding:"omitempty"`
}
