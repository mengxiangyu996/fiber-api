package admin

import (
	"breeze-api/internal/service"
	"breeze-api/pkg/response"

	"github.com/gofiber/fiber/v2"
)

// 配置请求
type Config struct{}

// 创建配置
func (*Config) Create(ctx *fiber.Ctx) error {

	type request struct {
		GroupName   string `json:"groupName"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Value       string `json:"value"`
		Remark      string `json:"remark"`
		Status      int    `json:"status"`
	}

	var req request

	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, err.Error())
	}

	if req.Name == "" {
		return response.Error(ctx, "参数错误")
	}

	if config := (&service.Config{}).DetailByName(req.Name); config.Id > 0 {
		return response.Error(ctx, "配置名称已存在")
	}

	if err := (&service.Config{}).Create(&service.Config{
		GroupName:   req.GroupName,
		Name:        req.Name,
		Description: req.Description,
		Value:       req.Value,
		Remark:      req.Remark,
		Status:      req.Status,
	}); err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 更新配置
func (*Config) Update(ctx *fiber.Ctx) error {

	type request struct {
		Id          int    `json:"id"`
		GroupName   string `json:"groupName"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Value       string `json:"value"`
		Remark      string `json:"remark"`
		Status      int    `json:"status"`
	}

	var req request

	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, err.Error())
	}

	if req.Id <= 0 || req.Name == "" {
		return response.Error(ctx, "参数错误")
	}

	if config := (&service.Config{}).DetailByName(req.Name); config.Id > 0 && config.Id != req.Id {
		return response.Error(ctx, "配置名称已存在")
	}

	if err := (&service.Config{}).Update(&service.Config{
		Id:          req.Id,
		GroupName:   req.GroupName,
		Name:        req.Name,
		Description: req.Description,
		Value:       req.Value,
		Remark:      req.Remark,
		Status:      req.Status,
	}); err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 删除配置
func (*Config) Delete(ctx *fiber.Ctx) error {

	type request struct {
		Id int `json:"id"`
	}

	var req request

	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, err.Error())
	}

	if req.Id <= 0 {
		return response.Error(ctx, "参数错误")
	}

	if err := (&service.Config{}).Delete(req.Id); err != nil {
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
