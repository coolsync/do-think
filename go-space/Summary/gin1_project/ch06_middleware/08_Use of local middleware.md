# Use of local middleware



## 1、自定义中间件

```go
 func Middle1Ware() gin.HandlerFunc{
    return func(c *gin.Context) {
        t := time.Now()
        fmt.Println("我是自定义中间件第2种定义方式---请求之前")
        //在gin上下文中定义一个变量
        c.Set("example", "CustomRouterMiddle2")
        //请求之前
        c.Next()
        fmt.Println("我是自定义中间件第2种定义方式---请求之后")
        //请求之后
        //计算整个请求过程耗时
        t2 := time.Since(t)
        fmt.Println(t2)
    }
}
```



## 2、局部使用中间件

```go
// 路由映射时可以传多个HandlerFunc
router := gin.Default()

router.GET("/hello",Middle1Ware(),Hello)
```





```go
// Ch02 Tmpls Render
ch02_routers.GET("/user", User)

ch02_routers.GET("/user_info", ch05.MiddleWare3, UserInfoStruct) // Partial use of middleware

ch02_routers.GET("/arr", ArrayHandler)
```



```shell
---- ---- This is custom middleware 3, Start
---- ---- This is custom middleware 3, End
```

