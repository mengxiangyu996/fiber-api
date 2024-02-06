package admin

import (
	"breeze-api/helper"
	"breeze-api/helper/encrypt"
	"breeze-api/helper/jwt"
	"breeze-api/internal/model"
	"breeze-api/internal/service"
	"breeze-api/pkg/response"
	"time"

	"github.com/gofiber/fiber/v2"
)

// 管理员请求
type Admin struct{}

// 管理员详情返回
type AdminResult struct {
	Id         int                 `json:"id"`
	CreateTime time.Time           `json:"createTime"`
	UpdateTime time.Time           `json:"updateTime"`
	Username   string              `json:"username"`
	Nickname   string              `json:"nickname"`
	Gender     int                 `json:"gender"`
	Email      string              `json:"email"`
	Phone      string              `json:"phone"`
	Avatar     string              `json:"avatar"`
	Status     int                 `json:"status"`
	Roles      []*AdminRole        `json:"roles"`
	Menus      []*service.MenuTree `json:"menus"`
}

// 管理员绑定的角色
type AdminRole struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// 创建管理员
func (*Admin) Create(ctx *fiber.Ctx) error {

	type request struct {
		Username string `json:"username"`
		Nickname string `json:"nickname"`
		Gender   int    `json:"gender"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Avatar   string `json:"avatar"`
		Status   int    `json:"status"`
	}

	var req request

	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, err.Error())
	}

	if req.Username == "" {
		return response.Error(ctx, "缺少参数")
	}

	admin := (&service.Admin{}).GetDetailByUsername(req.Username)
	if admin.Id > 0 {
		return response.Error(ctx, "管理员已存在")
	}

	if adminId := (&service.Admin{}).Create(&model.Admin{
		Username: req.Username,
		Nickname: req.Nickname,
		Password: encrypt.Generate("123456"),
		Gender:   req.Gender,
		Email:    req.Email,
		Phone:    req.Phone,
		Avatar:   req.Avatar,
		Status:   req.Status,
	}); adminId <= 0 {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 更新管理员
func (*Admin) Update(ctx *fiber.Ctx) error {

	type request struct {
		Id       int    `json:"id"`
		Nickname string `json:"nickname"`
		Gender   int    `json:"gender"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Avatar   string `json:"avatar"`
		Status   int    `json:"status"`
	}

	var req request

	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, err.Error())
	}

	if adminId := (&service.Admin{}).Update(&model.Admin{
		Id:       req.Id,
		Nickname: req.Nickname,
		Gender:   req.Gender,
		Email:    req.Email,
		Phone:    req.Phone,
		Avatar:   req.Avatar,
		Status:   req.Status,
	}); adminId <= 0 {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 删除管理员
func (*Admin) Delete(ctx *fiber.Ctx) error {

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

	if err := (&service.Admin{}).Delete(req.Id); err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 管理员列表
func (*Admin) Page(ctx *fiber.Ctx) error {

	page := ctx.QueryInt("page", 1)
	size := ctx.QueryInt("size", 10)

	type request struct {
		Username string `query:"username"`
		Nickname string `query:"nickname"`
		Email    string `query:"email"`
		Phone    string `query:"phone"`
	}

	var req request

	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, err.Error())
	}

	list, count := (&service.Admin{}).GetPage(page, size, req.Username, req.Nickname, req.Email, req.Phone)

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

	admin := (&service.Admin{}).GetDetail(id)
	if admin.Id > 0 {
		adminResult.Id = admin.Id
		adminResult.CreateTime = admin.CreateTime
		adminResult.UpdateTime = admin.UpdateTime
		adminResult.Username = admin.Username
		adminResult.Nickname = admin.Nickname
		adminResult.Gender = admin.Gender
		adminResult.Email = admin.Email
		adminResult.Phone = admin.Phone
		adminResult.Avatar = admin.Avatar
		adminResult.Status = admin.Status
		// 管理员权限
		adminRoles := (&service.AdminRoleRelation{}).GetList(admin.Id)
		if len(adminRoles) > 0 {
			for _, adminRole := range adminRoles {
				role := (&service.Role{}).GetDetail(adminRole.RoleId)
				if role.Status != 1 {
					continue
				}
				adminResult.Roles = append(adminResult.Roles, &AdminRole{
					Id:   role.Id,
					Name: role.Name,
				})
				// 角色绑定的菜单
				roleMenus := (&service.RoleMenuRelation{}).GetList(role.Id)
				if len(roleMenus) <= 0 {
					continue
				}
				var menuIds []int
				for _, roleMenu := range roleMenus {
					menuIds = append(menuIds, roleMenu.MenuId)
				}
				adminResult.Menus = (&service.Menu{}).ListToTree((&service.Menu{}).GetListByIds(menuIds), 0)
			}
		}
	}

	return response.Success(ctx, "成功", map[string]interface{}{
		"admin": adminResult,
	})
}

// 管理员登录
func (*Admin) Login(ctx *fiber.Ctx) error {

	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var req request

	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, err.Error())
	}

	if req.Username == "" || req.Password == "" {
		return response.Error(ctx, "参数错误")
	}

	admin := (&service.Admin{}).GetDetailByUsername(req.Username)
	if admin.Id <= 0 {
		return response.Error(ctx, "账号不存在")
	}

	if !encrypt.Compare(admin.Password, req.Password) {
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

// 修改密码
func (*Admin) ChangePassword(ctx *fiber.Ctx) error {

	type request struct {
		Password string `json:"password"`
	}

	var req request

	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, err.Error())
	}

	if req.Password == "" {
		return response.Error(ctx, "参数错误")
	}

	id, _ := helper.GetTokenPayload(ctx)

	if adminId := (&service.Admin{}).Update(&model.Admin{
		Id:       id,
		Password: encrypt.Generate(req.Password),
	}); adminId <= 0 {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 绑定角色
func (*Admin) BindRole(ctx *fiber.Ctx) error {

	type request struct {
		AdminId int   `json:"adminId"`
		RoleIds []int `json:"roleIds"`
	}

	var req request

	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, err.Error())
	}

	if req.AdminId <= 0 {
		return response.Error(ctx, "参数错误")
	}

	if err := (&service.AdminRoleRelation{}).Bind(req.AdminId, req.RoleIds); err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}
