# 参数绑定

它能够基于请求自动提取JSON、form表单和QueryString类型的数据，并把值绑定到指定的结构体对象

## 一、ShouldBind

示例代码：

```go
type User struct {
    Id int `form:"id" json:"id"`               
    Name string `form:"name" json:"name"`

}

structTag：指定字段名称，不用使用首字母大写的



func Index(ctx *gin.Context)  {

    var u User
    err := contxt.ShouldBind(&u)
    fmt.Println(err)
    fmt.Println(u)
    contxt.String(http.StatusOK, "Hello %s", u.Name)

}
```

## 二、ShouldBindWith

可以使用显式绑定声明绑定 multipart form：

```
c.ShouldBindWith(&form, binding.Form)
```

或者简单地使用 ShouldBind 方法自动绑定

## 三、ShouldBindQuery

ShouldBindQuery函数只绑定 url 查询参数而忽略 post 数据