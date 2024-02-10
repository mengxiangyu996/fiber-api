package service

import (
	"breeze-api/internal/model"
	"breeze-api/pkg/db"
)

// 权限数据服务
type Permission struct{}

// 创建权限
func (*Permission) Create(permission *model.Permission) error {
	return db.GormClient.Model(&model.Permission{}).Create(&permission).Error
}

// 更新权限
func (*Permission) Update(permission *model.Permission) error {
	return db.GormClient.Model(&model.Permission{}).Where("id = ?", permission.Id).Updates(&permission).Error
}

// 删除权限
func (*Permission) Delete(id int) error {
	return db.GormClient.Model(&model.Permission{}).Where("id = ?", id).Delete(nil).Error
}

// 权限列表
func (*Permission) Page(page, size int, name, groupName, path, method string) ([]*model.Permission, int) {

	var (
		list  []*model.Permission
		count int64
	)

	query := db.GormClient.Model(&model.Permission{}).Order("id desc, group_name")

	if name != "" {
		query.Where("name like ?", "%"+name+"%")
	}

	if groupName != "" {
		query.Where("group_name like ?", "%"+groupName+"%")
	}

	if path != "" {
		query.Where("path like ?", "%"+path+"%")
	}

	if method != "" {
		query.Where("method like ?", "%"+method+"%")
	}

	query.Count(&count).Limit(size).Offset((page - 1) * size).Find(&list)

	return list, int(count)
}

// 获取权限列表
func (*Permission) ListByIds(ids []int) []*model.Permission {

	var list []*model.Permission

	query := db.GormClient.Model(&model.Permission{}).Where("status = ?", 1)

	if len(ids) > 0 {
		query.Where("id in ?", ids)
	}

	query.Find(&list)

	return list
}

// 权限详情
func (*Permission) Detail(id int) *model.Permission {

	var detail *model.Permission

	db.GormClient.Model(&model.Permission{}).Where("id = ?", id).Take(&detail)

	return detail
}

// 权限详情
func (*Permission) DetailByPathWithMethod(path, method string) *model.Permission {

	var detail *model.Permission

	db.GormClient.Model(&model.Permission{}).Where("path = ?", path).Where("method = ?", method).Take(&detail)

	return detail
}
