package model

// 管理员模型
type Admin struct {
	BaseModel
	Username string
	Nickname string
	Gender   int
	Email    string
	Phone    string
	Password string
	Avatar   string
	Status   int `gorm:"default:1"`
}
