package entity

// 题解
type Solution struct {
	Id         uint64 `gorm:"primaryKey;autoIncrement;comment:题解ID" json:"id,omitempty"`
	ProblemId  uint64 `gorm:"not null;default:0;comment:题目ID" json:"problem_id,omitempty"`
	LanguageId uint64 `gorm:"not null;default:0;comment:语言ID" json:"language_id,omitempty"`
	SourceCode string `gorm:"type:longtext;not null;comment:源代码" json:"source_code,omitempty"`
}

func (Solution) TableName() string {
	return "tbl_solution"
}
