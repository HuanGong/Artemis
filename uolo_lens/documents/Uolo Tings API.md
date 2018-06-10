# Uolo Tings API 



- /things/v1/list/todo    GET    PUBLIC

获取当前用户的未完成的Things， 当请求是未授权用户时， 返回公共things



- /things/v1/list/finished GET  Public

获取用户已经完成的Things， 当请求是未授权用户时， 返回公共things



- /things/v1/new POST    Private

需登陆， 新增things, 表单或者json

```shell
content: string //things 的内容
comment: string  // 第一个针对things的评论， 可以为空
```



- /things/v1/delete POST  Private

私有api， 需登陆， 删除things 表单或者json

```shell
uuid： string   //things 的唯一标识id

```



- /things/v1/finish   POST  Private

私有api， 需登陆， 将某things 标识为完成  表单或json

```shell
uuid： string // things 的唯一标识
comment： string // things 完成之后的最后一次评论， 可以为空
```

