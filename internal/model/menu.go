package model

import (
	"time"

	"gorm.io/gorm"
)

// 菜单模型
type Menu struct {
	Id         int            `json:"id" gorm:"autoIncrement"`
	CreateTime time.Time      `json:"createTime" gorm:"autoCreateTime"`
	UpdateTime time.Time      `json:"updateTime" gorm:"autoUpdateTime"`
	DeleteTime gorm.DeletedAt `json:"deleteTime"`
	ParentId   int            `json:"parentId" gorm:"default:0"`
	Name       string         `json:"name"`
	Type       int            `json:"type"`
	Sort       int            `json:"sort" gorm:"default:0"`
	Path       string         `json:"path"`
	Component  string         `json:"component"`
	Icon       string         `json:"icon"`
	Redirect   string         `json:"redirect"`
	Status     int         	  `json:"status" gorm:"default:1"`
}
