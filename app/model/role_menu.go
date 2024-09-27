package model

import (
	"fiber-api/pkg/datetime"

	"gorm.io/gorm"
)

// 角色菜单关系模型
type RoleMenu struct {
	Id         int               `gorm:"autoIncrement"`
	CreateTime datetime.Datetime `gorm:"autoCreateTime"`
	UpdateTime datetime.Datetime `gorm:"autoUpdateTime"`
	DeleteTime gorm.DeletedAt
	RoleId     int
	MenuId     int
}
