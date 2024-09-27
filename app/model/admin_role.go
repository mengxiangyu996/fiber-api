package model

// 管理员角色关系模型
type AdminRole struct {
	BaseModel
	AdminId int
	RoleId  int
}
