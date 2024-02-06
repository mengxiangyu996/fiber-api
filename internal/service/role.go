package service

import (
	"breeze-api/internal/model"
	"breeze-api/pkg/db"
)

// 角色数据服务
type Role struct{}

// 创建角色
func (*Role) Create(role *model.Role) error {
	return db.GormClient.Model(&model.Role{}).Create(&role).Error
}

// 更新角色
func (*Role) Update(role *model.Role) error {
	return db.GormClient.Model(&model.Role{}).Where("id = ?", role.Id).Updates(&role).Error
}

// 删除角色
func (*Role) Delete(id int) error {
	return db.GormClient.Model(&model.Role{}).Where("id = ?", id).Delete(nil).Error
}

// 角色列表
func (*Role) GetPage(page, size int) ([]*model.Role, int) {

	var (
		list []*model.Role
		count int64
	)

	db.GormClient.Model(&model.Role{}).Order("id desc").Count(&count).Limit(size).Offset((page - 1) * size).Find(&list)

	return list, int(count)
}

// 角色详情
func (*Role) GetDetail(id int) *model.Role {

	var detail *model.Role

	db.GormClient.Model(&model.Role{}).Where("id = ?", id).Take(&detail)

	return detail
}

// 角色详情
func (*Role) GetDetailByName(name string) *model.Role {

	var detail *model.Role

	db.GormClient.Model(&model.Role{}).Where("name = ?", name).Take(&detail)

	return detail
}