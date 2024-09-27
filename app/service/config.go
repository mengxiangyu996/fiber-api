package service

import (
	"fiber-api/app/model"
	"fiber-api/pkg/db"
)

// 配置数据服务
type Config struct {
	Id          int    `json:"id"`
	GroupName   string `json:"groupName"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Value       string `json:"value"`
	Remark      string `json:"remark"`
	Status      int    `json:"status"`
}

// 创建配置
func (*Config) Create(config *Config) error {
	return db.GormClient.Model(&model.Config{}).Create(&model.Config{
		GroupName:   config.GroupName,
		Name:        config.Name,
		Description: config.Description,
		Value:       config.Value,
		Remark:      config.Remark,
		Status:      config.Status,
	}).Error
}

// 更新配置
func (*Config) Update(config *Config) error {
	return db.GormClient.Model(&model.Config{}).Where("id = ?", config.Id).Updates(&model.Config{
		GroupName:   config.GroupName,
		Name:        config.Name,
		Description: config.Description,
		Value:       config.Value,
		Remark:      config.Remark,
		Status:      config.Status,
	}).Error
}

// 删除配置
func (*Config) Delete(id int) error {
	return db.GormClient.Model(&model.Config{}).Where("id = ?", id).Delete(nil).Error
}

// 配置列表
func (*Config) Tab() map[string][]*Config {

	var (
		groupName []string
		list      = make(map[string][]*Config, 0)
	)

	db.GormClient.Model(&model.Config{}).Group("group_name").Pluck("group_name", &groupName)

	for _, item := range groupName {
		var configs []*Config
		db.GormClient.Model(&model.Config{}).Where("group_name = ?", item).Find(&configs)
		list[item] = configs
	}

	return list
}

// 配置详情
func (*Config) Detail(id int) *Config {

	var detail *Config

	db.GormClient.Model(&model.Config{}).Where("id = ?", id).Take(&detail)

	return detail
}

// 配置详情
func (*Config) DetailByName(name string) *Config {

	var detail *Config

	db.GormClient.Model(&model.Config{}).Where("name = ?", name).Take(&detail)

	return detail
}
