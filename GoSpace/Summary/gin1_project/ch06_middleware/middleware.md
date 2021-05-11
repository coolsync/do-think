# Middleware



# introduce

中间件是介于应用系统和[系统软件](https://baike.baidu.com/item/系统软件/215962)之间的一类软件，它使用系统软件所提供的基础服务（功能），衔接网络上应用系统的各个部分或不同的应用，能够达到资源共享、功能共享的目的。目前，它并没有很严格的定义，但是普遍接受IDC的定义：中间件是一种独立的系统软件服务程序，分布式应用软件借助这种软件在不同的技术之间共享资源，中间件位于客户机服务器的操作系统之上，管理计算资源和网络通信。从这个意义上可以用一个等式来表示中间件：中间件=平台+通信，这也就限定了只有用于分布式系统中才能叫中间件，同时也把它与支撑软件和实用软件区分开来。



使用 BasicAuth 中间件：https://gin-gonic.com/zh-cn/docs/examples/using-basicauth-middleware/

在中间件中使用 Goroutine：https://gin-gonic.com/zh-cn/docs/examples/goroutines-inside-a-middleware/



## 一、什么是中间件

- 开发者自定义的钩子（Hook） function 
- 类似python中的装饰器

## 二、中间件的作用

- 中间件适合处理一些公共的业务逻辑，比如登录认证、权限校验、数据分页、记录日志、耗时统计等
- 需要对某一类 function 进行通用的前置或者后置处理

## 三、使用场景

1. 逻辑执行之前
   - token等认证
   - 权限校验
   - 限流
   - 数据过滤
   - 白名单
   - ...
2. 逻辑执行之后
   - 数据过滤，比如敏感词等
   - 统一的响应头等
   - ...



# Use middleware

## 一、使用中间件

```go
router := gin.New()

router.Use(gin.Logger())
router.Use(gin.Recovery())
```

注意：中间件的回调要先于用户定义的路径处理 function 

## 二、中间价的使用位置说明

中间件的使用顺序绝对了什么时候执行中间件，比如有三个路由:

```
router := gin.Default()

router.Get("/login",xxx)

router.Get("/user_list",xxx)
router.Get("/news_list",xxx)
```

加入user_list和news_list需要在登陆后才可以访问，login不要登录认证就可访问，

这时候我们需要一个token认证的中间件，那这个中间件Use的位置会有影响，如下：

```
router := gin.Default()

router.Get("/login",xxx)

router.User(MiddleWare())

router.Get("/user_list",xxx)
router.Get("/news_list",xxx)

Use不能放在login的前面，不然也会对login进行拦截认证
```

注意：Use不能放在login的前面，不然也会对login进行拦截认证

## 三、中间件执行顺序示例

```go
func MiddlewareA() gin.HandlerFunc {
    return func(c *gin.Context) {
        fmt.Println("MiddlewareA before request")
        // before request
        c.Next()
        // after request
        fmt.Println("MiddlewareA after request")
    }
}

func MiddlewareB() gin.HandlerFunc {
    return func(c *gin.Context) {
        fmt.Println("MiddlewareB before request")
        // before request
        c.Next()
        // after request
        fmt.Println("MiddlewareB after request")
    }
}
```



# Built-middleware



## 一、gin内置中间件

- func BasicAuth(accounts Accounts) HandlerFunc
- func BasicAuthForRealm(accounts Accounts, realm string) HandlerFunc    realm：认证分组
- func Bind(val interface{}) HandlerFunc //拦截请求参数并进行绑定
- func ErrorLogger() HandlerFunc       //错误日志处理
- func ErrorLoggerT(typ ErrorType) HandlerFunc //自定义类型的错误日志处理
- func Logger() HandlerFunc //日志记录
- func LoggerWithConfig(conf LoggerConfig) HandlerFunc
- func LoggerWithFormatter(f LogFormatter) HandlerFunc
- func LoggerWithWriter(out io.Writer, notlogged ...string) HandlerFunc
- func Recovery() HandlerFunc
- func RecoveryWithWriter(out io.Writer) HandlerFunc
- func WrapF(f http.HandlerFunc) HandlerFunc //将http.HandlerFunc包装成中间件
- func WrapH(h http.Handler) HandlerFunc       //将http.Handler包装成中间件



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



# Middleware 的Next和Abort

## 一、Next

在我们定义的众多中间件，会形成一条中间件链，而通过 Next  function 来对后面的中间件进行执行
特点：

- 1.当遇到c.Next() function 时 它取出所有的没被执行过的注册的 function 都执⾏⼀遍，然后再回到本 function 中，有点类似递归 function 
- 2.Next  function 是在请求前执行，而 Next  function 后是在请求后执行。
- 3.可以用在token校验，把用户id存起来供给功能性 function 使用

## 二、Abort

1. ctx.Abort()方法的作用 终止调用整个链条
2. 比如：token认证没有通过，不能直接使用return返回，而是使用Abort来终止

## 三、中间件执行顺序示例

```go
func MiddleWare1(ctx *gin.Context) {
	fmt.Println("custom middleware 1 -- start")
	ctx.Next()
	fmt.Println("custom middleware 1 -- end")
}

func MiddleWare2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("custom middleware 2 -- start")
		b := true
		if b {	// 满足 condition
			c.Abort() // abort middleware
		}
		c.Next()
		fmt.Println("custom middleware 2 -- end")

	}
}

func MiddleWare3(ctx *gin.Context) {
	fmt.Println("custom middleware 3 -- start")
	ctx.Next()
	fmt.Println("custom middleware 3 -- end")
}


router := gin.Default()

router.Use(Middleware1,Middleware2(),Middleware3)
```

## 四、利用Next计算请求时间

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



# Router组 中间件的使用

## 一、使用

```go
v1 := router.Group("/v1")
v1.Use(GroupRouterMiddle())

{
    v1.Get...
}
```



# 全局中间件的使用

## 一、使用

```go
绑定在根router上即可

router := gin.Default()

//使用自定义的全局中间件
router.Use(GlobalMiddle)
```

## 二、中间件执行顺序

全局中间件 > 路由组中间件 > 路由中间件，如果是同一类别，那就取决于append的前后顺序了



# 局部中间件的使用

## 一、自定义中间件

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

## 二、局部使用中间件

```go
// 路由映射时可以传多个HandlerFunc
router := gin.Default()

router.GET("/hello",Middle1Ware(),Hello)
```



# BasicAuth MiddleWare

## 一、BasicAuth中间件的使用

```
// 局部使用中间价
chap05.GET("/basic",gin.BasicAuth(gin.Accounts{
        "zs":"123456",
        "ls":"123",
        "ww":"1234",
    }),BasicAuthTest)


// 私有数据
var map_data map[string]interface{} = map[string]interface{}{
    "zs":gin.H{"age":18,"addr":"zs-xx"},
    "ls":gin.H{"age":19,"addr":"ls-xx"},
    "ww":gin.H{"age":20,"addr":"ww-xx"},
}


// 获取私有数据。如果没有权限则获取不到
func BasicAuthTest(ctx *gin.Context)  {

    user_name := ctx.Query("user_name")

    data ,ok := map_data[user_name]

    if ok{
        ctx.JSON(http.StatusOK,gin.H{"user_name":user_name,"data":data})
    }else {
        ctx.JSON(http.StatusOK,gin.H{"user_name":user_name,"data":"没有权限"})
    }
}
```

一文读懂HTTP Basic身份认证:https://juejin.im/entry/6844903586405564430

执行逻辑：

登录页面(没有中间件) -- 会设置session -- 其他路由回去session的key -- 获取对应的数据

## 二、WrapF

```go
gin.WrapF(IndexHandler)


func IndexHandler(w http.ResponseWriter, r *http.Request)  {
    ...
}
```

## 三、WrapH和WrapF的区别

```
需要自己去定义struct实现这个Handler接口

type TestStruct struct {}

func (test *TestStruct) TestH(w http.ResponseWriter, r *http.Request) {
    ...
}
```