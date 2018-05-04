package model

type (
	Article struct {
		Id           int64  `json:"id" xorm:"id"`                       // 唯一标示
		Tag          string `json:"tag" xorm:"tag"`                     // 类型标签
		Mime         string `json:"mime" xorm:"mime"`                   // 资源类型：md, html
		Title        string `json:"title" xorm:"title"`                 // 标题
		Origin       string `json:"origin" xorm:"origin"`               // 如果是转载的，　原始来源
		Author       string `json:"author" xorm:"author"`               // 作者
		Summary      string `json:"summary" xorm:"summary"`             //　简介
		ResourcePath string `json:"resource_path" xorm:"resource_path"` // md, html 存储路径
		Status       int32  `json:"status" xorm:"status"`               // 用于内容审核，防止非法内容
	}
)

func (posts *Article) GetName() string {
	return "articles"
}
