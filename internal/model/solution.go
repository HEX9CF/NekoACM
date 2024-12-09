package model

// 题解
type Solution struct {
	Language    string `json:"language,omitempty"`
	SourceCode  string `json:"source_code,omitempty"`
	Explanation string `json:"explanation,omitempty"`
}

// 题解说明
type SolutionInstruction struct {
	Title        string   `json:"title,omitempty" binding:"omitempty"`
	Description  string   `json:"description,omitempty" binding:"omitempty"`
	Input        string   `json:"input,omitempty" binding:"omitempty"`
	Output       string   `json:"output,omitempty" binding:"omitempty"`
	SampleInput  string   `json:"sample_input,omitempty" binding:"omitempty"`
	SampleOutput string   `json:"sample_output,omitempty" binding:"omitempty"`
	Hint         string   `json:"hint,omitempty" binding:"omitempty"`
	Tags         []string `json:"tags,omitempty" binding:"omitempty"`
	Solution     string   `json:"solution,omitempty" binding:"omitempty"`
	Language     string   `json:"language,omitempty" binding:"omitempty"`
}
