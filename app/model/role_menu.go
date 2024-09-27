package model

// 角色菜单关系模型
type RoleMenu struct {
	BaseModel
	RoleId int
	MenuId int
}
