# Next and Abort in middleware



## 1、Next

在我们定义的众多中间件，会形成一条中间件链，而通过 Next Function 来对后面的中间件进行执行
特点：

1. 当遇到 ctx.Next() Function时 它取出所有的没被执行过的注册的Function都执⾏⼀遍，然后再回到本Function中，有点类似递归Function

2. Next() 在 client 请求前执行，Next() 结束后 再对 请求 data handle。

3. 可以用在token校验，把用户id存起来供给功能性Function使用



## 2、Abort

1. ctx.Abort()方法的作用 终止调用整个链条
2. 比如：token认证没有通过，不能直接使用return返回，而是使用Abort来终止



## 3、中间件执行顺序示例

```go
func MiddleWare1(ctx *gin.Context)  {

    fmt.Println("这是自定义中间件1--开始")
    ctx.Next()
    fmt.Println("这是自定义中间件1--结束")
}

func MiddleWare2() gin.HandlerFunc {

    return func(ctx *gin.Context) {
        fmt.Println("这是自定义中间件2--开始")

        if 3 < 4{   // 满足条件
            ctx.Abort()
        }
        ctx.Next()
        fmt.Println("这是自定义中间件2--结束")
    }
}

func MiddleWare3(ctx *gin.Context)  {
    fmt.Println("这是自定义中间件3--开始")
    ctx.Next()
    fmt.Println("这是自定义中间件3--结束")
}


router := gin.Default()

router.Use(Middleware1,Middleware2(),Middleware3)
```



### Next and Abort Example:

```go
func MiddleWare1(ctx *gin.Context) {
	fmt.Println("This is custom middleware 1, Start")
	ctx.Next()
	fmt.Println("This is custom middleware 1, End")
}

func MiddleWare2() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("---- This is custom middleware 2, Start")
		// b := true
		// if b {
		// 	ctx.Abort() // ctx.Abort()方法的作用 终止调用整个 middle chain, back all content is not visit
		// }
		ctx.Next()
		fmt.Println("---- This is custom middleware 2, End")
	}
}

func MiddleWare3(ctx *gin.Context) {
	fmt.Println("---- ---- This is custom middleware 3, Start")
	ctx.Next()
	fmt.Println("---- ---- This is custom middleware 3, End")
}
```



use

```go
// router.Use(ch05.MiddleWare1) // global middle ware
// router.Use(ch05.MiddleWare2())
// router.Use(ch05.MiddleWare3)
router.Use(ch05.MiddleWare1, ch05.MiddleWare2(), ch05.MiddleWare3)
```





## 4、利用Next计算请求时间

```go
func Middle(ctx *gin.Context){
    t := time.Now()

    //可以设置一些公共参数
    c.Set("example", "12345")
    //等其他中间件先执行
    c.Next()
    //获取耗时
    latency := time.Since(t)
    fmt.Printf("cost time:%d us", latency/1000)
}
```