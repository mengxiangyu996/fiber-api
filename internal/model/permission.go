package model

import (
	"breeze-api/pkg/datetime"

	"gorm.io/gorm"
)

// 权限模型
type Permission struct {
	Id         int               `gorm:"autoIncrement"`
	CreateTime datetime.DateTime `gorm:"autoCreateTime"`
	UpdateTime datetime.DateTime `gorm:"autoUpdateTime"`
	DeleteTime gorm.DeletedAt
	Name       string
	GroupName  string
	Path       string
	Method     string
	Status     int `grom:"default:1"`
}
