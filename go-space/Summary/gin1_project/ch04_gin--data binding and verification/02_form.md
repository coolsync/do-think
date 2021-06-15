# ch04 Gin -- data bind and verify



# 1 data bind introduce



Gin提供了两类绑定方法：

- Type
  - Must bind
    - **Methods** -`Bind`,`BindJSON`,`BindXML`,`BindQuery`,`BindYAML`
    - Behavior
      - 这些方法属于`MustBindWith`的具体调用。 如果发生绑定错误，则请求终止，并触发 `c.AbortWithError(400, err).SetType(ErrorTypeBind)` 。响应状态码被设置为 400 并且`Content-Type`被设置为`text/plain; charset=utf-8` 。 如果您在此之后尝试设置响应状态码，Gin会输出日志`[GIN-debug] [WARNING] Headers were already written. Wanted to override status code 400 with 422` 。 如果您希望更好地控制绑定，考虑使用`ShouldBind`等效方法。
- Type
  - Should bind
    - **Methods** -`ShouldBind`,`ShouldBindJSON`,`ShouldBindXML`,`ShouldBindQuery`,`ShouldBindYAML`
    - Behavior
      - 这些方法属于`ShouldBindWith`的具体调用。 如果发生绑定错误，Gin 会返回错误并由开发者处理错误和请求。



一般使用 Should bind





# 2 data bind -- Should bind



## Form



`bind_form.html`:

```html
{{define "chapter04/bind_form.html"}}
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>

<body>
    form -- bind page <br>
    <form action="/do_bind_form" method="POST">
        <input type="text" name="name" placeholder="please enter name"><br>
        <input type="text" name="age" placeholder="please enter age"><br>
        <input type="text" name="addr" placeholder="please enter addr"><br>
        <input type="submit" name="submit"><br>
    </form>
</body>

</html>
{{end}}
```



`main.go`:

```go
	// Bind Form
	router.GET("/to_bind_form", chapter04.ToBindForm)
	router.POST("/do_bind_form", chapter04.DoBindForm)
```



`chapter04/bind_form.go`:

```go
package chapter04

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `form:"name"`
	Age  int    `form:"age"`
	Addr string `form:"addr"`
}

func ToBindForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chapter04/bind_form.html", nil)
}

func DoBindForm(ctx *gin.Context) {
	var user User
	err := ctx.ShouldBind(&user)
	fmt.Println(err)
	if err != nil {
		ctx.String(http.StatusNotFound, "bind form failed!")
	}

	fmt.Println(user)

	ctx.String(http.StatusOK, "bind form successful!")
}
```



