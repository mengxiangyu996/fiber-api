package model

import (
	"breeze-api/pkg/datetime"

	"gorm.io/gorm"
)

// 角色模型
type Role struct {
	Id         int               `gorm:"autoIncrement"`
	CreateTime datetime.DateTime `gorm:"autoCreateTime"`
	UpdateTime datetime.DateTime `gorm:"autoUpdateTime"`
	DeleteTime gorm.DeletedAt
	Name       string
	Status     int `gorm:"default:1"`
}
