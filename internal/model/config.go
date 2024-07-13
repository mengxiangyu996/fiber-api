package model

import (
	"breeze-api/pkg/datetime"

	"gorm.io/gorm"
)

// 配置模型
type Config struct {
	Id          int               `gorm:"autoIncrement"`
	CreateTime  datetime.Datetime `gorm:"autoCreateTime"`
	UpdateTime  datetime.Datetime `gorm:"autoUpdateTime"`
	DeleteTime  gorm.DeletedAt
	GroupName   string
	Name        string
	Description string
	Value       string
	Remark      string
	Status      int `gorm:"default:1"`
}
