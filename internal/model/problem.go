package model

type Difficulty uint8

type ProblemStatus uint8

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
