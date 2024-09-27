package admin

import (
	"fiber-api/app/service"
	"fiber-api/pkg/response"

	"github.com/gofiber/fiber/v2"
)

// 角色请求
type Role struct{}

// 创建角色
func (*Role) Create(ctx *fiber.Ctx) error {

	var param struct {
		Name   string `json:"name"`
		Status int    `json:"status"`
	}

	if err := ctx.BodyParser(&param); err != nil {
		return response.Error(ctx, err.Error())
	}

	if param.Name == "" {
		return response.Error(ctx, "参数错误")
	}

	role := (&service.Role{}).DetailByName(param.Name)
	if role.Id > 0 {
		return response.Error(ctx, "角色已存在")
	}

	if err := (&service.Role{}).Create(&service.Role{
		Name:   param.Name,
		Status: param.Status,
	}); err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 更新角色
func (*Role) Update(ctx *fiber.Ctx) error {

	var param struct {
		Id     int    `json:"id"`
		Name   string `json:"name"`
		Status int    `json:"status"`
	}

	if err := ctx.BodyParser(&param); err != nil {
		return response.Error(ctx, err.Error())
	}

	if param.Id <= 0 || param.Name == "" {
		return response.Error(ctx, "参数错误")
	}

	role := (&service.Role{}).DetailByName(param.Name)
	if role.Id > 0 && param.Id != role.Id {
		return response.Error(ctx, "角色已存在")
	}

	if err := (&service.Role{}).Update(&service.Role{
		Id:     param.Id,
		Name:   param.Name,
		Status: param.Status,
	}); err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 删除角色
func (*Role) Delete(ctx *fiber.Ctx) error {

	id := ctx.QueryInt("id")

	if id <= 0 {
		return response.Error(ctx, "参数错误")
	}

	if err := (&service.Role{}).Delete(id); err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 角色列表
func (*Role) Page(ctx *fiber.Ctx) error {

	page := ctx.QueryInt("page", 1)
	size := ctx.QueryInt("size", 10)

	list, count := (&service.Role{}).Page(page, size)

	return response.Success(ctx, "成功", map[string]interface{}{
		"list":  list,
		"count": count,
	})
}

// 角色详情
func (*Role) Detail(ctx *fiber.Ctx) error {

	id := ctx.QueryInt("id")
	if id <= 0 {
		return response.Error(ctx, "参数错误")
	}

	role := (&service.Role{}).Detail(id)

	return response.Success(ctx, "成功", map[string]interface{}{
		"role": role,
	})
}

// 绑定菜单
func (*Role) BindMenu(ctx *fiber.Ctx) error {

	var param struct {
		RoleId  int   `json:"roleId"`
		MenuIds []int `json:"menuIds"`
	}

	if err := ctx.BodyParser(&param); err != nil {
		return response.Error(ctx, err.Error())
	}

	if param.RoleId <= 0 {
		return response.Error(ctx, "参数错误")
	}

	if err := (&service.RoleMenu{}).Bind(param.RoleId, param.MenuIds); err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 绑定菜单列表
func (*Role) Menus(ctx *fiber.Ctx) error {

	roleId := ctx.QueryInt("roleId")
	if roleId <= 0 {
		return response.Error(ctx, "参数错误")
	}

	var menuIds []int

	roleMenus := (&service.RoleMenu{}).List(roleId)
	if len(roleMenus) > 0 {
		for _, roleMenu := range roleMenus {
			menuIds = append(menuIds, roleMenu.MenuId)
		}
	}

	tree := (&service.Menu{}).ListToTree((&service.Menu{}).ListByIds(nil), 0)
	bindTree := (&service.Menu{}).ListToTree((&service.Menu{}).ListByIds(menuIds), 0)

	return response.Success(ctx, "成功", map[string]interface{}{
		"tree":     tree,
		"bindTree": bindTree,
	})
}

// 绑定权限
func (*Role) BindPermission(ctx *fiber.Ctx) error {

	var param struct {
		RoleId        int   `json:"roleId"`
		PermissionIds []int `json:"permissionIds"`
	}

	if err := ctx.BodyParser(&param); err != nil {
		return response.Error(ctx, err.Error())
	}

	if param.RoleId <= 0 {
		return response.Error(ctx, "参数错误")
	}

	if err := (&service.RolePermission{}).Bind(param.RoleId, param.PermissionIds); err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 绑定权限列表
func (*Role) Permissions(ctx *fiber.Ctx) error {

	roleId := ctx.QueryInt("roleId")
	if roleId <= 0 {
		return response.Error(ctx, "参数错误")
	}

	var permissionIds []int

	rolePermissions := (&service.RolePermission{}).List(roleId)
	if len(rolePermissions) > 0 {
		for _, rolePermission := range rolePermissions {
			permissionIds = append(permissionIds, rolePermission.PermissionId)
		}
	}

	tree := (&service.Permission{}).ListByIds(nil)
	bindTree := (&service.Permission{}).ListByIds(permissionIds)

	return response.Success(ctx, "成功", map[string]interface{}{
		"tree":     tree,
		"bindTree": bindTree,
	})
}
