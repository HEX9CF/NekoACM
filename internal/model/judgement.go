package model

type JudgeStatus uint64

// 评测结果
type Judgement struct {
	Stdout        string `json:"stdout,omitempty"`
	Stderr        string `json:"stderr,omitempty"`
	CompileOutput string `json:"compile_output,omitempty"`
	Message       string `json:"message,omitempty"`
	Status        string `json:"status"`
}
