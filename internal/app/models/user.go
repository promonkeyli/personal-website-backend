package models

// User represents a tool entity
// @Description User represents a user entity
type User struct {
	ID       uint   `json:"id,omitempty" gorm:"column:id" swaggerignore:"true"`
	UserName string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
}

// TableName 自定义表名
func (User) TableName() string {
	return "user"
}
