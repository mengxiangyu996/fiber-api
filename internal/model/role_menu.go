package model

import (
	"time"

	"gorm.io/gorm"
)

// 角色菜单关系模型
type RoleMenu struct {
	Id         int       `gorm:"autoIncrement"`
	CreateTime time.Time `gorm:"autoCreateTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime"`
	DeleteTime gorm.DeletedAt
	RoleId     int
	MenuId     int
}
