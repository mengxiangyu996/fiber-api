package admin

import (
	"breeze-api/helper"
	"breeze-api/helper/encrypt"
	"breeze-api/helper/jwt"
	"breeze-api/internal/service"
	"breeze-api/pkg/response"
	"time"

	"github.com/gofiber/fiber/v2"
)

// 管理员请求
type Admin struct{}

// 管理员详情返回
type AdminResult struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Gender   int    `json:"gender"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
	Status   int    `json:"status"`
}

// 管理员绑定的角色
type AdminRole struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type RoleMenu struct {
	Roles []*AdminRole        `json:"roles"`
	Menus []*service.MenuTree `json:"menus"`
}

// 创建管理员
func (*Admin) Create(ctx *fiber.Ctx) error {

	var param struct {
		Username string `json:"username"`
		Nickname string `json:"nickname"`
		Gender   int    `json:"gender"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Avatar   string `json:"avatar"`
		Status   int    `json:"status"`
	}

	if err := ctx.BodyParser(&param); err != nil {
		return response.Error(ctx, err.Error())
	}

	if param.Username == "" {
		return response.Error(ctx, "参数错误")
	}

	admin := (&service.Admin{}).DetailByUsername(param.Username)
	if admin.Id > 0 {
		return response.Error(ctx, "管理员已存在")
	}

	if adminId := (&service.Admin{}).Create(&service.Admin{
		Username: param.Username,
		Nickname: param.Nickname,
		Password: encrypt.Generate("123456"),
		Gender:   param.Gender,
		Email:    param.Email,
		Phone:    param.Phone,
		Avatar:   param.Avatar,
		Status:   param.Status,
	}); adminId <= 0 {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 更新管理员
func (*Admin) Update(ctx *fiber.Ctx) error {

	var param struct {
		Id       int    `json:"id"`
		Nickname string `json:"nickname"`
		Gender   int    `json:"gender"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Avatar   string `json:"avatar"`
		Status   int    `json:"status"`
	}

	if err := ctx.BodyParser(&param); err != nil {
		return response.Error(ctx, err.Error())
	}

	if adminId := (&service.Admin{}).Update(&service.Admin{
		Id:       param.Id,
		Nickname: param.Nickname,
		Gender:   param.Gender,
		Email:    param.Email,
		Phone:    param.Phone,
		Avatar:   param.Avatar,
		Status:   param.Status,
	}); adminId <= 0 {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 删除管理员
func (*Admin) Delete(ctx *fiber.Ctx) error {

	id := ctx.QueryInt("id", 1)

	if id <= 0 {
		return response.Error(ctx, "参数错误")
	}

	if err := (&service.Admin{}).Delete(id); err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 管理员列表
func (*Admin) Page(ctx *fiber.Ctx) error {

	page := ctx.QueryInt("page", 1)
	size := ctx.QueryInt("size", 10)

	var param struct {
		Username string `query:"username"`
		Nickname string `query:"nickname"`
		Email    string `query:"email"`
		Phone    string `query:"phone"`
	}

	if err := ctx.QueryParser(&param); err != nil {
		return response.Error(ctx, err.Error())
	}

	list, count := (&service.Admin{}).Page(page, size, param.Username, param.Nickname, param.Email, param.Phone)

	// 清除密码
	for _, item := range list {
		item.Password = ""
	}

	return response.Success(ctx, "成功", map[string]interface{}{
		"list":  list,
		"count": count,
	})
}

// 管理员详情
func (*Admin) Detail(ctx *fiber.Ctx) error {

	id := ctx.QueryInt("id")
	// 未传值默认当前用户
	if id <= 0 {
		id, _ = helper.GetTokenPayload(ctx)
	}

	var adminResult AdminResult

	admin := (&service.Admin{}).Detail(id)
	if admin.Id > 0 {
		adminResult.Id = admin.Id
		adminResult.Username = admin.Username
		adminResult.Nickname = admin.Nickname
		adminResult.Gender = admin.Gender
		adminResult.Email = admin.Email
		adminResult.Phone = admin.Phone
		adminResult.Avatar = admin.Avatar
		adminResult.Status = admin.Status
	}

	return response.Success(ctx, "成功", map[string]interface{}{
		"admin": adminResult,
	})
}

// 管理员登录
func (*Admin) Login(ctx *fiber.Ctx) error {

	var param struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.BodyParser(&param); err != nil {
		return response.Error(ctx, err.Error())
	}

	if param.Username == "" || param.Password == "" {
		return response.Error(ctx, "参数错误")
	}

	admin := (&service.Admin{}).DetailByUsername(param.Username)
	if admin.Id <= 0 {
		return response.Error(ctx, "账号不存在")
	}

	if !encrypt.Compare(admin.Password, param.Password) {
		return response.Error(ctx, "密码错误")
	}

	token := jwt.Generate(&jwt.Payload{
		Id:     admin.Id,
		Expire: time.Now().AddDate(0, 0, 7),
	})

	return response.Success(ctx, "成功", map[string]interface{}{
		"token": token,
	})
}

// 管理员角色和菜单
func (*Admin) AdminRoleMenu(ctx *fiber.Ctx) error {

	id, err := helper.GetTokenPayload(ctx)
	if err != nil {
		return response.Error(ctx, err.Error())
	}

	if id == 1 {
		return response.Success(ctx, "ok", &RoleMenu{
			Menus: (&service.Menu{}).ListToTree((&service.Menu{}).ListByIds(nil), 0),
		})
	}

	adminRoles := (&service.AdminRole{}).List(id)
	if len(adminRoles) <= 0 {
		return response.Error(ctx, "角色不存在")
	}

	var roles []*AdminRole
	var menus []*service.MenuTree

	for _, adminRole := range adminRoles {
		role := (&service.Role{}).Detail(adminRole.RoleId)
		if role.Status != 1 {
			continue
		}
		roles = append(roles, &AdminRole{
			Id:   role.Id,
			Name: role.Name,
		})
		// 角色绑定的菜单
		roleMenus := (&service.RoleMenu{}).List(role.Id)
		if len(roleMenus) <= 0 {
			continue
		}
		var menuIds []int
		for _, roleMenu := range roleMenus {
			menuIds = append(menuIds, roleMenu.MenuId)
		}
		menus = (&service.Menu{}).ListToTree((&service.Menu{}).ListByIds(menuIds), 0)
	}

	return response.Success(ctx, "ok", &RoleMenu{
		Roles: roles,
		Menus: menus,
	})
}

// 修改密码
func (*Admin) ChangePassword(ctx *fiber.Ctx) error {

	var param struct {
		Password string `json:"password"`
	}

	if err := ctx.BodyParser(&param); err != nil {
		return response.Error(ctx, err.Error())
	}

	if param.Password == "" {
		return response.Error(ctx, "参数错误")
	}

	id, _ := helper.GetTokenPayload(ctx)

	if adminId := (&service.Admin{}).Update(&service.Admin{
		Id:       id,
		Password: encrypt.Generate(param.Password),
	}); adminId <= 0 {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 绑定角色
func (*Admin) BindRole(ctx *fiber.Ctx) error {

	var param struct {
		AdminId int   `json:"adminId"`
		RoleIds []int `json:"roleIds"`
	}

	if err := ctx.BodyParser(&param); err != nil {
		return response.Error(ctx, err.Error())
	}

	if param.AdminId <= 0 {
		return response.Error(ctx, "参数错误")
	}

	if err := (&service.AdminRole{}).Bind(param.AdminId, param.RoleIds); err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}
