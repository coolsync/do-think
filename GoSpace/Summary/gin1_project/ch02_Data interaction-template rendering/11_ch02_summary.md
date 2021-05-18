# ch02 summary

## 一、http请求补充

```go
router := gin.Default()

router.GET("/someGet", getting)
router.POST("/somePost", posting)
router.PUT("/somePut", putting)
router.DELETE("/someDelete", deleting)
router.PATCH("/somePatch", patching)
router.HEAD("/someHead", head)
router.OPTIONS("/someOptions", options)
```

二、设置启动参数

```
router := gin.Default()

s := &http.Server{
        Addr:           ":8090"
        Handler:        router,
        ReadTimeout:    60 * time.Second,
        WriteTimeout:   60 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }

s.ListenAndServe()
```