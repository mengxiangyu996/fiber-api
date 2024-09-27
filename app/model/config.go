package model

// 配置模型
type Config struct {
	BaseModel
	GroupName   string
	Name        string
	Description string
	Value       string
	Remark      string
	Status      int `gorm:"default:1"`
}
