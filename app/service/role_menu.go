package service

import (
	"fiber-api/app/model"
	"fiber-api/pkg/db"
)

// 角色菜单关系数据服务
type RoleMenu struct {
	Id     int `json:"id"`
	RoleId int `json:"roleId"`
	MenuId int `json:"menuId"`
}

// 绑定菜单
func (*RoleMenu) Bind(roleId int, menuIds []int) error {

	// 开启事务
	tx := db.GormClient.Begin()

	// 删除已绑定菜单
	if err := tx.Model(&model.RoleMenu{}).Where("role_id = ?", roleId).Delete(nil).Error; err != nil {
		tx.Rollback()
		return err
	}

	if len(menuIds) > 0 {
		// 重新绑定角色
		for _, menuId := range menuIds {
			if err := tx.Model(&model.RoleMenu{}).Create(&model.RoleMenu{
				RoleId: roleId,
				MenuId: menuId,
			}).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}

// 绑定菜单列表
func (*RoleMenu) List(roleId int) []*RoleMenu {

	var list []*RoleMenu

	db.GormClient.Model(&model.RoleMenu{}).Where("role_id = ?", roleId).Find(&list)

	return list
}
