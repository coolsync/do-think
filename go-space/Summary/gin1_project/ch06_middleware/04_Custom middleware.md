# Custom middleware



## 一、自定义中间件的两种方式

```go
//自定义中间件第1种定义方式
func MiddleWare1(ctx *gin.Context)  {

    fmt.Println("这是自定义中间件1")
}

自定义中间件第2种定义方式
func MiddleWare2() gin.HandlerFunc  {
    return func(ctx *gin.Context) {

        fmt.Println("这是自定义i中间件2")
    }
}



router := gin.New()

router.Use(MiddleWare1)      // 需要加括号
router.Use(MiddleWare2())        // 不需要加括号，当成参数
```



## Middle Ware 1

main.go:

```go
router := gin.Default()
// router := gin.New()
// router.Use(gin.Logger(), gin.Recovery()) // built middleware

router.Use(ch05.MiddleWare1) // global middle ware

// Set consume tpl func
router.SetFuncMap(template.FuncMap{
    "add_num":   ch03.AddNum,
    "str_len":   ch03.SubStr,
    "safe_html": ch03.SafeHTML,
})

...

// Use tmpl Regular, specify 多级目录
router.LoadHTMLGlob("templates/**/*")

// Specify static file dir
router.Static("/static", "static")

...

```



controllers/ch05/middleware.go:

```go
func MiddleWare1(ctx *gin.Context) {
	fmt.Println("---- The middleware 1")
}
```



visit: http://localhost:8090/

result:

```shell
---- The middleware 1
[GIN] 2021/05/22 - 15:03:52 | 200 |    6.778062ms |       127.0.0.1 | GET      "/"
---- The middleware 1
[GIN] 2021/05/22 - 15:03:52 | 200 |    8.335864ms |       127.0.0.1 | GET      "/static/css/index.css"
---- The middleware 1
[GIN] 2021/05/22 - 15:03:52 | 200 |     416.531µs |       127.0.0.1 | GET      "/static/images/OIP.jpeg"
```



## Middle Ware 2



main.go:

```go
// router.Use(ch05.MiddleWare1) // global middle ware
router.Use(ch05.MiddleWare2())
```





controllers/ch05/middleware.go:

```go
func MiddleWare2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("---- This is custom middleware 2")
	}
}
```



visit: http://localhost:8090/ch02/user_info



shell:

```shell
---- This is custom middleware 2
[GIN] 2021/05/22 - 15:15:14 | 200 |    6.214591ms |       127.0.0.1 | GET      "/ch02/user_info"
```

