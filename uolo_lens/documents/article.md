# article相关的内容与测试命令

## 文章新建流程

- 第三方来源类型的

第三方来源URL， 通过API `/utils/extract/article` 接口解析抽取出主要文章结构，暂时只开放输出成markdown形式, 陈旭智能分析出下面内容（可能缺失部分内容）
```go
		Response struct {
			Code     int32  `json:"code" query:"code" form:"code"`
			Message  string `json:"message" query:"message" form:"message"`
			Author   string `json:"author" query:"author" form:"author"`
			Title    string `json:"title" query:"title" form:"title"`
			Mime     string `json:"mime" query:"mime" form:"mime"`
			Desc     string `json:"desc" query:"desc" form:"desc"`
			Origin   string `json:"origin" query:"origin" form:"origin"`
			Date     string `json:"date" query:"date" form:"date"`
			Keywords string `json:"keywords" query:"keywords" form:"keywords"`
			Content  string `json:"content" query:"content" form:"content"`
		}
```
得到程序返回的json格式内容后， 检查原始来源，author，Desc，关键字， content 是否正确或缺失， 少量的修正后通过API
```POST /lens/article/new`  加**表单**内容的形式进行提交, 目前此API需要登录授权，

    - [] 增加公开API的形式提交内容，需要后台审核界面，确认内容ok后发布到Lens文章数据库

- 原创类型文章

    - 有文章在其他平台发表过的情况下， 提供URL的方式通过上面的方式提交

    - 没有发表的一手资料， 建议以pdf 或者 markdown 文章的形式直接上传
        - [] 需要后台界面的支持


### 文章列表获取流程

由用户id为key存储在Redis中的一份推荐列表作为数据源，通过uuid作为索引获取到文章简洁格式以json数据格式放回给客户端


