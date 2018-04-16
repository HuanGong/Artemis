package model

type (
	User struct {
		Id       string `xorm:"pk" json:"id"`
		Name     string `xorm:"username"  json:"username" form:"username"`
		Email    string `xorm:"email" json:"email,omitempty" form:"email"`
		Password string `xorm:"password" json:"password,omitempty" form:"password"`
	}
)

func (u User) TableName() string {
	return "uolo_users"
}
