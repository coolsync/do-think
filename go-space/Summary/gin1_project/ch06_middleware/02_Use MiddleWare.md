# 中间件的使用

## 一、使用中间件

```go
router := gin.Default()
// =
router := gin.New()

// router.Use(gin.Logger(), gin.Recovery())
// =
router.Use(gin.Logger())
router.Use(gin.Recovery())
```

注意：中间件的回调要先于用户定义的路径处理函数

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

Use不能放在login的前面，会对login进行拦截认证
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