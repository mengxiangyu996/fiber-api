package model

// 权限模型
type Permission struct {
	BaseModel
	Name      string
	GroupName string
	Path      string
	Method    string
	Status    int `grom:"default:1"`
}
