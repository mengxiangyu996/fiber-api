package model

import (
	"time"

	"gorm.io/gorm"
)

// 角色权限关系模型
type RolePermission struct {
	Id           int       `gorm:"autoIncrement"`
	CreateTime   time.Time `gorm:"autoCreateTime"`
	UpdateTime   time.Time `gorm:"autoUpdateTime"`
	DeleteTime   gorm.DeletedAt
	RoleId       int
	PermissionId int
}
