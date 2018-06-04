package model

import (
	"time"
)

/*
	这个功能就像一个日记本， 希望类似于Sticks Note
	它告诉自己什么事情是重要的，敦促自己，过程中能对
	一个UOLO事件完成前，可以写下很多comment，当一件事情完成后， 能看到自己做整个事情的一个心理路程
*/
const (
	UoloThingTableName   = "things"
	ThingComentTableName = "things_comment"
)

type (
	UoloThing struct {
		Uuid        string    `json:"uuid" xorm:"uuid"`
		Owner       string    `json:"-" xorm:"owner"`
		Content     string    `json:"content" xorm:"content"`
		StartTime   time.Time `json:"start_time" xorm:"start_time"`
		ArchiveGoal bool      `json:"archive_goal" xorm:"archive_goal"`
	}

	ThingsComment struct {
		ThingsId string `json:"thing_id" xorm:"thing_id"` // 对应UoloThings的Uuid
		Comment  string `json:"comment" xorm:"comment"`
	}
)

func (data UoloThing) TableName() string {
	return UoloThingTableName
}

func (data ThingsComment) TableName() string {
	return ThingComentTableName
}
