package model

import (
	"fiber-api/pkg/datetime"

	"gorm.io/gorm"
)

// 角色权限关系模型
type RolePermission struct {
	Id           int               `gorm:"autoIncrement"`
	CreateTime   datetime.Datetime `gorm:"autoCreateTime"`
	UpdateTime   datetime.Datetime `gorm:"autoUpdateTime"`
	DeleteTime   gorm.DeletedAt
	RoleId       int
	PermissionId int
}
