# 3 Bind Query String



main.go

```go
	// Bind query string
	router.GET("/bind_query_string", chapter04.BindQueryString)
```

bind_query_string.go

```go
package chapter04

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BindQueryString(ctx *gin.Context) {
	var user User

	err := ctx.ShouldBind(&user)
	fmt.Println(user)

	if err != nil {
		ctx.String(http.StatusNotFound, "bind query string failed!")
	}

	ctx.String(http.StatusOK, "bind query string successful!")
}
```



# 4 Bind Json



`template/chapter04/bind_json.html` :

```html
{{define "chapter04/bind_json.html"}}
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <script src="/static/js/jquery-3.6.0.min.js"></script>
</head>

<body>
    bind -- json page <br>
    <form>
        <input type="text" id="name" placeholder="please enter name"><br>
        <input type="text" id="age" placeholder="please enter age"><br>
        <input type="text" id="addr" placeholder="please enter addr"><br>
        <input type="button" value="btn_add" id="btn_add"><br>
    </form>

    <script>
        let btn_add = document.getElementById("btn_add");

        btn_add.onclick = function () {
            let name = document.getElementById("name").value;
            let age = document.getElementById("age").value;
            let addr = document.getElementById("addr").value;


            $.ajax({
                url: "/do_bind_json",
                type: "POST",
        
                contentType: "application/json",
                dataType: "json",
                data: JSON.stringify({
                    "name": name,
                    "age": Number(age),
                    "addr": addr,
                }),

                success: function (data) {
                    alert(data["code"]);
                    alert(data["msg"]);
                },
                fail: function (data) {
                    alert("server error");
                }
            })
        }
    </script>
</body>

</html>
{{end}}
```



`chapter04/bind_json.go`

```go
package chapter04

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ToBindJson(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chapter04/bind_json.html", nil)
}

func DoBindJson(ctx *gin.Context) {
	var user User

	err := ctx.ShouldBind(&user)
	
	fmt.Println(err)
	fmt.Println(user)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "failed",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}
```



`main.go`: 

```go
// Bind Json
router.GET("/to_bind_json", chapter04.ToBindJson)
router.POST("/do_bind_json", chapter04.DoBindJson)
```



# 5 Bind Uri



`bind_uri.go` : 

```go
package chapter04

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BindUri(ctx *gin.Context) {
	var user User

	err := ctx.ShouldBindUri(&user)
	fmt.Println(user)

	if err != nil {
		ctx.String(http.StatusNotFound, "Bind uri failed")
	}

	ctx.String(http.StatusOK, "Bind Uri Ok")
}
```



`main.go` : 

```go
// Bind Uri
router.GET("/bind_uri/:name/:age/:addr", chapter04.BindUri)
```



# 6 Bind valid data



# 数据验证

go-playground/validator.v8进行验证

## 一、使用

使用structTag的binding，如：binding:"required"

如果没有***空值或者类型不匹配***就会报错，重定向到400 (Bad Request)

错误信息：Key: 'User.Name' Error:Field validation for 'Name' failed on the 'required' tag

示例代码：



`chapter04/bind_valid_data.go`

`template/chapter04/bind_valid.html`



## 二、其他验证器

注意：

- 多个验证器之间用英文输入法下的逗号(,)隔开，并且是按照验证器的顺序执行的

- 如果希望在参数中包含逗号（即excludesall =,），则需要使用UTF-8十六进制表示形式0x2C

  - e.g.

    `validate："excludesall=0x2C"`

1.-                     忽略字段，如：binding:"-"

2.required：    必填字段，如：binding:"required"

3.min               最小长度，如：binding:"min=10"

4.max              最大长度，如：binding:"max=10"

5.|                    或，如：binding:"rgb|rgba"

6.***structonly***     如果有嵌套，可以决定只验证结构体上的，binding:"structonly"

7.Exists

8.omitempty  省略空，如果为空，则不会继续验证该字段上其他的规则，只有不为空才会继续验证其他的，如max等

9.dive              嵌套验证



29.***alpha***           字符串值仅包含字母字符

38.***email***              字符串值包含有效的电子邮件

39.***url***                  字符串值包含有效的网址，必须包含http://等

42.***contains***       字符串值包含子字符串值，contains=@

45.***excludes***          字符串值不包含子字符串值，excludes = @

51.***uuid***                 字符串值包含有效的UUID。

62.***ip***                        字符串值包含有效的IP地址

官方文档：

https://godoc.org/gopkg.in/go-playground/validator.v8#hdr-Baked_In_Validators_and_Tags

