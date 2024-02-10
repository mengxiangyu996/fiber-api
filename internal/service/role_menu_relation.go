package service

import (
	"breeze-api/internal/model"
	"breeze-api/pkg/db"
)

// 角色菜单关系数据服务
type RoleMenuRelation struct{}

// 绑定菜单
func (*RoleMenuRelation) Bind(roleId int, menuIds []int) error {

	// 开启事务
	tx := db.GormClient.Begin()

	// 删除已绑定菜单
	if err := tx.Model(&model.RoleMenuRelation{}).Where("role_id = ?", roleId).Delete(nil).Error; err != nil {
		tx.Rollback()
		return err
	}

	if len(menuIds) > 0 {
		// 重新绑定角色
		for _, menuId := range menuIds {
			if err := tx.Model(&model.RoleMenuRelation{}).Create(&model.RoleMenuRelation{
				RoleId: roleId,
				MenuId: menuId,
			}).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	tx.Commit()

	return nil
}

// 绑定菜单列表
func (*RoleMenuRelation) List(roleId int) []*model.RoleMenuRelation {

	var list []*model.RoleMenuRelation

	db.GormClient.Model(&model.RoleMenuRelation{}).Where("role_id = ?", roleId).Find(&list)

	return list
}
