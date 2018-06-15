package model

import (
	"regexp"
	"time"
)

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

/*CREATE TABLE `user_profile` (
  `id` varchar(40) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '用户主键Id',
  `username` 	varchar(24)   NOT NULL COMMENT '用户名',
  `password` 	varchar(64)   NOT NULL COMMENT  '密文密码',
  `nickname` 	varchar(24)   NOT NULL COMMENT '昵称',
  `avast` 	 	varchar(50)   NOT NULL DEFAULT '' COMMENT '头像Id',
  `phone` 		varchar(50)   NOT NULL DEFAULT '' COMMENT '用户手机号',
  `email` 		varchar(50)   NOT NULL COMMENT '用户邮箱',
  `sex` 			int(2)    NOT NULL DEFAULT 0 COMMENT '性别：0未知 1男，2女',
  `birth` 		DATE                  COMMENT '用户年龄',
  `tags` 			bigint(20)    DEFAULT NULL COMMENT '用户标签',
  `status` 		tinyint(2)    NOT NULL DEFAULT '0' COMMENT '逻辑删除位',
  `createtime` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updatetime` timestamp    NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_username_password` (`username`,`password`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;*/
type (
	UserProfile struct {
		Id       string    `xorm:"pk" json:"id"`
		Name     string    `xorm:"username"  json:"username" form:"username"`
		Email    string    `xorm:"email" json:"email,omitempty" form:"email"`
		Password string    `xorm:"password" json:"-" form:"password"`
		NickName string    `xorm:"nickname" json:"nickname" form:"nickname"`
		AvastIco string    `xorm:"avast" json:"avast" form:"avast"`
		Phone    string    `xorm:"phone" json:"phone" form:"phone"`
		Sex      int32     `xorm:"sex" json:"sex" form:"sex"`
		Birth    time.Time `xorm:"birth" json:"birth" form:"birth"`
		Tags     int64     `xorm:"tags" json:"-"`
		Status   int32     `xorm:"status" json:"-"`
		CreateAt time.Time `xorm:"createtime created not null" json:"createtime"`
		UpdateAt time.Time `xorm:"updatetime updated not null" json:"updatetime"`
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
