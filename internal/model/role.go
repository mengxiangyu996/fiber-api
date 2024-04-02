package model

import (
	"time"

	"gorm.io/gorm"
)

// 角色模型
type Role struct {
	Id         int       `gorm:"autoIncrement"`
	CreateTime time.Time `gorm:"autoCreateTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime"`
	DeleteTime gorm.DeletedAt
	Name       string
	Status     int `gorm:"default:1"`
}
