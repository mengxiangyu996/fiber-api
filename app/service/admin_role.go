package service

import (
	"fiber-api/app/model"
	"fiber-api/pkg/db"
)

// 管理员角色关系数据服务
type AdminRole struct {
	Id      int `json:"id"`
	AdminId int `json:"adminId"`
	RoleId  int `json:"roleId"`
}

// 绑定角色
func (*AdminRole) Bind(adminId int, roleIds []int) error {

	// 开启事务
	tx := db.GormClient.Begin()

	// 删除已绑定角色
	if err := tx.Model(&model.AdminRole{}).Where("admin_id = ?", adminId).Delete(nil).Error; err != nil {
		tx.Rollback()
		return err
	}

	if len(roleIds) > 0 {
		// 重新绑定角色
		for _, roleId := range roleIds {
			if err := tx.Model(&model.AdminRole{}).Create(&model.AdminRole{
				AdminId: adminId,
				RoleId:  roleId,
			}).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}

// 绑定角色列表
func (*AdminRole) List(adminId int) []*AdminRole {

	var list []*AdminRole

	db.GormClient.Model(&model.AdminRole{}).Where("admin_id = ?", adminId).Find(&list)

	return list
}
