package model

import (
	"regexp"
	"time"
)

const (
	UoloCidName       = "uolo_cid"
	JwtCookieAuthName = "_Authorization"
)

type (
	UserProfile struct {
		Id         string    `xorm:"pk" json:"id"`
		Name       string    `xorm:"username"  json:"username" form:"username"`
		Email      string    `xorm:"email" json:"email,omitempty" form:"email"`
		Password   string    `xorm:"password" json:"-" form:"password"`
		NickName   string    `xorm:"nickname" json:"nickname" form:"nickname"`
		Avatar     string    `xorm:"avatar" json:"-" form:"avatar"`
		Phone      string    `xorm:"phone" json:"phone" form:"phone"`
		Sex        int32     `xorm:"sex" json:"sex" form:"sex"`
		Birth      time.Time `xorm:"birth" json:"birth" form:"birth"`
		Tags       int64     `xorm:"tags" json:"tags,omitempty" form:"tags"`
		Profession string    `xorm:"profession" json:"profession" form:"profession"`
		Status     int32     `xorm:"status" json:"-"`
		CreateAt   time.Time `xorm:"createtime created not null" json:"-"`
		UpdateAt   time.Time `xorm:"updatetime updated not null" json:"-"`
	}
)

func (u UserProfile) TableName() string {
	return "user_profile"
}

func CheckEmail(email string) bool {
	if m, _ := regexp.MatchString("^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+", email); !m {
		return false
	}
	return true
}
