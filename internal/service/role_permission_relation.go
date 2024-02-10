package service

import (
	"breeze-api/internal/model"
	"breeze-api/pkg/db"
)

// 角色权限关系数据服务
type RolePermissionRelation struct{}

// 绑定权限
func (*RolePermissionRelation) Bind(roleId int, permissionIds []int) error {

	// 开启事务
	tx := db.GormClient.Begin()

	// 删除已绑定权限
	if err := tx.Model(&model.RolePermissionRelation{}).Where("role_id = ?", roleId).Delete(nil).Error; err != nil {
		tx.Rollback()
		return err
	}

	if len(permissionIds) > 0 {
		// 重新绑定权限
		for _, permissionId := range permissionIds {
			if err := tx.Model(&model.RolePermissionRelation{}).Create(&model.RolePermissionRelation{
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
func (*RolePermissionRelation) List(roleId int) []*model.RolePermissionRelation {

	var list []*model.RolePermissionRelation

	db.GormClient.Model(&model.RolePermissionRelation{}).Where("role_id = ?", roleId).Find(&list)

	return list
}

// 权限绑定详情
func (*RolePermissionRelation) DetailByRoleIdWithPermissionId(roleId, permissionId int) *model.RolePermissionRelation {

	var detail *model.RolePermissionRelation

	db.GormClient.Model(&model.RolePermissionRelation{}).Where("role_id = ?", roleId).Where("permission_id = ?", permissionId).Take(&detail)

	return detail
}
