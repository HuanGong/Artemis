package model

type (
	PostType int32

	Posts struct {
		Id 			int64	`json:"id" xorm:"id"`
		Type 		int32	`json:"type" xorm:"type"`
		Title   	string	`json:"title" xorm:"title"`
		From 		string	`json:"from" xorm:"from"`
		Author  	string	`json:"author" xorm:"author"`
		Summary 	string	`json:"summary" xorm:"summary"`
		OffsetPath 	string	`json:"offsetpath" xorm:"offsetpath"`
	}
)
const (
	PostTypeOriginal  PostType 	= 0
	PostTypeThirdpart PostType 	= 1
)

func (posts *Posts) GetName() string {
	return "posts"
}
