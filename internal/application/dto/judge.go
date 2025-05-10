package dto

// 提交信息
type Submission struct {
	SourceCode     string `json:"source_code"`
	Language       string `json:"language"`
	Stdin          string `json:"stdin"`
	ExpectedOutput string `json:"expected_output"`
}

// 评测结果
type Judgement struct {
	Stdout        string `json:"stdout"`
	Stderr        string `json:"stderr"`
	CompileOutput string `json:"compile_output"`
	Message       string `json:"message"`
	Status        string `json:"status"`
}
