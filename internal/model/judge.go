package model

type JudgeStatus uint64

// 提交信息
type Submission struct {
	SourceCode     string `json:"source_code"`
	Language       string `json:"language"`
	Stdin          string `json:"stdin,omitempty"`
	ExpectedOutput string `json:"expected_output,omitempty"`
}

// 评测结果
type Judgement struct {
	Stdout        string `json:"stdout,omitempty"`
	Stderr        string `json:"stderr,omitempty"`
	CompileOutput string `json:"compile_output,omitempty"`
	Message       string `json:"message,omitempty"`
	Status        string `json:"status"`
}
