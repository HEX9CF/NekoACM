package dto

// 题目
type Problem struct {
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Input        string   `json:"input"`
	Output       string   `json:"output"`
	SampleInput  string   `json:"sample_input"`
	SampleOutput string   `json:"sample_output"`
	Hint         string   `json:"hint"`
	Tags         []string `json:"tags"`
}

// 题目说明
type ProblemInstruction struct {
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

// 题目数据
type ProblemData struct {
	Data string `json:"data" binding:"omitempty"`
}

// 翻译题目说明
type TranslateInstruction struct {
	Title       string   `json:"title" binding:"omitempty"`
	Description string   `json:"description" binding:"omitempty"`
	Input       string   `json:"input" binding:"omitempty"`
	Output      string   `json:"output" binding:"omitempty"`
	Hint        string   `json:"hint" binding:"omitempty"`
	Tags        []string `json:"tags" binding:"omitempty"`
	TargetLang  string   `json:"target_lang" binding:"omitempty"`
}
