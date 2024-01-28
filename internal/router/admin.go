package router

import (
	"breeze-api/internal/handler/admin"
	"breeze-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

// 后台路由
func AdminRouter(app *fiber.App) {

	// 未授权路由组
	api := app.Group("api/admin")
	{
		api.Post("admin/login", (&admin.Admin{}).Login) // 管理员登录
	}

	// 授权路由组
	authApi := app.Group("api/admin", (&middleware.Admin{}).Handle)
	{
		// 管理员操作
		authApi.Post("admin/create", (&admin.Admin{}).Create)                 // 创建管理员
		authApi.Post("admin/update", (&admin.Admin{}).Update)                 // 更新管理员
		authApi.Post("admin/delete", (&admin.Admin{}).Delete)                 // 删除管理员
		authApi.Get("admin/page", (&admin.Admin{}).Page)                      // 管理员列表
		authApi.Get("admin/detail", (&admin.Admin{}).Detail)                  // 管理员详情
		authApi.Post("admin/changePassword", (&admin.Admin{}).ChangePassword) // 管理员修改密码
		authApi.Post("admin/bindRole", (&admin.Admin{}).BindRole)             // 管理员绑定角色

		// 角色操作
		authApi.Post("role/create", (&admin.Role{}).Create)                 // 创建角色
		authApi.Post("role/update", (&admin.Role{}).Update)                 // 更新角色
		authApi.Post("role/delete", (&admin.Role{}).Delete)                 // 删除角色
		authApi.Get("role/page", (&admin.Role{}).Page)                      // 角色列表
		authApi.Get("role/detail", (&admin.Role{}).Detail)                  // 角色详情
		authApi.Post("role/bindMenu", (&admin.Role{}).BindMenu)             // 角色绑定菜单
		authApi.Get("role/menus", (&admin.Role{}).Menus)                    // 角色菜单列表
		authApi.Post("role/bindPermission", (&admin.Role{}).BindPermission) // 角色绑定权限
		authApi.Get("role/permissions", (&admin.Role{}).Permissions)        // 角色权限列表

		// 菜单操作
		authApi.Post("menu/create", (&admin.Menu{}).Create) // 创建菜单
		authApi.Post("menu/update", (&admin.Menu{}).Update) // 更新菜单
		authApi.Post("menu/delete", (&admin.Menu{}).Delete) // 删除菜单
		authApi.Get("menu/tree", (&admin.Menu{}).Tree)      // 菜单列表
		authApi.Get("menu/detail", (&admin.Menu{}).Detail)  // 菜单详情

		// 权限操作
		authApi.Post("permission/create", (&admin.Permission{}).Create) // 创建权限
		authApi.Post("permission/update", (&admin.Permission{}).Update) // 更新权限
		authApi.Post("permission/delete", (&admin.Permission{}).Delete) // 删除权限
		authApi.Get("permission/page", (&admin.Permission{}).Page)      // 权限列表
		authApi.Get("permission/detail", (&admin.Permission{}).Detail)  // 权限详情
	}

}
