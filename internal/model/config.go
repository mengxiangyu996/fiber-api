package model

import (
	"time"

	"gorm.io/gorm"
)

// 配置模型
type Config struct {
	Id          int       `gorm:"autoIncrement"`
	CreateTime  time.Time `gorm:"autoCreateTime"`
	UpdateTime  time.Time `gorm:"autoUpdateTime"`
	DeleteTime  gorm.DeletedAt
	GroupName   string
	Name        string
	Description string
	Value       string
	Remark      string
	Status      int `gorm:"default:1"`
}
