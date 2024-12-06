package model

// 评测点数据
type Testcase struct {
	Serial     uint64 `json:"serial,omitempty"`
	TestInput  string `json:"test_input,omitempty"`
	TestOutput string `json:"test_output,omitempty"`
}
