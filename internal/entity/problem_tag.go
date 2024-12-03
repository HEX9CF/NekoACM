package entity

// 题目标签关系
type ProblemTag struct {
	ProblemId uint64 `gorm:"primaryKey;autoIncrement:false;comment:题目ID"`
	TagId     uint64 `gorm:"primaryKey;autoIncrement:false;comment:标签ID"`
}

func (ProblemTag) TableName() string {
	return "tbl_problem_tag"
}
