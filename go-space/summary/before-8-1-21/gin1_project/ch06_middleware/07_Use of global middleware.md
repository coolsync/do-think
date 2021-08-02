# Use of global middleware



## 1、Use

```go
绑定在根router上即可

router := gin.Default()

//使用自定义的全局中间件
router.Use(GlobalMiddle)
```



```go
// router.Use(ch05.MiddleWare1) // global middle ware
// router.Use(ch05.MiddleWare2())
// router.Use(ch05.MiddleWare3)
router.Use(ch05.MiddleWare1, ch05.MiddleWare2(), ch05.MiddleWare3)
```



## 二、中间件执行顺序

全局中间件 > 路由组中间件 > 路由中间件，如果是同一类别，那就取决于append的前后顺序了