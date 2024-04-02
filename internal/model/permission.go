package model

import (
	"time"

	"gorm.io/gorm"
)

// 权限模型
type Permission struct {
	Id         int       `gorm:"autoIncrement"`
	CreateTime time.Time `gorm:"autoCreateTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime"`
	DeleteTime gorm.DeletedAt
	Name       string
	GroupName  string
	Path       string
	Method     string
	Status     int `grom:"default:1"`
}
