package model

import (
	"time"

	"gorm.io/gorm"
)

// 管理员角色关系模型
type AdminRole struct {
	Id         int       `gorm:"autoIncrement"`
	CreateTime time.Time `gorm:"autoCreateTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime"`
	DeleteTime gorm.DeletedAt
	AdminId    int
	RoleId     int
}
