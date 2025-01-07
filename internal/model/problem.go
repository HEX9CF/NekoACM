package model

// 题目
type Problem struct {
	Title        string   `json:"title,omitempty"`
	Description  string   `json:"description,omitempty"`
	Input        string   `json:"input,omitempty"`
	Output       string   `json:"output,omitempty"`
	SampleInput  string   `json:"sample_input,omitempty"`
	SampleOutput string   `json:"sample_output,omitempty"`
	Hint         string   `json:"hint,omitempty"`
	Tags         []string `json:"tags,omitempty"`
}

// 题目说明
type ProblemInstruction struct {
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

// 题目数据
type ProblemData struct {
	Data string `json:"data,omitempty" binding:"omitempty"`
}

// 翻译题目说明
type TranslateInstruction struct {
	Title       string   `json:"title,omitempty" binding:"omitempty"`
	Description string   `json:"description,omitempty" binding:"omitempty"`
	Input       string   `json:"input,omitempty" binding:"omitempty"`
	Output      string   `json:"output,omitempty" binding:"omitempty"`
	Hint        string   `json:"hint,omitempty" binding:"omitempty"`
	Tags        []string `json:"tags,omitempty" binding:"omitempty"`
	TargetLang  string   `json:"target_lang,omitempty" binding:"omitempty"`
}
