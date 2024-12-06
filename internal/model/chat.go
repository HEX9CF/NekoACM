package model

type ProblemInstruction struct {
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Input        string   `json:"input"`
	Output       string   `json:"output"`
	SampleInput  string   `json:"sample_input"`
	SampleOutput string   `json:"sample_output"`
	Hint         string   `json:"hint"`
	Tags         []string `json:"tags"`
	Solution     string   `json:"solution"`
}
