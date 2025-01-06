package model

// 提交信息
type Submission struct {
	SourceCode     string `json:"source_code"`
	Language       string `json:"language"`
	Stdin          string `json:"stdin,omitempty"`
	ExpectedOutput string `json:"expected_output,omitempty"`
}
