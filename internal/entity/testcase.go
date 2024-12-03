package entity

// 评测点数据
type Testcase struct {
	Id         uint64 `gorm:"primaryKey;autoIncrement;comment:评测点数据ID" json:"id,omitempty"`
	Serial     uint64 `gorm:"not null;default:0;comment:评测点序号" json:"serial,omitempty"`
	ProblemId  uint64 `gorm:"not null;default:0;comment:题目ID" json:"problem_id,omitempty"`
	TestInput  string `gorm:"type:longtext;not null;comment:测试输入" json:"test_input,omitempty"`
	TestOutput string `gorm:"type:longtext;not null;comment:测试输出" json:"test_output,omitempty"`
}

func (Testcase) TableName() string {
	return "tbl_testcase"
}
