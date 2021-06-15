# Redirect



main

```go
	// Redirect
	router.GET("/redirect_a", ch02.RedirectA)
	router.GET("/redirect_b", ch02.RedirectB)
```

ch02

```go
func RedirectA(ctx *gin.Context) {
	fmt.Println("Router A")
	ctx.Redirect(http.StatusFound, "/redirect_b") // 302
	// ctx.Redirect(http.StatusFound, "https://cn.bing.com/") // 302
}

func RedirectB(ctx *gin.Context) {
	fmt.Println("Router B")
	ctx.String(http.StatusOK, "Router B")
}
```

