package admin

import (
	"breeze-api/internal/model"
	"breeze-api/internal/service"
	"breeze-api/pkg/response"

	"github.com/gofiber/fiber/v2"
)

// 角色请求
type Role struct{}

// 创建角色
func (*Role) Create(ctx *fiber.Ctx) error {

	type request struct {
		Name   string `json:"name"`
		Status int    `json:"status"`
	}

	var req request

	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, err.Error())
	}

	if req.Name == "" {
		return response.Error(ctx, "参数错误")
	}

	role := (&service.Role{}).GetDetailByName(req.Name)
	if role.Id > 0 {
		return response.Error(ctx, "角色已存在")
	}

	if err := (&service.Role{}).Create(&model.Role{
		Name:   req.Name,
		Status: req.Status,
	}); err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 更新角色
func (*Role) Update(ctx *fiber.Ctx) error {

	type request struct {
		Id     int    `json:"id"`
		Name   string `json:"name"`
		Status int    `json:"status"`
	}

	var req request

	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, err.Error())
	}

	if req.Id <= 0 || req.Name == "" {
		return response.Error(ctx, "参数错误")
	}

	role := (&service.Role{}).GetDetailByName(req.Name)
	if role.Id > 0 && req.Id != role.Id {
		return response.Error(ctx, "角色已存在")
	}

	if err := (&service.Role{}).Update(&model.Role{
		Id:     req.Id,
		Name:   req.Name,
		Status: req.Status,
	}); err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 删除角色
func (*Role) Delete(ctx *fiber.Ctx) error {

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

	if err := (&service.Role{}).Delete(req.Id); err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 角色列表
func (*Role) Page(ctx *fiber.Ctx) error {

	page := ctx.QueryInt("page", 1)
	size := ctx.QueryInt("size", 10)

	list, count := (&service.Role{}).GetPage(page, size)

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

	role := (&service.Role{}).GetDetail(id)

	return response.Success(ctx, "成功", map[string]interface{}{
		"role": role,
	})
}

// 绑定菜单
func (*Role) BindMenu(ctx *fiber.Ctx) error {

	type request struct {
		RoleId  int   `json:"roleId"`
		MenuIds []int `json:"menuIds"`
	}

	var req request

	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, err.Error())
	}

	if req.RoleId <= 0 {
		return response.Error(ctx, "参数错误")
	}

	if err := (&service.RoleMenuRelation{}).Bind(req.RoleId, req.MenuIds); err != nil {
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

	roleMenus := (&service.RoleMenuRelation{}).GetList(roleId)
	if len(roleMenus) > 0 {
		for _, roleMenu := range roleMenus {
			menuIds = append(menuIds, roleMenu.MenuId)
		}
	}

	tree := (&service.Menu{}).ListToTree((&service.Menu{}).GetListByIds(nil), 0)
	bindTree := (&service.Menu{}).ListToTree((&service.Menu{}).GetListByIds(menuIds), 0)

	return response.Success(ctx, "成功", map[string]interface{}{
		"tree":     tree,
		"bindTree": bindTree,
	})
}

// 绑定权限
func (*Role) BindPermission(ctx *fiber.Ctx) error {

	type request struct {
		RoleId        int   `json:"roleId"`
		PermissionIds []int `json:"permissionIds"`
	}

	var req request

	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, err.Error())
	}

	if req.RoleId <= 0 {
		return response.Error(ctx, "参数错误")
	}

	if err := (&service.RolePermissionRelation{}).Bind(req.RoleId, req.PermissionIds); err != nil {
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

	rolePermissions := (&service.RolePermissionRelation{}).GetList(roleId)
	if len(rolePermissions) > 0 {
		for _, rolePermission := range rolePermissions {
			permissionIds = append(permissionIds, rolePermission.PermissionId)
		}
	}

	tree := (&service.Permission{}).GetListByIds(nil)
	bindTree := (&service.Permission{}).GetListByIds(permissionIds)

	return response.Success(ctx, "成功", map[string]interface{}{
		"tree":     tree,
		"bindTree": bindTree,
	})
}
