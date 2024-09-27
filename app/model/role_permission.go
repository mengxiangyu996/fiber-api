package model

// 角色权限关系模型
type RolePermission struct {
	BaseModel
	RoleId       int
	PermissionId int
}
