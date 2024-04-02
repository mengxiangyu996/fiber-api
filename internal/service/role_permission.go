package service

import (
	"breeze-api/internal/model"
	"breeze-api/pkg/db"
)

// 角色权限关系数据服务
type RolePermission struct {
	Id           int `json:"id"`
	RoleId       int `json:"roleId"`
	PermissionId int `json:"permissionId"`
}

// 绑定权限
func (*RolePermission) Bind(roleId int, permissionIds []int) error {

	// 开启事务
	tx := db.GormClient.Begin()

	// 删除已绑定权限
	if err := tx.Model(&model.RolePermission{}).Where("role_id = ?", roleId).Delete(nil).Error; err != nil {
		tx.Rollback()
		return err
	}

	if len(permissionIds) > 0 {
		// 重新绑定权限
		for _, permissionId := range permissionIds {
			if err := tx.Model(&model.RolePermission{}).Create(&model.RolePermission{
				RoleId:       roleId,
				PermissionId: permissionId,
			}).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	tx.Commit()

	return nil
}

// 绑定权限列表
func (*RolePermission) List(roleId int) []*RolePermission {

	var list []*RolePermission

	db.GormClient.Model(&model.RolePermission{}).Where("role_id = ?", roleId).Find(&list)

	return list
}

// 权限绑定详情
func (*RolePermission) DetailByRoleIdWithPermissionId(roleId, permissionId int) *RolePermission {

	var detail *RolePermission

	db.GormClient.Model(&model.RolePermission{}).Where("role_id = ?", roleId).Where("permission_id = ?", permissionId).Take(&detail)

	return detail
}
