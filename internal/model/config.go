package model

import (
	"time"

	"gorm.io/gorm"
)

// 配置模型
type Config struct {
	Id          int            `json:"id" gorm:"autoIncrement"`
	CreateTime  time.Time      `json:"createTime" gorm:"autoCreateTime"`
	UpdateTime  time.Time      `json:"updateTime" gorm:"autoUpdateTime"`
	DeleteTime  gorm.DeletedAt `json:"deleteTime"`
	GroupName   string         `json:"groupName"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Value       string         `json:"value"`
	Remark      string         `json:"remark"`
	Status      int            `json:"status" gorm:"default:1"`
}
