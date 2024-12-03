package entity

// 编程语言
type Language struct {
	Id   uint64 `gorm:"primaryKey;autoIncrement;comment:语言ID" json:"id,omitempty"`
	Name string `gorm:"type:varchar(255);not null;comment:语言名" json:"name,omitempty"`
}

func (Language) TableName() string {
	return "tbl_language"
}
