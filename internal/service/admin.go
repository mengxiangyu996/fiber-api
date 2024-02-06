package service

import (
	"breeze-api/internal/model"
	"breeze-api/pkg/db"
)

// 管理员数据服务
type Admin struct{}

// 创建管理员
func (*Admin) Create(admin *model.Admin) int {

	err := db.GormClient.Model(&model.Admin{}).Create(&admin).Error
	if err != nil {
		return 0
	}

	return admin.Id
}

// 更新管理员
func (*Admin) Update(admin *model.Admin) int {

	err := db.GormClient.Model(&model.Admin{}).Where("id = ?", admin.Id).Updates(&admin).Error
	if err != nil {
		return 0
	}

	return admin.Id
}

// 删除管理员
func (*Admin) Delete(id int) error {
	return db.GormClient.Model(&model.Admin{}).Where("id = ?", id).Delete(nil).Error
}

// 管理员列表
func (*Admin) GetPage(page, size int, username, nickname, email, phone string) ([]*model.Admin, int) {

	var (
		list  []*model.Admin
		count int64
	)

	query := db.GormClient.Model(&model.Admin{}).Omit("password").Order("id desc")

	if username != "" {
		query.Where("username like ?", "%" + username + "%")
	}

	if nickname != "" {
		query.Where("nickname like ?", "%" + nickname + "%")
	}

	if email != "" {
		query.Where("email like ?", "%" + email + "%")
	}

	if phone != "" {
		query.Where("phone like ?", "%" + phone + "%")
	}

	query.Count(&count).Limit(size).Offset((page - 1) * size).Find(&list)

	return list, int(count)
}

// 管理员详情
func (*Admin) GetDetail(id int) *model.Admin {

	var detail *model.Admin

	db.GormClient.Model(&model.Admin{}).Where("id = ?", id).Take(&detail)

	return detail
}

// 管理员详情
func (*Admin) GetDetailByUsername(username string) *model.Admin {

	var detail *model.Admin

	db.GormClient.Model(&model.Admin{}).Where("username = ?", username).Take(&detail)

	return detail
}
