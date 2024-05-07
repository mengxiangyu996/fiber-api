package model

import (
	"breeze-api/pkg/datetime"

	"gorm.io/gorm"
)

// 管理员模型
type Admin struct {
	Id         int           `gorm:"autoIncrement"`
	CreateTime datetime.Time `gorm:"autoCreateTime"`
	UpdateTime datetime.Time `gorm:"autoUpdateTime"`
	DeleteTime gorm.DeletedAt
	Username   string
	Nickname   string
	Gender     int
	Email      string
	Phone      string
	Password   string
	Avatar     string
	Status     int `gorm:"default:1"`
}
