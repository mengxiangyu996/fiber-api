package service

import (
	"breeze-api/api/model"
	"breeze-api/pkg/db"
)

// 角色数据服务
type Role struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

// 创建角色
func (*Role) Create(role *Role) error {
	return db.GormClient.Model(&model.Role{}).Create(&model.Role{
		Name:   role.Name,
		Status: role.Status,
	}).Error
}

// 更新角色
func (*Role) Update(role *Role) error {
	return db.GormClient.Model(&model.Role{}).Where("id = ?", role.Id).Updates(&model.Role{
		Name:   role.Name,
		Status: role.Status,
	}).Error
}

// 删除角色
func (*Role) Delete(id int) error {
	return db.GormClient.Model(&model.Role{}).Where("id = ?", id).Delete(nil).Error
}

// 角色列表
func (*Role) Page(page, size int) ([]*Role, int) {

	var (
		list  []*Role
		count int64
	)

	db.GormClient.Model(&model.Role{}).Order("id desc").Count(&count).Limit(size).Offset((page - 1) * size).Find(&list)

	return list, int(count)
}

// 角色详情
func (*Role) Detail(id int) *Role {

	var detail *Role

	db.GormClient.Model(&model.Role{}).Where("id = ?", id).Take(&detail)

	return detail
}

// 角色详情
func (*Role) DetailByName(name string) *Role {

	var detail *Role

	db.GormClient.Model(&model.Role{}).Where("name = ?", name).Take(&detail)

	return detail
}
