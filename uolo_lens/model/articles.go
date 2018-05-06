package model

import "time"

type (
	Article struct {
		Id        int64     `json:"id" xorm:"id"`   // 唯一标示
		Tag       string    `json:"tag" xorm:"tag"` // 类型标签
		Uuid      string    `json:"uuid" xorm:"uuid"`
		Mime      string    `json:"mime" xorm:"mime"`             // 资源类型：md, html
		Title     string    `json:"title" xorm:"title"`           // 标题
		Origin    string    `json:"origin" xorm:"origin"`         // 如果是转载的，　原始来源
		Author    string    `json:"author" xorm:"author"`         // 作者
		Summary   string    `json:"summary" xorm:"summary"`       // 简介
		Rpath     string    `json:"rpath" xorm:"rpath"`           // md, html 存储路径
		CreatedAt time.Time `json:"created_at" xorm:"created_at"` // 创建时间
		UpdatedAt time.Time `json:"updated_at" xorm:"updated_at"` // 更新时间
		Status    int32     `json:"status" xorm:"status"`         // 用于内容审核，防止非法内容
		Count     int32     `json:"count" xorm:"count"`           // 阅读量
	}
)

func (db Article) TableName() string {
	return "articles"
}
