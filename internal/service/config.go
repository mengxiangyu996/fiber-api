package service

import (
	"breeze-api/internal/model"
	"breeze-api/pkg/db"
)

// 配置数据服务
type Config struct{}

// 创建配置
func (*Config) Create(config *model.Config) error {
	return db.GormClient.Model(&model.Config{}).Create(&config).Error
}

// 更新配置
func (*Config) Update(config *model.Config) error {
	return db.GormClient.Model(&model.Config{}).Where("id = ?", config.Id).Updates(&config).Error
}

// 删除配置
func (*Config) Delete(id int) error {
	return db.GormClient.Model(&model.Config{}).Where("id = ?", id).Delete(nil).Error
}

// 配置列表
func (*Config) Tab() map[string][]*model.Config {

	var (
		groupName []string
		list      = make(map[string][]*model.Config, 0)
	)

	db.GormClient.Model(&model.Config{}).Group("group_name").Pluck("group_name", &groupName)

	for _, item := range groupName {
		var configs []*model.Config
		db.GormClient.Model(&model.Config{}).Where("group_name = ?", item).Find(&configs)
		list[item] = configs
	}

	return list
}

// 配置详情
func (*Config) Detail(id int) *model.Config {

	var detail *model.Config

	db.GormClient.Model(&model.Config{}).Where("id = ?", id).Take(&detail)

	return detail
}

// 配置详情
func (*Config) DetailByName(name string) *model.Config {

	var detail *model.Config

	db.GormClient.Model(&model.Config{}).Where("name = ?", name).Take(&detail)

	return detail
}
