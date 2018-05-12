

## 新建文章

POST /article/new

表单内容:
```
		ArticleNewForm struct {
			Tag     string `json:"tag" query:"tag" form:"tag"`
			Author  string `json:"author" query:"author" form:"author"`
			Title   string `json:"title" query:"title" form:"title"`
			Mime    string `json:"mime" query:"mime" form:"mime"`
			Origin  string `json:"origin" query:"origin" form:"origin"`
			Summary string `json:"summary" query:"summary" form:"summary"`
			Content string `json:"content" query:"content" form:"content"`
		}
```
响应内容:
```
		Response struct {
			Code    int32  `json:"code"`
			Message string `json:"message"`
		}
```


## 获取文章详情页

GET|POST /article/detail
参数
```
		PostContentForm struct {
			Path string `json:"path" query:"path" form:"path"`
			Type string `json:"type" query:"type" form:"type"`
		}
```
响应
```
		Response struct {
			Code    int32  `json:"code"`
			Message string `json:"message"`
			Content string `json:"content"`
		}
```

## 获取文章列表

GET /lenslist
参数 NONE

响应
```
	Article struct {
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
		Count     int32     `json:"count" xorm:"count"`           // 阅读量
	}

	LensList struct {
		Articles []*Article `json:"articles" xorm:"articles"`
	}

	type Response struct {
		Code     int32           `json:"code"`
		Message  string          `json:"message"`
		LensList *model.LensList `json:"lensList"`
	}

```


##
/utils/extract/article