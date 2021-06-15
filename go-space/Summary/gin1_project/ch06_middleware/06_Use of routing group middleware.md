# Use of routing group middleware



## 1、use

```go
v1 := router.Group("/v1")
v1.Use(GroupRouterMiddle())

{
    v1.Get...
}
```



routers/routers.go

```go
	ch01_router_group := router.Group("/")
	ch02_router_group := router.Group("/ch02")

	ch02_router_group.Use(ch05.MiddleWare1) // Use of routing group middleware

	ch03_router_group := router.Group("/ch03")
	ch04_router_group := router.Group("/ch04")
	ch05_router_group := router.Group("/ch05")
```



visit: http://localhost:8090/ch02/user_info



shell:

```shell
This is custom middleware 1, Start
This is custom middleware 1, End
total cost time: 0.007193059
[GIN] 2021/05/22 - 16:05:30 | 200 |    7.209751ms |       127.0.0.1 | GET      "/ch02/user_info"
```

