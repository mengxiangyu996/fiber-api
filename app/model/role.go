package model

import (
	"fiber-api/pkg/datetime"

	"gorm.io/gorm"
)

// 角色模型
type Role struct {
	Id         int               `gorm:"autoIncrement"`
	CreateTime datetime.Datetime `gorm:"autoCreateTime"`
	UpdateTime datetime.Datetime `gorm:"autoUpdateTime"`
	DeleteTime gorm.DeletedAt
	Name       string
	Status     int `gorm:"default:1"`
}
