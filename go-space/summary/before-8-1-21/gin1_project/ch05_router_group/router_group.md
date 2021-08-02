# Router Group

## Introduce:

https://gin-gonic.com/zh-cn/docs/examples/grouping-routes/

## 一、路由组的作用

1. 区分版本
2. 区分模块

## 二、使用

```go
func main() {
	router := gin.Default()

	// 简单的路由组: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}

	// 简单的路由组: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}

	router.Run(":8080")
}
```



# 路由抽取

## 一、抽取规则

1. 总路由中设置路由组
2. 模块中的路由负责映射具体的业务

## 二、Example

1.main.go中

```go
"gin_project/routers"

router := gin.Default()
// Router group
routers.Routers(router)   // routers是总路由的package名称
```

2.总路由

```go
package routers

import (
	"ginproject/controller/chapter01"
	"ginproject/controller/chapter02"
	"ginproject/controller/chapter03"
	"ginproject/controller/chapter04"

	"github.com/gin-gonic/gin"
)
func Routers(router *gin.Engine) {
	ch01 := router.Group("/chapter01")	
	ch02 := router.Group("/chapter02")
	ch03 := router.Group("/chapter03")
	ch04 := router.Group("/chapter04")

	chapter01.Router(ch01)	// chapter01 是 project 内 chapter01 mudle package 名称
	chapter02.Router(ch02)
	chapter03.Router(ch03)
	chapter04.Router(ch04)
}
```

3.模块路由

```go
package chapter04

import "github.com/gin-gonic/gin"

// Chapter04
func Router(ch04 *gin.RouterGroup) {
	// Bind Form
	ch04.GET("/to_bind_form", ToBindForm)
	ch04.POST("/do_bind_form", DoBindForm)

	// Bind query string
	ch04.GET("/bind_query_string", BindQueryString)

	// Bind Json
	ch04.GET("/to_bind_json", ToBindJson)
	ch04.POST("/do_bind_json", DoBindJson)

	// Bind Uri
	ch04.GET("/bind_uri/:name/:age/:addr", BindUri)

	// Bind valid data
	ch04.GET("/to_valid", ToBindValid)
	ch04.POST("/do_valid", DoBindValid)
}
```