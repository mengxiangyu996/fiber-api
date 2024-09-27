package model

// 菜单模型
type Menu struct {
	BaseModel
	ParentId  int `gorm:"default:0"`
	Name      string
	Type      int
	Sort      int `gorm:"default:0"`
	Path      string
	Component string
	Icon      string
	Hidden    int
	KeepAlive int
	Redirect  string
	Status    int `gorm:"default:1"`
}
