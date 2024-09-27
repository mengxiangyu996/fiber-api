package model

// 角色模型
type Role struct {
	BaseModel
	Name   string
	Status int `gorm:"default:1"`
}
