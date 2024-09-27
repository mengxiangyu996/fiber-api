package model

import (
	"fiber-api/pkg/datetime"

	"gorm.io/gorm"
)

// 菜单模型
type Menu struct {
	Id         int               `gorm:"autoIncrement"`
	CreateTime datetime.Datetime `gorm:"autoCreateTime"`
	UpdateTime datetime.Datetime `gorm:"autoUpdateTime"`
	DeleteTime gorm.DeletedAt
	ParentId   int `gorm:"default:0"`
	Name       string
	Type       int
	Sort       int `gorm:"default:0"`
	Path       string
	Component  string
	Icon       string
	Redirect   string
	Status     int `gorm:"default:1"`
}
