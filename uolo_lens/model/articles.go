package model

type (
	PostType   int32
	SourceType int32
	Posts      struct {
		Id           int64  `json:"id" xorm:"id"`                       // 唯一标示
		Tag          string `json:"tag" xorm:"tag"`                     // 类型标签
		Type         int32  `json:"type" xorm:"type"`                   // 来源，转载还是原创的
		Source       int32  `json:"source" xorm:"source"`               // 资源类型：markdown, html
		Title        string `json:"title" xorm:"title"`                 // 标题
		From         string `json:"from" xorm:"from"`                   // 如果是转载的，　原始来源
		Author       string `json:"author" xorm:"author"`               // 作者
		Summary      string `json:"summary" xorm:"summary"`             //　简介
		ResourcePath string `json:"resource_path" xorm:"resource_path"` // md, html 存储路径
	}

)

const (
	TypeOriginal  PostType = 0
	TypeReproduce PostType = 1

	SourceMarkdown SourceType = 0
	SourceHtml     SourceType = 1
)

func (posts *Posts) GetName() string {
	return "posts"
}
