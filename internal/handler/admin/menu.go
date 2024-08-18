package admin

import (
	"breeze-api/internal/service"
	"breeze-api/pkg/response"

	"github.com/gofiber/fiber/v2"
)

// 菜单请求
type Menu struct{}

// 创建菜单
func (*Menu) Create(ctx *fiber.Ctx) error {

	var param struct {
		ParentId  int    `json:"parentId"`
		Name      string `json:"name"`
		Type      int    `json:"type"`
		Sort      int    `json:"sort"`
		Path      string `json:"path"`
		Component string `json:"component"`
		Icon      string `json:"icon"`
		Redirect  string `json:"redirect"`
		Status    int    `json:"status"`
	}

	if err := ctx.BodyParser(&param); err != nil {
		return response.Error(ctx, err.Error())
	}

	if param.Name == "" || param.Type <= 0 || param.Path == "" {
		return response.Error(ctx, "参数错误")
	}

	if err := (&service.Menu{}).Create(&service.Menu{
		ParentId:  param.ParentId,
		Name:      param.Name,
		Type:      param.Type,
		Sort:      param.Sort,
		Path:      param.Path,
		Component: param.Component,
		Icon:      param.Icon,
		Redirect:  param.Redirect,
		Status:    param.Status,
	}); err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 更新菜单
func (*Menu) Update(ctx *fiber.Ctx) error {

	var param struct {
		Id        int    `json:"id"`
		ParentId  int    `json:"parentId"`
		Name      string `json:"name"`
		Type      int    `json:"type"`
		Sort      int    `json:"sort"`
		Path      string `json:"path"`
		Component string `json:"component"`
		Icon      string `json:"icon"`
		Redirect  string `json:"redirect"`
		Status    int    `json:"status"`
	}

	if err := ctx.BodyParser(&param); err != nil {
		return response.Error(ctx, err.Error())
	}

	if param.Id <= 0 {
		return response.Error(ctx, "参数错误")
	}

	if err := (&service.Menu{}).Update(&service.Menu{
		Id:        param.Id,
		ParentId:  param.ParentId,
		Name:      param.Name,
		Type:      param.Type,
		Sort:      param.Sort,
		Path:      param.Path,
		Component: param.Component,
		Icon:      param.Icon,
		Redirect:  param.Redirect,
		Status:    param.Status,
	}); err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 删除菜单
func (*Menu) Delete(ctx *fiber.Ctx) error {

	id := ctx.QueryInt("id")

	if id <= 0 {
		return response.Error(ctx, "参数错误")
	}

	list := (&service.Menu{}).ChildrenList(id, 0)
	if len(list) > 0 {
		return response.Error(ctx, "存在下级菜单")
	}

	if err := (&service.Menu{}).Delete(id); err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 菜单列表
func (*Menu) Tree(ctx *fiber.Ctx) error {

	tree := (&service.Menu{}).ChildrenList(0, 0)

	return response.Success(ctx, "成功", map[string]interface{}{
		"tree": tree,
	})
}

// 菜单详情
func (*Menu) Detail(ctx *fiber.Ctx) error {

	id := ctx.QueryInt("id")
	if id <= 0 {
		return response.Error(ctx, "参数错误")
	}

	menu := (&service.Menu{}).Detail(id)

	return response.Success(ctx, "成功", map[string]interface{}{
		"menu": menu,
	})
}
