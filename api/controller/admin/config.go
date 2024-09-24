package admin

import (
	"breeze-api/api/service"
	"breeze-api/pkg/response"

	"github.com/gofiber/fiber/v2"
)

// 配置请求
type Config struct{}

// 创建配置
func (*Config) Create(ctx *fiber.Ctx) error {

	var param struct {
		GroupName   string `json:"groupName"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Value       string `json:"value"`
		Remark      string `json:"remark"`
		Status      int    `json:"status"`
	}

	if err := ctx.BodyParser(&param); err != nil {
		return response.Error(ctx, err.Error())
	}

	if param.Name == "" {
		return response.Error(ctx, "参数错误")
	}

	if config := (&service.Config{}).DetailByName(param.Name); config.Id > 0 {
		return response.Error(ctx, "配置名称已存在")
	}

	if err := (&service.Config{}).Create(&service.Config{
		GroupName:   param.GroupName,
		Name:        param.Name,
		Description: param.Description,
		Value:       param.Value,
		Remark:      param.Remark,
		Status:      param.Status,
	}); err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 更新配置
func (*Config) Update(ctx *fiber.Ctx) error {

	var param struct {
		Id          int    `json:"id"`
		GroupName   string `json:"groupName"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Value       string `json:"value"`
		Remark      string `json:"remark"`
		Status      int    `json:"status"`
	}

	if err := ctx.BodyParser(&param); err != nil {
		return response.Error(ctx, err.Error())
	}

	if param.Id <= 0 || param.Name == "" {
		return response.Error(ctx, "参数错误")
	}

	if config := (&service.Config{}).DetailByName(param.Name); config.Id > 0 && config.Id != param.Id {
		return response.Error(ctx, "配置名称已存在")
	}

	if err := (&service.Config{}).Update(&service.Config{
		Id:          param.Id,
		GroupName:   param.GroupName,
		Name:        param.Name,
		Description: param.Description,
		Value:       param.Value,
		Remark:      param.Remark,
		Status:      param.Status,
	}); err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 删除配置
func (*Config) Delete(ctx *fiber.Ctx) error {

	id := ctx.QueryInt("id")

	if id <= 0 {
		return response.Error(ctx, "参数错误")
	}

	if err := (&service.Config{}).Delete(id); err != nil {
		return response.Error(ctx, err.Error())
	}

	return response.Success(ctx, "成功", nil)
}

// 配置列表
func (*Config) Tab(ctx *fiber.Ctx) error {

	list := (&service.Config{}).Tab()

	return response.Success(ctx, "成功", map[string]interface{}{
		"list": list,
	})
}

// 配置详情
func (*Config) Detail(ctx *fiber.Ctx) error {

	id := ctx.QueryInt("id")

	config := (&service.Config{}).Detail(id)

	return response.Success(ctx, "成功", map[string]interface{}{
		"config": config,
	})
}
