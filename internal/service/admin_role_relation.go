package service

import (
	"breeze-api/internal/model"
	"breeze-api/pkg/db"
)

// 管理员角色关系数据服务
type AdminRoleRelation struct{}

// 绑定角色
func (*AdminRoleRelation) Bind(adminId int, roleIds []int) error {

	// 开启事务
	tx := db.GormClient.Begin()

	// 删除已绑定角色
	if err := tx.Model(&model.AdminRoleRelation{}).Where("admin_id = ?", adminId).Delete(nil).Error; err != nil {
		tx.Rollback()
		return err
	}

	if len(roleIds) > 0 {
		// 重新绑定角色
		for _, roleId := range roleIds {
			if err := tx.Model(&model.AdminRoleRelation{}).Create(&model.AdminRoleRelation{
				AdminId: adminId,
				RoleId:  roleId,
			}).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	tx.Commit()

	return nil
}

// 绑定角色列表
func (*AdminRoleRelation) List(adminId int) []*model.AdminRoleRelation {

	var list []*model.AdminRoleRelation

	db.GormClient.Model(&model.AdminRoleRelation{}).Where("admin_id = ?", adminId).Find(&list)

	return list
}
